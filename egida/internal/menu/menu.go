package menu

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/antonioalfa22/egida/internal/menu/cis"
)

func SelectHardeningMode(connection string) {
	// LOGO
	fmt.Print("\033[2J") //Clear screen
	printLogo()
	// MENU
	qs := crearMenu()
	respuestas := struct {
		HardMode string
		Option   string
	}{}
	// RESPUESTAS
	_ = survey.Ask(qs, &respuestas)
	if respuestas.HardMode == "CIS Benchmarks" {
		switch respuestas.Option {
		case "Select CIS Controls":
			cis.ShowControlsMenu(connection)
		case "Select CIS Benchmarks Points":
			cis.ShowPointsMenu(connection)
		case "Select CIS Sections":
			cis.ShowSectionsMenu(connection)
		case "Run all CIS Benchmarks":
			cis.ShowAllMenu(connection)
		}
	}
}

func crearMenu() []*survey.Question {
	var qs = []*survey.Question{
		{
			Name: "HardMode",
			Prompt: &survey.Select{
				Message: "Select server hardening mode",
				Options: []string{"CIS Benchmarks", "LAMP Server (soon)", "LEMP Server(soon)"},
				Default: "CIS Benchmarks",
			},
		},
		{
			Name: "Option",
			Prompt: &survey.Select{
				Message: "How do you want to do the hardening?",
				Options: []string{
					"Select CIS Controls",
					"Select CIS Benchmarks Points",
					"Select CIS Sections",
					"Run all CIS Benchmarks",
				},
				Default: "Select CIS Controls",
			},
		},
	}
	return qs
}

func printLogo() {
	fmt.Println()
	fmt.Println("\t@@@@@@@@@@@@@&%%%%((((((((((                        ((((((((((%%%%%@@@@@@@@@@@@@")
	fmt.Println("\t@@@@@@@@@@@%%%%((((((((   *******               ******,  ((((((((#%%%%@@@@@@@@@@")
	fmt.Println("\t@@@@@@@@@%%%#(((((((    ***    .***.         ****     ***   ((((((((%%%%@@@@@@@@")
	fmt.Println("\t@@@@@@@%%%(((((((      ***      ***, .*****  ***       ***     (((((((%%%%@@@@@@")
	fmt.Println("\t@@@@@%%%%((((((        ***        .   **.**            **        (((((((%%%%@@@@")
	fmt.Println("\t@@@@%%%((((((            ******************************.,          ((((((%%%%@@@")
	fmt.Println("\t@@@%%%((((((           %%%%%%**%%%%***,,*,,**#%%%(**%%%%%%          ((((((%%%%@@")
	fmt.Println("\t@@%%%((((((          *%%%%%%%%%******,,%%#,,******%%%%%%%%%          ((((((%%%%@")
	fmt.Println("\t@%%%((((((         %%**%%%%***********%%%%%***********%%%%**%%        ((((((%%%@")
	fmt.Println("\t@%%%(((((         %%%*******%****%%%%%%%%%%%%%%%****%******%%%,        ((((((%%%")
	fmt.Println("\t@%%((((((       ...%%%%%%%*****%%%%%%%%%%%%%%%%%%******%%%%%%%...      ((((((%%%")
	fmt.Println("\t%%%((((((        #%*****%%%***%%%%%   %%%%   %%%%%***%%%*****%#        ((((((%%%")
	fmt.Println("\t%%%((((((      #**%%%%%%*******%%%%%%%%%%%%%%%%%%%*******%%%%%%**#     ((((((%%%")
	fmt.Println("\t@%%((((((        %%%**%********%%%%%%%%% %%%%%%%%%********%**%%%       ((((((%%%")
	fmt.Println("\t@%%%(((((         %%%%%/%%%%****%%%%%%%%%%%%%%%%%****%%%%/%%%%%        ((((((%%%")
	fmt.Println("\t@%%%((((((         %(**%%/*******%%%%#     %%%%%*******/%%**%%        ((((((%%%@")
	fmt.Println("\t@@%%%((((((          %%%%%%%%%****%%%%%%%%%%%%****%%%%%%%%%%         ((((((%%%%@")
	fmt.Println("\t@@@%%%((((((          %%*********%#**%%%%%%%**#%********%%          ((((((%%%%@@")
	fmt.Println("\t@@@@%%%((((((   ***   . . #%%%***/%%%%****(%%%%***(%%% ..    ***   ((((((%%%%@@@")
	fmt.Println("\t@@@@@%%%%((((((             ____________________  ___            (((((((%%%@@@@@")
	fmt.Println("\t@@@@@@@%%%%((((((          / ____/ ____/  _/ __ \\/   |         (((((((%%%%@@@@@@")
	fmt.Println("\t@@@@@@@@@%%%%((((((       / __/ / / __ / // / / / /| |      ((((((((%%%%@@@@@@@@")
	fmt.Println("\t@@@@@@@@@@%%%%(((((((    / /___/ /_/ // // /_/ / ___ |   ((((((((%%%%%@@@@@@@@@@")
	fmt.Println("\t@@@@@@@@@%%%%((((((((   /_____/\\____/___/_____/_/  |_|   (((((((%%%%%@@@@@@@@@@@")
	fmt.Println("                              ")
	fmt.Println()
}
