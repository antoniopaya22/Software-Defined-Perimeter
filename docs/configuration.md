# Configuration

## Environment (`.env`)

Copy `.env.example` to `.env`. Only the passwords are required; everything else has working defaults.

| Variable | Default | Description |
|---|---|---|
| `DB_ROOT_PASSWORD` | — (required) | MariaDB root password. |
| `DB_USER` | `sdp` | Database user used by the controller. |
| `DB_PASSWORD` | — (required) | Password for `DB_USER`. |
| `CA_KEY_PASSWORD` | — (required) | Passphrase protecting the CA private key. The key never leaves the controller. |
| `PUBLIC_SUBNET` | `10.231.0.0/24` | Subnet of `sdp-public`. |
| `PRIVATE_SUBNET` | `10.231.1.0/24` | Subnet of `sdp-private`. |
| `CONTROLLER_IP` | `10.231.0.10` | Controller address on the public network. |
| `GATEWAY_IP` | `10.231.0.20` | Gateway public address; clients send SPA packets here. |
| `GATEWAY_PRIVATE_IP` | `10.231.1.20` | Gateway address on the private network. |
| `CLIENT_IP` | `10.231.0.40` | Client address. |
| `PROTECTED_SERVICE_IP` | `10.231.1.30` | Address of the protected nginx. |
| `PROTECTED_PORT` | `8080` | Gateway port forwarded to the protected service once authorized. |

If the subnets clash with networks on your machine, change the subnet and IP variables together.

## Members, services and permissions

The controller reads everything from the database. On the first start, MariaDB runs:

- `database/01-schema.sql`: the SDPController schema, including the triggers that tell the controller to push updates.
- `database/02-seed.sh`: the example data.

| Table | Seed rows |
|---|---|
| `sdpid` | `1` gateway, `2` controller, `3` client |
| `service` | `1` `protected-web` |
| `service_gateway` | service 1 on gateway 1: `tcp/${PROTECTED_PORT}` → `${PROTECTED_SERVICE_IP}:80` |
| `sdpid_service` | SDP ID 3 may use service 1 |

The seed only runs when the `db-data` volume is empty. To apply changes, run `docker compose down -v`. Alternatively, edit the tables on a running lab: the controller's database triggers push the changes to the gateway within a few seconds.

## Adding a second client

1. Add the member and its permission in `database/02-seed.sh`:

    ```sql
    INSERT INTO sdpid (sdpid, valid, type, country, state, locality, org, org_unit, alt_name, email, serial)
    VALUES (4, 1, 'client', 'ES', 'Asturias', 'Oviedo', 'SDP Lab', 'SDP', 'client-2', 'client2@sdp.test', 'cli-0002');
    INSERT INTO sdpid_service (id, sdpid, service_id) VALUES (2, 4, 1);
    ```

2. Ask the controller to issue its credentials: set `SDP_IDS: "1 2 3 4"` for `sdp-controller` in `docker-compose.yml`.

3. Add a client service that mounts bundle `4`:

    ```yaml
    sdp-client-2:
      <<: *defaults
      build: ./client
      environment:
        SDP_ID: "4"
        GATEWAY_IP: ${GATEWAY_IP:-10.231.0.20}
        SERVICE_IDS: "1"
        CTRL_ADDR: sdp-controller
      volumes:
        - type: volume
          source: bundles
          target: /bundle
          read_only: true
          volume:
            subpath: "4"
        - client-2-home:/root
      networks:
        sdp-public:
          ipv4_address: 10.231.0.41
      depends_on:
        sdp-controller:
          condition: service_healthy
    ```

    Then declare `client-2-home:` under `volumes:`.

4. Recreate the lab: `docker compose down -v && docker compose up -d --build --wait`.

## Adding a service

Add a row to `service`, and map it to the gateway in `service_gateway` (protocol, gateway port, NAT target IP and port). Then grant it to clients in `sdpid_service`. The NAT target must be reachable from the gateway's private network. On the client, list the new service ID in `SERVICE_IDS`.

## Upstream versions

| Component | Where | Pin |
|---|---|---|
| SDPController | `controller/Dockerfile` | `SDPCONTROLLER_COMMIT` build arg |
| fwknop (Waverley Labs fork) | `gateway/Dockerfile`, `client/Dockerfile` | `FWKNOP_COMMIT` build arg |
| Node.js | `controller/Dockerfile` | `node:22-bookworm-slim` |
| MariaDB, nginx, phpMyAdmin | `docker-compose.yml` | image tags |

Dependabot proposes updates for the base images, the Compose images and the GitHub Actions every week. CI runs the end-to-end demo on every pull request.
