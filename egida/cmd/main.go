package main

import (
	"fmt"
	"os"

	"github.com/akamensky/argparse"
	"github.com/antonioalfa22/egida/internal/config"
	"github.com/antonioalfa22/egida/internal/dsl"
	"github.com/antonioalfa22/egida/internal/info"
	"github.com/antonioalfa22/egida/internal/menu"
)

var parser *argparse.Parser

func main() {
	if os.Geteuid() != 0 {
		fmt.Println("You're not sudo!")
		return
	}
	parser = argparse.NewParser("Egida CLI", "")
	// Commands
	menuCmd := parser.NewCommand("menu", "")
	addGroupCmd := parser.NewCommand("add-group", "")
	removeGroupCmd := parser.NewCommand("remove-group", "")
	infoCmd := parser.NewCommand("info", "")
	compileCmd := parser.NewCommand("compile", "")
	// Flags
	hostsgroup := parser.String("g", "group",
		&argparse.Options{Required: false, Help: "Host group: -group local"})
	hostslist := parser.StringList("H", "hosts",
		&argparse.Options{Required: false, Help: "List of hosts: -H 192.128.2.1 --hosts localhost -H 129.1.1.1"})
	services := parser.Selector("s", "services", []string{"all", "running", "stopped"},
		&argparse.Options{Required: false, Help: "Services info (all | running | stopped): --services all"})
	packages := parser.Selector("p", "packages", []string{"all"},
		&argparse.Options{Required: false, Help: "Packages info (all): --packages all"})
	hardening := parser.Selector("z", "hardscores", []string{"lynis"},
		&argparse.Options{Required: false, Help: "Hardening scores info (lynis): --hardscores lynis"})
	connection := parser.Selector("c", "connection", []string{"local", "ssh"},
		&argparse.Options{Required: false, Help: "Connection (local | ssh): --connection ssh"})
	file := parser.File("f", "file", os.O_RDWR, 0600,
		&argparse.Options{Required: false, Help: "Aspida File"})

	err := parser.Parse(os.Args)
	if err != nil {
		invalidArgs()
		return
	}
	if menuCmd.Happened() {
		setMenu(*connection)
	} else if compileCmd.Happened() {
		setCompile(*file)
	} else if addGroupCmd.Happened() {
		setAddGroup(*hostsgroup, *hostslist, *connection)
	} else if removeGroupCmd.Happened() {
		setRemoveGroup(*hostsgroup)
	} else if infoCmd.Happened() {
		setInfo(*hostslist, *services, *packages, *hardening)
	} else {
		invalidArgs()
	}
}

// EGIDA OPTIONS
func invalidArgs() {
	fmt.Print(parser.Usage(nil))
}

func setMenu(connection string) {
	if connection != "" {
		menu.SelectHardeningMode(connection)
	} else {
		invalidArgs()
	}
}

func setCompile(file os.File) {
	dsl.ParseFile(file.Name())
}

func setAddGroup(hostsgroup string, hostslist []string, connection string) {
	if hostsgroup != "" && len(hostslist) != 0 {
		config.AddHostGroup(hostsgroup, hostslist, connection)
	} else {
		invalidArgs()
	}
}

func setRemoveGroup(hostsgroup string) {
	if hostsgroup != "" {
		config.RemoveHostGroup(hostsgroup)
	} else {
		invalidArgs()
	}
}

func setInfo(hostslist []string, services string, packages string, hardening string) {
	if len(hostslist) != 0 {
		info.GetWorkerInfo(hostslist, services, packages, hardening)
	} else {
		invalidArgs()
	}
}
