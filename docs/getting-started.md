# Getting started

## Requirements

- Docker Engine **26 or newer** with the Compose plugin (`docker compose`). Version 26 is needed for volume `subpath` mounts.
- Linux host or a Linux VM. The gateway needs the `NET_ADMIN` and `NET_RAW` capabilities to manage iptables and sniff SPA packets.
- About 2 GB of free disk space for the images.

## Start the lab

```bash
git clone https://github.com/antoniopaya22/Software-Defined-Perimeter.git
cd Software-Defined-Perimeter

cp .env.example .env
# Replace every "change-me" value, for example with: openssl rand -base64 24

docker compose up -d --build --wait
```

`--wait` returns when the database and the controller are healthy. On the first start the controller creates the CA and issues credentials for every member; this takes a few seconds.

## Run the end-to-end check

```bash
./scripts/demo.sh
```

The script verifies four things and exits non-zero if any of them fails:

1. Before SPA, the protected service on the gateway does not answer.
2. The client cannot reach the private network directly.
3. `fwknop -n service_gate` sends the SPA packet and completes a session with the controller.
4. After SPA, the protected service answers with HTTP 200.

The same steps by hand:

```bash
docker compose exec sdp-client curl -m 4 http://10.231.0.20:8080/   # times out
docker compose exec sdp-client fwknop -n service_gate                # send SPA
docker compose exec sdp-client curl http://10.231.0.20:8080/         # nginx page
```

!!! info "Access window"
    The gateway keeps the path open for 30 seconds and only for the client's source address. Another host on the same network still sees nothing, even during that window.

## Watch what happens

```bash
docker compose logs -f sdp-gateway sdp-controller
```

When the SPA packet arrives, the gateway logs the rules it adds and removes:

```text
(stanza #0) SPA Packet from IP: 10.231.0.40 received with access source match
Added FORWARD rule to FWKNOP_FORWARD for 10.231.0.40 -> 10.231.1.30 port 80, expires at ...
Added DNAT rule to FWKNOP_PREROUTING for 10.231.0.40 -> 0.0.0.0/0 port 8080, expires at ...
```

## Reset

```bash
docker compose down -v
```

This removes the containers and every volume: database, CA, issued credentials and rotated keys. The next `up` starts from scratch.

## Inspect the database (optional)

```bash
docker compose --profile debug up -d phpmyadmin
```

phpMyAdmin is bound to <http://127.0.0.1:8081> only. Log in with the `DB_USER` and `DB_PASSWORD` from `.env`.
