package ansible

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/antonioalfa22/go-utils/collections"
	"github.com/antonioalfa22/go-utils/io"
	"os"
	"path/filepath"
)

func getVarsAndHosts() ([]string, string) {
	varsfiles, hostlist := getVarsAndHostsFiles()
	qs := crearMenu(varsfiles, hostlist)
	respuestas := struct {
		VarsFile   string
		HostsGroup string
	}{}
	// RESPUESTAS
	_ = survey.Ask(qs, &respuestas)
	vars := collections.Map(io.ReadFile(respuestas.VarsFile), func(x string) string {
		return "    " + x + "\n"
	}).([]string)
	return vars, respuestas.HostsGroup
}

func getVarsAndHostsFiles() ([]string, []string) {
	varsroot := "/etc/egida/vars"
	var varsfiles []string
	err := filepath.Walk(varsroot, func(path string, info os.FileInfo, err error) error {
		if path != "/etc/egida/vars" {
			varsfiles = append(varsfiles, path)
		}
		return nil
	})
	if err != nil {
		fmt.Println("Cant find /etc/egida/vars directory")
	}
	hostlist := io.ReadFile("/etc/egida/hostsgroups")
	return varsfiles, hostlist
}

func crearMenu(varsfiles []string, hostslist []string) []*survey.Question {
	var qs = []*survey.Question{
		{
			Name: "VarsFile",
			Prompt: &survey.Select{
				Message: "Which vars file do you want to use?",
				Options: varsfiles,
				Default: "/etc/egida/vars/vars_template.yml",
			},
		},
		{
			Name: "HostsGroup",
			Prompt: &survey.Select{
				Message: "Which hosts group do you want to use?",
				Options: hostslist,
				Default: "local",
			},
		},
	}
	return qs
}
