# Architecture

The lab follows the Cloud Security Alliance SDP model: a **controller** decides who may reach what, a **gateway** (Accepting Host side) enforces it, and **clients** (Initiating Hosts) must authenticate before any service is visible to them.

## Networks

```mermaid
flowchart LR
  subgraph public["sdp-public (10.231.0.0/24)"]
    client["sdp-client<br/>Initiating Host<br/>10.231.0.40"]
    ctrl["sdp-controller<br/>10.231.0.10"]
    gwpub["sdp-gateway<br/>10.231.0.20"]
  end
  subgraph private["sdp-private (internal, 10.231.1.0/24)"]
    gwpriv["sdp-gateway<br/>10.231.1.20"]
    svc["protected-service<br/>nginx, 10.231.1.30:80"]
  end
  subgraph mgmt["sdp-mgmt (internal)"]
    db[("sdp-db<br/>MariaDB")]
  end
  client -- "1. SPA (UDP 62201)" --> gwpub
  client -- "2. mTLS :5000" --> ctrl
  gwpub -- "mTLS :5000" --> ctrl
  ctrl --- db
  gwpub -. "same container" .- gwpriv
  gwpriv -- "3. DNAT :8080 → :80" --> svc
```

| Network | Members | Purpose |
|---|---|---|
| `sdp-public` | client, gateway, controller | The untrusted side. The gateway drops everything here by default. |
| `sdp-private` | gateway, protected service | Internal Docker network: only reachable through the gateway. |
| `sdp-mgmt` | controller, database | Internal Docker network: the database is never exposed. |

## Components

### SDP Controller

The [Waverley Labs SDP Controller](https://github.com/waverleylabs/SDPController) (Node.js), pinned to a fixed commit. It:

- authenticates every member with **mutual TLS**: the CN of each certificate is the member's SDP ID;
- issues and rotates credentials: a TLS certificate plus SPA encryption and HMAC keys;
- sends each gateway its **service list** (which ports to forward where) and its **access list** (which SDP IDs may use which services, with their SPA keys).

On start, `controller/entrypoint.sh` creates the CA if it does not exist yet. It then runs `genCredentials.js` for each SDP ID that has no keys, and exports one credential bundle per gateway and client. Each container mounts only its own bundle, and the CA private key never leaves the controller.

### SDP Gateway

The SDP-enabled [fwknop fork](https://github.com/waverleylabs/fwknop), built from source on Debian 12. The entrypoint sets `INPUT` and `FORWARD` to `DROP` and starts `fwknopd`. fwknopd gets its configuration from the controller and sniffs SPA packets on the public interface. For each valid packet it adds a DNAT and FORWARD rule limited to the client's IP address and the authorized service, and removes them when they expire.

### SDP Client

The fwknop client in SDP mode. `fwknop -n service_gate` sends an encrypted, HMAC-authenticated SPA packet that names the client's SDP ID and the requested service IDs. It then opens a control session with the controller, which rotates the client's keys and notifies the gateway.

## Bootstrap sequence

```mermaid
sequenceDiagram
  autonumber
  participant DB as sdp-db
  participant C as sdp-controller
  participant G as sdp-gateway
  participant IH as sdp-client
  DB->>DB: load schema and seed data (first start)
  C->>C: create CA (if missing)
  C->>DB: genCredentials for SDP IDs 1, 2, 3 (store SPA keys)
  C->>C: export bundles (cert, key, SPA keys, CA cert)
  G->>G: iptables default DROP, render config from bundle
  G->>C: mTLS connect (SDP ID 1)
  C-->>G: credentials good + credential update
  C-->>G: service list (service 1: tcp/8080 → 10.231.1.30:80)
  C-->>G: access list (SDP ID 3 may use service 1, with its SPA keys)
  IH->>IH: render ~/.fwknoprc and control client config from bundle
```

## Access sequence

```mermaid
sequenceDiagram
  autonumber
  participant IH as sdp-client
  participant G as sdp-gateway
  participant C as sdp-controller
  participant AH as protected-service
  IH->>G: TCP :8080 (no rule) → dropped
  IH->>G: SPA packet (SDP ID 3, service 1), encrypted + HMAC
  G->>G: verify against access list, add DNAT/FORWARD for 10.231.0.40 (30 s)
  IH->>C: mTLS control session
  C-->>IH: credential update (new SPA keys)
  C-->>G: access update (client's new keys)
  IH->>G: HTTP :8080
  G->>AH: DNAT to 10.231.1.30:80 (masqueraded)
  AH-->>IH: HTTP 200
```

## Design choices

- **No secrets in git.** Passwords come from `.env`, and the CA and member credentials are generated at runtime into Docker volumes.
- **Credentials survive restarts.** Keys rotated by the controller are kept in the `gateway-state` and `client-home` volumes.
- **Pinned upstream code.** SDPController and fwknop are cloned at fixed commits, so builds are reproducible.
- **fwknop SNAT disabled.** fwknopd builds its MASQUERADE rule with `--to-ports <gateway port>`. Every forwarded connection would then reuse one source port, and back-to-back connections would collide in conntrack. The gateway adds a plain MASQUERADE rule on the private side instead. Filtering still happens per client in fwknopd's FORWARD rules.
