package cis

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/antonioalfa22/egida/pkg/ansible"
	"github.com/antonioalfa22/go-utils/collections"
)

func ShowSectionsMenu(connection string) {
	var sections []string
	prompt := &survey.MultiSelect{
		Message: "Select CIS Sections:",
		Options: []string{
			"1.1-Filesystem Configuration",
			"1.2-Console Software Updates",
			"1.3-Configure sudo",
			"1.4-Filesystem Integrity",
			"1.5-Secure Boot Settings",
			"1.6-Additional Process Hardering",
			"1.7-Mandatory Access Control",
			"1.8-Warning Banners",
			"1.9-Updates",

			"2.1-Initd Services",
			"2.2-Special Purpose Services",
			"2.3-Service Clients",

			"3.1-Network Parameters Host",
			"3.2-Network Parameters Host and Router",
			"3.3-TCP Wrappers",
			"3.4-Uncommon Network Protocols",
			"3.5-Firewall Configuration",
			"3.6-Wireless",
			"3.7-Disable IPv6",

			"4.1-Configure System Accounting (auditd)",
			"4.2-Configure logging",
			"4.3-Ensure logrotate is configured",

			"5.1-Configure cron",
			"5.2-SSH Server configuration",
			"5.3-Configure PAM",
			"5.4-User accounts and environment",
			"5.5-Ensure root login is restricted to system console",
			"5.6-Ensure access to the su command is restricted",

			"6.1-System file permisions",
			"6.2-User and Group Settings",
		},
	}
	_ = survey.AskOne(prompt, &sections)
	result := collections.Map(sections, func(p string) string { return "section_" + strings.Split(p, "-")[0] })
	ansible.RunMenuPlaybook(result.([]string), connection)
}
