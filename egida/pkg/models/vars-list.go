package models

func GetVarsList() []Variable {
	var vars []Variable
	vars = append(vars, Variable{
		Name: "aide_cron",
		Section: "section_1.4",
		Points: []string{"rule_1.4.2"},
		Tags: []string{"aide"},
		Controls: []string{"control_14.9"},
	})
	vars = append(vars, Variable{
		Name: "ufw_ports_allow",
		Section: "section_3.5",
		Points: []string{"rule_3.5.2.1"},
		Tags: []string{"firewall"},
		Controls: []string{"control_9.4"},
	})
	vars = append(vars, Variable{
		Name: "sshd_access",
		Section: "section_5.2",
		Points: []string{"rule_5.2.1","rule_5.2.2","rule_5.2.3","rule_5.2.4","rule_5.2.5","rule_5.2.6","rule_5.2.7","rule_5.2.8","rule_5.2.9","rule_5.2.10","rule_5.2.11","rule_5.2.12","rule_5.2.13","rule_5.2.14","rule_5.2.15","rule_5.2.16","rule_5.2.17","rule_5.2.18","rule_5.2.19","rule_5.2.20","rule_5.2.21","rule_5.2.22","rule_5.2.23","rule_5.2.24"},
		Tags: []string{"sshd"},
		Controls: []string{"control_14.6"},
	})
	return vars
}
