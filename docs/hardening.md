# Component hardening

An SDP hides services from unauthorized clients. The controller, the gateway and the hosts behind it are still ordinary systems, though, and a compromise of any of them weakens the whole perimeter. SecureSDP extends Zero Trust down to each component: every element is hardened with security controls chosen for its role.

## Dynamic hardening with Egida

SecureSDP uses [Egida](https://github.com/Egida-Kassandra/egida) ([Payá et al., Computers & Security, 2022](https://doi.org/10.1016/j.cose.2022.102638)) to apply hardening controls based on the CIS Benchmarks. Egida suits an SDP for three reasons:

- **Master–slave model.** Egida's master can be the SDP controller, which already knows the role of every Accepting Host. It can order each host to apply the controls that its services allow.
- **Detects applied controls.** Egida knows which controls are already in place, so it can apply and roll back controls quickly as the services of a host change.
- **A DSL (Aspida) with conditions.** One policy file can choose, per host, only the controls that do not interfere with legitimate services.

!!! example "Illustrative Aspida policy for a gateway"
    A gateway relies on its own iptables rules managed by fwknopd. The policy below therefore leaves out the CIS firewall section and applies filesystem, integrity, auditing and SSH controls:

    ```text
    MAIN : {
        name: "SDP gateway hardening";
        connection: ssh;
        description: "CIS controls compatible with fwknopd";
    }

    HOST : "sdp-gateways";

    TASKS : {
        sections: ["1.1", "1.4", "4.1", "5.2"];
    }
    ```

    Run it from the Egida master with `egida compile -f gateway.aspida`. See the [Egida documentation](https://github.com/Egida-Kassandra/egida) for installation, host groups and variables.

!!! note "Containers in this lab are not hardened"
    Egida applies controls to full hosts through Ansible: kernel parameters, mount options, auditd, PAM, SSH and so on. Most of these do not apply to Docker containers, so this lab does not run Egida. To reproduce the hardening experiments, deploy the components on Debian/Ubuntu virtual machines and run Egida against them.

## Results reported in the paper

The SecureSDP paper audited the components with three tools, first with the default configuration and then after applying the Egida controls. The **values below are as reported in the paper** ([Tables 2 and 3](https://doi.org/10.1007/s10207-024-00863-7)); they were not regenerated with this lab.

| Component | Tool | Before | After | Improvement |
|---|---|---|---|---|
| SDP Controller | Lynis | 52.00 % | 86.00 % | ×1.65 |
| SDP Controller | Chef InSpec | 36.17 % | 64.25 % | ×1.78 |
| SDP Controller | OpenSCAP | 63.52 % | 82.94 % | ×1.30 |
| SDP Gateway (AH) | Lynis | 54.00 % | 88.00 % | ×1.63 |

According to the paper, the hardened systems had no safety-critical test failures, only warnings. The InSpec and OpenSCAP results of the gateway were the same as the controller's.

## Auditing your own deployment

The paper used these tools and profiles:

=== "Lynis"

    ```bash
    sudo apt-get install -y lynis
    sudo lynis audit system --quick
    # The "Hardening index" line at the end is the score.
    ```

    Warnings and suggestions count as failures; skipped tests are ignored.

=== "Chef InSpec"

    ```bash
    inspec exec https://github.com/dev-sec/cis-dil-benchmark
    ```

    Profile: CIS Distribution Independent Linux Benchmark.

=== "OpenSCAP"

    ```bash
    sudo apt-get install -y openscap-scanner ssg-debderived
    sudo oscap xccdf eval \
      --profile xccdf_org.ssgproject.content_profile_cis_level1_server \
      --report report.html \
      /usr/share/xml/scap/ssg/content/ssg-ubuntu2004-ds.xml
    ```

    Profile: CIS Level 1 Server for Ubuntu 20.04 LTS. Use the data stream that matches your distribution.

!!! tip
    A score of 100 % is not the goal: some controls would break legitimate functionality. Egida's contribution is choosing the controls that raise the score **without** disrupting the services each component provides.
