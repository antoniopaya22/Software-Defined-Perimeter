package config

import (
	"bufio"
	"fmt"
	"html/template"
	"os"
	"strings"

	"github.com/antonioalfa22/go-utils/collections"
	"github.com/antonioalfa22/go-utils/command"
	"github.com/antonioalfa22/go-utils/io"
)

func AddHostGroup(group string, hostslist []string, connection string) {
	groups := io.ReadFile("/etc/egida/hostsgroups")
	if collections.Find(groups, func(s string) bool { return s == group }) == nil {
		groups = append(groups, group)
		io.WriteFile(groups, "/etc/egida/hostsgroups")
		lines := io.ReadFile("/etc/ansible/hosts")
		g := "[" + group + "]"
		lines = append(lines, g)
		for _, host := range hostslist {
			lines = append(lines, host)
		}
		io.WriteFile(lines, "/etc/ansible/hosts")
	} else {
		fmt.Println("Group " + group + " already exists")
	}
	SetupGroup(group, connection)
}

func SetupGroup(group string, connection string) {
	renderFile(group, connection)
	if connection == "ssh" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("SSH User: ")
		user, _ := reader.ReadString('\n')
		err := command.RunCommandPrintOutput("ansible-playbook", "-u", user, "--ask-pass", "/etc/egida/generated_setup.yml")
		if err != nil {
			fmt.Println("Error on running playbook, Do you have Ansible installed?")
		}
	} else {
		err := command.RunCommandPrintOutput("ansible-playbook", "/etc/egida/generated_setup.yml")
		if err != nil {
			fmt.Println("Error on running playbook, Do you have Ansible installed?")
		}
	}
}

func renderFile(group string, connection string) {
	type Options struct {
		Hosts      string
		Connection string
	}
	options := Options{Hosts: group, Connection: connection}
	tmpl := template.New("PlaybookSetupTemplate")
	tmpl, _ = tmpl.Parse(
		"---\n" +
			"\n" +
			"- name: egida-role-setup\n" +
			"  hosts: {{ .Hosts }}\n" +
			"  connection: {{ .Connection }}\n" +
			"  become: yes\n" +
			"\n" +
			"  roles:\n" +
			"    - antonioalfa22.egida_role_setup\n")
	f, err := os.Create("/etc/egida/generated_setup.yml")
	if err != nil {
		fmt.Println("Error creating playbook: " + err.Error())
	}
	_ = tmpl.Execute(f, options)
}

func RemoveHostGroup(group string) {
	groups := io.ReadFile("/etc/egida/hostsgroups")
	if collections.Find(groups, func(s string) bool { return s == group }) != nil {
		groups = remove(groups, group)
		io.WriteFile(groups, "/etc/egida/hostsgroups")
		var lines []string
		for _, gp := range groups {
			g := "[" + gp + "]"
			lines = append(lines, g)
			hostslist := GetHostsFromGroup(gp)
			for _, host := range hostslist {
				lines = append(lines, host)
			}
		}

		io.WriteFile(lines, "/etc/ansible/hosts")
	} else {
		fmt.Println("Group " + group + " not exists")
	}
}

func GetHostsFromGroup(group string) []string {
	lines := io.ReadFile("/etc/ansible/hosts")
	var hosts []string
	gpStarts := false
	for _, line := range lines {
		if strings.HasPrefix(line, "[") {
			line = strings.Replace(line, "[", "", -1)
			line = strings.Replace(line, "]", "", -1)
			line = strings.Replace(line, " ", "", -1)
			if line == group && !gpStarts {
				gpStarts = true
			} else if gpStarts {
				gpStarts = false
			}
		} else if gpStarts && !strings.HasPrefix(line, "[") {
			hosts = append(hosts, line)
		}
	}
	return hosts
}

func remove(s []string, r string) []string {
	for i, v := range s {
		if v == r {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}
