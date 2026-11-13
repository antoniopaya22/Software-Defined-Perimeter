# Paper

## SecureSDP (2024)

> **SecureSDP: a novel software-defined perimeter implementation for enhanced network security and scalability**<br>
> Antonio Payá, Vicente García-Díaz, Alberto Gómez<br>
> *International Journal of Information Security* 23(4), 2793–2808, 2024. [doi:10.1007/s10207-024-00863-7](https://doi.org/10.1007/s10207-024-00863-7)

??? quote "Abstract"
    Software-defined perimeters (SDP) revolutionize network security by offering dynamic and adaptive access controls, focusing on user and device authentication to substantially reduce the attack surface. Despite their potential, traditional SDP models grapple with challenges in scalability and component-level security, limiting their effectiveness in complex digital environments. To overcome these limitations, this article introduces SecureSDP, a sophisticated evolution of the SDP framework designed to enhance scalability and bolster security for each network component. SecureSDP stands out for its seamless integration into varied organizational structures, delivering a robust and comprehensive security solution that strengthens the network's defenses across all layers. The key advancement of SecureSDP is evidenced by rigorous experimental evaluations, which demonstrate its superior performance in improving network security and scalability. Specifically, SecureSDP achieved substantial increases in the hardening scores across various tools: 65% in Lynis, 78% in Chef Inspec, and 30% in OpenSCAP for the SDP controller. These enhancements underscore SecureSDP's significant contributions to the field, marking a pivotal step forward in the development of more secure, scalable network infrastructures.

    — Payá et al., *International Journal of Information Security*, 2024.

The paper has two contributions:

- **Scalability with the Bridge model.** Bridged controllers relay authorization requests to a root controller over mTLS, which keeps each organization's existing topology. The [Secure-SDP](https://antoniopaya22.github.io/Secure-SDP/) repository implements this architecture.
- **Per-component dynamic hardening with Egida.** The controller and the Accepting Hosts are hardened with controls that fit each role. This lab provides the fwknop-based SDP that those components run, and [Component hardening](hardening.md) summarizes the reported results.

### Earlier version

The work was first released as a preprint: *Securing the Software-Defined Perimeter Framework with Automated Security Configuration Deployment Systems*, Payá González, Redondo López, García Díaz and Gómez Gómez, SSRN, 2023. [doi:10.2139/ssrn.4326496](https://doi.org/10.2139/ssrn.4326496).

## How to cite

If you use this lab, please cite the published article:

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

GitHub's **Cite this repository** button uses [`CITATION.cff`](https://github.com/antoniopaya22/Software-Defined-Perimeter/blob/main/CITATION.cff), which points to the same article.

## Follow-up work

> **Enhancing software-defined perimeters with integrated identity solutions and threat detection for robust zero trust security**<br>
> Antonio Payá, Vicente García-Díaz, Alberto Gómez<br>
> *International Journal of Information Security* 24(4), 178, 2025. [doi:10.1007/s10207-025-01099-9](https://doi.org/10.1007/s10207-025-01099-9)

This paper adds Keycloak-based identity and access management (MFA, RBAC, session tokens) and Apollon-based intrusion detection to the SDP. Code and lab: [iam-ids-sdp](https://antoniopaya22.github.io/iam-ids-sdp/).

```bibtex
@article{paya2025enhancing,
  title   = {Enhancing software-defined perimeters with integrated identity solutions and threat detection for robust zero trust security},
  author  = {Pay{\'a}, Antonio and Garc{\'i}a-D{\'i}az, Vicente and G{\'o}mez, Alberto},
  journal = {International Journal of Information Security},
  volume  = {24},
  number  = {4},
  pages   = {178},
  year    = {2025},
  doi     = {10.1007/s10207-025-01099-9}
}
```

See the [research line](research-line.md) for the full sequence of publications and repositories.
