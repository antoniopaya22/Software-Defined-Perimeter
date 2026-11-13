---
template: home.html
title: Software-Defined Perimeter Lab
description: Reproducible, CI-tested Software-Defined Perimeter lab with fwknop Single Packet Authorization. Companion lab of SecureSDP (IJIS 2024).
hide:
  - navigation
  - toc
---

# Software-Defined Perimeter Lab

A reproducible, CI-tested Software-Defined Perimeter built with Docker Compose. Services stay hidden behind a default-deny gateway until a client authenticates with Single Packet Authorization (SPA) using fwknop. The SDP controller runs the PKI, authenticates members with mutual TLS and pushes service and access lists to the gateway.

Companion lab of SecureSDP, International Journal of Information Security, 2024 (doi:10.1007/s10207-024-00863-7). Quick start: `cp .env.example .env`, `docker compose up -d --build --wait`, `./scripts/demo.sh`.
