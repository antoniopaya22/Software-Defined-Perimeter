package cis

import "github.com/antonioalfa22/egida/pkg/ansible"

func ShowAllMenu(connection string) {
	ansible.RunMenuPlaybook([]string{}, connection)
}
