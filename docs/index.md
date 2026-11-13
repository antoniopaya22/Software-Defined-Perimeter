---
template: home.html
title: Software-Defined Perimeter Lab
hide:
  - navigation
  - toc
hero:
  eyebrow: Companion lab · SecureSDP, IJIS 2024
  title: Invisible <em>until</em> authorized.
  lead: >-
    A reproducible Software-Defined Perimeter in one command. The protected
    service drops every connection until a client sends a single encrypted,
    authenticated packet. Then a path opens for that client alone, for
    thirty seconds.
  buttons:
    - text: Get started
      link: getting-started/
    - text: How it works
      link: architecture/
    - text: Read the paper
      link: https://doi.org/10.1007/s10207-024-00863-7
  meta:
    - Single Packet Authorization (fwknop)
    - Mutual TLS controller
    - Docker Compose · CI-tested
  denied: "DROP  10.231.0.40 → :8080  no rule"
  granted: "ACCEPT  10.231.0.40 → :8080  30 s"
---

# Software-Defined Perimeter Lab

## One packet opens the door

<p class="sdp-kicker">The Software-Defined Perimeter turns "connect, then authenticate" around. Nothing on the network answers until the client has proven who it is.</p>

<ol class="sdp-steps">
  <li><strong>Hidden by default</strong>The gateway's firewall drops everything. A port scan finds no service and gets no reply.</li>
  <li><strong>One packet</strong>The client sends a single SPA packet, encrypted and HMAC-authenticated, that names its identity and the service it wants.</li>
  <li><strong>Policy check</strong>The gateway checks the packet against the access list that the controller pushed over mutual TLS.</li>
  <li><strong>A path for one</strong>A DNAT rule opens for that client's address only, and expires after 30 seconds.</li>
</ol>

## What's in the lab

<div class="grid cards" markdown>

- :material-brain:{ .lg } **SDP Controller**

    ---

    Waverley Labs controller. Runs the PKI, authenticates every member with mutual TLS, rotates keys and pushes service and access lists.

- :material-wall-fire:{ .lg } **SDP Gateway**

    ---

    fwknop daemon behind a default-drop iptables firewall. It sniffs SPA packets and opens per-client DNAT and FORWARD rules.

- :material-laptop:{ .lg } **Initiating Host**

    ---

    The client. It sends SPA packets and refreshes its credentials with the controller on every session.

- :material-server-network:{ .lg } **Accepting Host**

    ---

    An nginx service on an internal network, reachable only through the gateway.

</div>

## Results reported in the paper

<p class="sdp-kicker">SecureSDP hardens every component with <a href="https://github.com/Egida-Kassandra/egida">Egida</a> as well as hiding services. The paper audited the SDP controller before and after applying the hardening controls.</p>

<div class="sdp-stats">
  <div class="sdp-stat"><span class="sdp-stat__value">+65%</span><span class="sdp-stat__label"><strong>Lynis</strong> hardening score<br>52 → 86</span></div>
  <div class="sdp-stat"><span class="sdp-stat__value">+78%</span><span class="sdp-stat__label"><strong>Chef InSpec</strong> CIS benchmark<br>36.2 → 64.3</span></div>
  <div class="sdp-stat"><span class="sdp-stat__value">+30%</span><span class="sdp-stat__label"><strong>OpenSCAP</strong> CIS Level 1<br>63.5 → 82.9</span></div>
  <div class="sdp-stat"><span class="sdp-stat__value">0</span><span class="sdp-stat__label"><strong>Safety-critical failures</strong> after hardening, only warnings</span></div>
</div>

<p class="sdp-footnote">Values as reported in Payá et al., <em>IJIS</em> 2024, Table 2. They were not regenerated with this lab. See <a href="hardening/">Component hardening</a>.</p>

## Run it in a minute

<div class="sdp-terminal" markdown>

```console
$ cp .env.example .env            # set the passwords
$ docker compose up -d --build --wait
$ ./scripts/demo.sh
[1/4] Request before SPA ................ PASS  service hidden
[2/4] Direct access to private net ...... PASS  unreachable
[3/4] fwknop -n service_gate ............ PASS  SPA accepted
[4/4] Request after SPA ................. PASS  HTTP 200
```

</div>

Every pull request runs the same demo in CI, so the lab keeps working.

## Part of a research line

<ol class="sdp-timeline">
  <li><span class="sdp-timeline__year">2022 · COMPUTERS &amp; SECURITY</span><span class="sdp-timeline__title">Egida</span><p>Automated security configuration with a DSL that avoids breaking legitimate services.</p></li>
  <li class="is-current"><span class="sdp-timeline__year">2024 · IJIS</span><span class="sdp-timeline__title">SecureSDP: this lab</span><p>Bridge-model SDP and per-component dynamic hardening.</p></li>
  <li><span class="sdp-timeline__year">2025 · IJIS</span><span class="sdp-timeline__title">SDP with IAM and intrusion detection</span><p>Keycloak MFA and roles, plus Apollon-based behaviour monitoring.</p></li>
</ol>

[Explore the full research line →](research-line.md)

<div class="sdp-cta" markdown>

## Use it in your research

Clone it, adapt the policy, and cite the paper if it helps your work.

[Get started](getting-started.md){ .md-button .md-button--primary } [Cite SecureSDP](research.md#how-to-cite){ .md-button }

</div>
