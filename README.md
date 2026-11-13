# Software-Defined Perimeter Lab

[![Paper](https://img.shields.io/badge/IJIS_2024-10.1007%2Fs10207--024--00863--7-0E7490)](https://doi.org/10.1007/s10207-024-00863-7)
[![CI](https://github.com/antoniopaya22/Software-Defined-Perimeter/actions/workflows/ci.yml/badge.svg)](https://github.com/antoniopaya22/Software-Defined-Perimeter/actions/workflows/ci.yml)
[![Docs](https://github.com/antoniopaya22/Software-Defined-Perimeter/actions/workflows/docs.yml/badge.svg)](https://antoniopaya22.github.io/Software-Defined-Perimeter/)
[![License](https://img.shields.io/github/license/antoniopaya22/Software-Defined-Perimeter?color=green)](LICENSE)

A reproducible **Software-Defined Perimeter (SDP)** lab built with Docker Compose. A protected service stays invisible behind a default-drop gateway until a client authenticates with **Single Packet Authorization (SPA)**; then the gateway opens a short-lived path for that client alone.

It is the companion lab of the article **[SecureSDP: a novel software-defined perimeter implementation for enhanced network security and scalability](https://doi.org/10.1007/s10207-024-00863-7)** (*International Journal of Information Security*, 2024). An earlier version of that work was released as an [SSRN preprint](https://doi.org/10.2139/ssrn.4326496) in 2023. The lab deploys the open-source SDP Controller and the SDP-enabled fork of fwknop from [Waverley Labs](https://github.com/waverleylabs).

📖 **Documentation: <https://antoniopaya22.github.io/Software-Defined-Perimeter/>**

![SDP architecture](docs/assets/sdp-architecture.svg)

## Quick start

Requirements: Docker Engine 26+ with the Compose plugin, on Linux.

```bash
cp .env.example .env        # then replace every "change-me"
docker compose up -d --build --wait
./scripts/demo.sh
```

```text
[1/4] Request to http://10.231.0.20:8080/ before SPA (expect: blocked)
  PASS no response, the service is hidden
[2/4] Direct request to the private network (expect: unreachable)
  PASS private network not reachable from the client
[3/4] Sending SPA packet: fwknop -n service_gate
  PASS SPA sent and controller session completed
[4/4] Request to http://10.231.0.20:8080/ after SPA (expect: HTTP 200)
  PASS protected service reachable through the gateway
```

Reset everything (database, CA, keys) with `docker compose down -v`.

| Container | Role |
|---|---|
| `sdp-controller` | Waverley Labs SDP Controller: PKI, mutual TLS, service and access lists. |
| `sdp-db` | MariaDB with members, services and permissions (internal network only). |
| `sdp-gateway` | fwknop in a default-drop firewall; opens per-client DNAT/FORWARD rules after a valid SPA packet. |
| `protected-service` | Accepting Host (nginx), reachable only through the gateway. |
| `sdp-client` | Initiating Host: sends SPA packets and refreshes its credentials. |

The documentation covers the [architecture](https://antoniopaya22.github.io/Software-Defined-Perimeter/architecture/), [configuration](https://antoniopaya22.github.io/Software-Defined-Perimeter/configuration/) (adding clients and services), [component hardening with Egida](https://antoniopaya22.github.io/Software-Defined-Perimeter/hardening/) and [troubleshooting](https://antoniopaya22.github.io/Software-Defined-Perimeter/troubleshooting/).

> **Lab use only.** This setup is for research, teaching and testing. Review it carefully before reusing any part in production.

## Research line

| Year | Work | Code |
|---|---|---|
| 2022 | Egida: automated security configuration deployment ([C&S](https://doi.org/10.1016/j.cose.2022.102638)) | [Egida-Kassandra/egida](https://github.com/Egida-Kassandra/egida) |
| 2024 | **SecureSDP**: Bridge-model SDP with per-component hardening ([IJIS](https://doi.org/10.1007/s10207-024-00863-7)) | **this lab** · [Secure-SDP](https://github.com/antoniopaya22/Secure-SDP) |
| 2024 | Apollon: IDS robust to adversarial machine learning ([C&S](https://doi.org/10.1016/j.cose.2023.103546)) | [apollon](https://github.com/antoniopaya22/apollon) |
| 2025 | SDP with Keycloak IAM and IDS ([IJIS](https://doi.org/10.1007/s10207-025-01099-9)) | [iam-ids-sdp](https://github.com/antoniopaya22/iam-ids-sdp) |

More in the [research line](https://antoniopaya22.github.io/Software-Defined-Perimeter/research-line/) page.

## Citation

If you use this lab, please cite:

```bibtex
@article{paya2024securesdp,
  title   = {SecureSDP: a novel software-defined perimeter implementation for enhanced network security and scalability},
  author  = {Pay{\'a}, Antonio and Garc{\'i}a-D{\'i}az, Vicente and G{\'o}mez, Alberto},
  journal = {International Journal of Information Security},
  volume  = {23},
  number  = {4},
  pages   = {2793--2808},
  year    = {2024},
  doi     = {10.1007/s10207-024-00863-7}
}
```

## Contributing and security

Issues and pull requests are welcome. CI builds the whole lab and runs `scripts/demo.sh` on every pull request. To report a vulnerability, see [SECURITY.md](SECURITY.md).

## License

[Apache 2.0](LICENSE)
