package ansible

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"text/template"

	"github.com/antonioalfa22/go-utils/command"
)

func RunMenuPlaybook(tags []string, connection string) {
	vars, hosts := getVarsAndHosts()
	renderFile(vars, hosts, connection)
	runPlaybook(tags, connection)
}

func RunDSLPlaybook(tags []string, vars []string, hosts string, connection string) {
	renderFile(vars, hosts, connection)
	runPlaybook(tags, connection)
}

func runPlaybook(tags []string, connection string) {
	if connection == "ssh" {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("SSH User: ")
		user, _ := reader.ReadString('\n')
		err := command.RunCommandPrintOutput("ansible-playbook", "-u", user, "--ask-pass", "/etc/egida/generated.yml",
			"--tags="+getTags(tags))
		if err != nil {
			fmt.Println("Error on running playbook, Do you have Ansible installed?")
		}
	} else {
		if len(tags) == 0 {
			err := command.RunCommandPrintOutput("ansible-playbook", "/etc/egida/generated.yml")
			if err != nil {
				fmt.Println("Error on running playbook, Do you have Ansible installed?")
			}
		} else {
			err := command.RunCommandPrintOutput("ansible-playbook", "/etc/egida/generated.yml",
				"--tags="+getTags(tags))
			if err != nil {
				fmt.Println("Error on running playbook, Do you have Ansible installed?")
			}
		}
	}
}

func renderFile(vars []string, hosts string, connection string) {
	type Options struct {
		Hosts      string
		Connection string
		Vars       []string
	}
	options := Options{Hosts: hosts, Vars: vars, Connection: connection}
	tmpl := template.New("PlaybookTemplate")
	tmpl, _ = tmpl.Parse(
		"---\n" +
			"\n" +
			"- name: Harden Server\n" +
			"  hosts: {{ .Hosts }}\n" +
			"  connection: {{ .Connection }}\n" +
			"  become: yes\n" +
			"  vars:\n" +
			"{{ range .Vars }}{{ . }}{{ end }}\n" +
			"\n" +
			"  roles:\n" +
			"    - antonioalfa22.egida_role_cis\n")
	f, err := os.Create("/etc/egida/generated.yml")
	if err != nil {
		fmt.Println("Error creating playbook: " + err.Error())
	}
	_ = tmpl.Execute(f, options)
}

func getTags(tags []string) string {
	return strings.Join(tags[:], ",")
}
