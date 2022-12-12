package cis

import (
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/antonioalfa22/egida/pkg/ansible"
	"github.com/antonioalfa22/go-utils/collections"
)

func ShowControlsMenu(connection string) {
	var controls []string
	prompt := &survey.MultiSelect{
		Message: "Select CIS Controls:",
		Options: []string{
			"2.-Inventory and Control of Software Assets",
			"3.-Continuous Vulnerability Management",
			"4.-Controlled Use of Administrative Privileges",
			"5.-Secure Configuration for Hardware and Software on Mobile Devices, Laptops, Workstations and Servers",
			"6.-Maintenance, Monitoring and Analysis of Audit Logs",
			"8.-Malware Defenses",
			"9.-Limitation and Control of Network Ports, Protocols and Services",
			"13.-Data Protection",
			"14.-Controlled Access Based on the Need to Know",
			"16.-Account Monitoring and Control",
		},
	}
	_ = survey.AskOne(prompt, &controls)
	temp := collections.Map(controls, func(p string) string { return strings.Split(p, ".")[0] })
	result := getControls(temp.([]string))
	ansible.RunMenuPlaybook(result, connection)
}

func getControls(controls []string) []string {
	var result []string
	all_controls := map[string][]string{
		"2":  {"control_2.6"},
		"3":  {"control_3.4", "control_3.5"},
		"4":  {"control_4.3", "control_4.4", "control_4.5", "control_4.8", "control_4.9"},
		"5":  {"control_5.1", "control_5.5"},
		"6":  {"control_6", "control_6.1", "control_6.2", "control_6.3"},
		"8":  {"control_8.3"},
		"9":  {"control_9.2", "control_9.4"},
		"13": {"control_13"},
		"14": {"control_14.4", "control_14.6", "control_14.9"},
		"16": {"control_16", "control_16.3", "control_16.4", "control_16.5", "control_16.7", "control_16.11", "control_16.13"},
	}
	for _, v := range controls {
		if val, ok := all_controls[v]; ok {
			for _, element := range val {
				result = append(result, element)
			}
		}
	}
	return result
}
