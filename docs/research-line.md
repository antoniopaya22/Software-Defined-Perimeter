# Research line

This project is part of a line of research on Zero Trust networking and Software-Defined Perimeters (SDP) at the University of Oviedo. Each step builds on the previous one.

<ol class="sdp-timeline">
  <li>
    <span class="sdp-timeline__year">2022 · COMPUTERS &amp; SECURITY 116, 102638</span>
    <span class="sdp-timeline__title">Egida: automated security configuration deployment with early error detection</span>
    <p>Security orchestration with a DSL (Aspida) that applies CIS-based hardening controls without breaking legitimate services.</p>
    <div class="sdp-timeline__links">
      <a href="https://doi.org/10.1016/j.cose.2022.102638">Paper</a>
      <a href="https://github.com/Egida-Kassandra/egida">Egida-Kassandra/egida</a>
    </div>
  </li>
  <li>
    <span class="sdp-timeline__year">2024 · IJIS 23, 2793–2808</span>
    <span class="sdp-timeline__title">SecureSDP: a novel software-defined perimeter implementation for enhanced network security and scalability</span>
    <p>A Bridge-model SDP for scalability, plus dynamic per-component hardening with Egida. The paper reports Lynis +65 %, Chef InSpec +78 % and OpenSCAP +30 % on the controller. An earlier version was released as an <a href="https://doi.org/10.2139/ssrn.4326496">SSRN preprint</a> in 2023.</p>
    <div class="sdp-timeline__links">
      <a href="https://doi.org/10.1007/s10207-024-00863-7">Paper</a>
      <a href="https://antoniopaya22.github.io/Software-Defined-Perimeter/">Software-Defined-Perimeter (fwknop lab)</a>
      <a href="https://antoniopaya22.github.io/Secure-SDP/">Secure-SDP (Go implementation)</a>
    </div>
  </li>
  <li>
    <span class="sdp-timeline__year">2024 · COMPUTERS &amp; SECURITY 136, 103546</span>
    <span class="sdp-timeline__title">Apollon: a robust defense system against adversarial machine learning attacks in IDS</span>
    <p>An IDS that uses Multi-Armed Bandits to choose the classifier for each input, which makes adversarial evasion harder.</p>
    <div class="sdp-timeline__links">
      <a href="https://doi.org/10.1016/j.cose.2023.103546">Paper</a>
      <a href="https://github.com/antoniopaya22/apollon">apollon</a>
    </div>
  </li>
  <li>
    <span class="sdp-timeline__year">2025 · IJIS 24, 178</span>
    <span class="sdp-timeline__title">Enhancing software-defined perimeters with integrated identity solutions and threat detection</span>
    <p>Keycloak IAM (MFA, RBAC, session tokens) and Apollon-based behaviour monitoring inside the SDP. The paper reports about 75–100 ms of extra connection setup, with throughput largely unaffected.</p>
    <div class="sdp-timeline__links">
      <a href="https://doi.org/10.1007/s10207-025-01099-9">Paper</a>
      <a href="https://antoniopaya22.github.io/iam-ids-sdp/">iam-ids-sdp</a>
    </div>
  </li>
</ol>

## Related repositories

| Repository | Description |
|---|---|
| [Harpe](https://github.com/antoniopaya22/Harpe) | Adversarial traffic generation (WGAN-GP) that evades the IDS of an SDP. It is the offensive counterpart of Apollon. |
| [SimpleSDP](https://github.com/antoniopaya22/SimpleSDP) | Minimal educational SDP in Python. |
| [SPA-Example](https://github.com/antoniopaya22/SPA-Example) | Standalone fwknop Single Packet Authorization lab. |

## Documentation sites

[Software-Defined-Perimeter](https://antoniopaya22.github.io/Software-Defined-Perimeter/){ .md-button } [Secure-SDP](https://antoniopaya22.github.io/Secure-SDP/){ .md-button } [iam-ids-sdp](https://antoniopaya22.github.io/iam-ids-sdp/){ .md-button }
