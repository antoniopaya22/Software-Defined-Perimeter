# Research line

This project is part of a line of research on Zero Trust networking and Software-Defined Perimeters (SDP) at the University of Oviedo. Each step builds on the previous one.

| Year | Work | What it contributes | Code |
|---|---|---|---|
| 2022 | **Egida**: *Automated security configuration deployment systems with early error detection* ([Computers & Security 116, 102638](https://doi.org/10.1016/j.cose.2022.102638)) | Security orchestration with a DSL (Aspida) that applies CIS-based hardening controls without breaking legitimate services. | [Egida-Kassandra/egida](https://github.com/Egida-Kassandra/egida) |
| 2024 | **SecureSDP**: *A novel software-defined perimeter implementation for enhanced network security and scalability* ([IJIS 23, 2793–2808](https://doi.org/10.1007/s10207-024-00863-7)) | Bridge-model SDP for scalability plus dynamic per-component hardening with Egida (Lynis +65 %, Chef InSpec +78 %, OpenSCAP +30 % on the controller, as reported in the paper). Earlier version: [SSRN preprint, 2023](https://doi.org/10.2139/ssrn.4326496). | [Software-Defined-Perimeter](https://github.com/antoniopaya22/Software-Defined-Perimeter) (fwknop lab) · [Secure-SDP](https://github.com/antoniopaya22/Secure-SDP) (Go implementation) |
| 2024 | **Apollon**: *A robust defense system against Adversarial Machine Learning attacks in Intrusion Detection Systems* ([Computers & Security 136, 103546](https://doi.org/10.1016/j.cose.2023.103546)) | IDS that uses Multi-Armed Bandits to pick classifiers per input, resisting adversarial evasion. | [apollon](https://github.com/antoniopaya22/apollon) |
| 2025 | **IAM + IDS for SDP**: *Enhancing software-defined perimeters with integrated identity solutions and threat detection for robust zero trust security* ([IJIS 24, 178](https://doi.org/10.1007/s10207-025-01099-9)) | Keycloak IAM (MFA, RBAC, session tokens) and Apollon-based behaviour monitoring inside the SDP; ~75–100 ms extra connection setup, throughput largely unaffected (as reported in the paper). | [iam-ids-sdp](https://github.com/antoniopaya22/iam-ids-sdp) |

Related repositories: [Harpe](https://github.com/antoniopaya22/Harpe) (adversarial traffic that evades the IDS of an SDP, the offensive counterpart of Apollon), [SimpleSDP](https://github.com/antoniopaya22/SimpleSDP) (minimal educational SDP), [SPA-Example](https://github.com/antoniopaya22/SPA-Example) (standalone fwknop SPA lab).

Documentation sites: [Software-Defined-Perimeter](https://antoniopaya22.github.io/Software-Defined-Perimeter/) · [Secure-SDP](https://antoniopaya22.github.io/Secure-SDP/) · [iam-ids-sdp](https://antoniopaya22.github.io/iam-ids-sdp/)
