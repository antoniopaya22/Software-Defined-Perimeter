# Troubleshooting

## `docker compose` asks for variables

```text
required variable DB_PASSWORD is missing a value: set DB_PASSWORD in .env
```

Create `.env` from `.env.example` and set the passwords.

## The subnet overlaps with an existing network

```text
Pool overlaps with other one on this address space
```

Set `PUBLIC_SUBNET`, `PRIVATE_SUBNET` and the IP variables in `.env` to a free range, keeping every address inside its subnet.

## The gateway cannot change iptables

The gateway needs the `NET_ADMIN` and `NET_RAW` capabilities, which `docker-compose.yml` grants. Rootless Docker, some hardened hosts and some CI runners do not allow them.

## `volume subpath` is not supported

Docker Engine 26 or newer is required for the `subpath` mounts that give each container its own credential bundle. Check with `docker version`.

## The demo fails at step 4

Look at the gateway logs right after sending the SPA packet:

```bash
docker compose logs --since 2m sdp-gateway
```

- **No `SPA Packet from IP` line:** the packet did not reach the gateway. Check `GATEWAY_IP` and that the client and the gateway share `sdp-public`.
- **`SPA Packet ... received` but no rules added:** the client's keys do not match the gateway's access list. This happens when credentials are regenerated for only one side, for example after deleting a single volume. Reset everything with `docker compose down -v`.
- **Rules added but the request times out:** check that `protected-service` is running and that `PROTECTED_SERVICE_IP` matches its address.

## The database was initialized with wrong data

The seed runs only on the first start. Remove the volumes to start again:

```bash
docker compose down -v
```

## Logs

```bash
docker compose logs -f sdp-controller sdp-gateway
docker compose exec sdp-gateway iptables -S
docker compose exec sdp-gateway iptables -t nat -S
```
