package generator

import (
	"cloudstart/config"
	"cloudstart/defaults"
	"cloudstart/version"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type Runcmd struct {
	Cmds []defaults.Runcmd
}

func (r *Runcmd) Generate() {
	msg := fmt.Sprintf("echo '####################_Cloudstart_v%s_####################'", version.Version)
	config.Store.Runcmd = append(config.Store.Runcmd, msg)

	if len(r.Cmds) <= 0 {
		fmt.Println("No commands found, can be defined in '$HOME/.config/cloudstart/defaults.yml'")
		return
	}

	r.SetCmds()
}

func (r *Runcmd) SetCmds() {
	var cmdNames []string
	for _, cmd := range r.Cmds {
		cmdNames = append(cmdNames, cmd.Name)
	}

	var selected []string
	if err := survey.AskOne(&survey.MultiSelect{
		Message: "Select commands to run",
		Options: cmdNames,
	}, &selected); err != nil {
		panic(err)
	}

	var cmds []string
	for _, cmd := range r.Cmds {
		for _, name := range selected {
			if cmd.Name == name {
				var addCmds []string
				addCmds = append(addCmds, fmt.Sprintf("echo 'Execute script %s'", cmd.Name))
				addCmds = append(addCmds, cmd.Cmds...)
				cmds = append(cmds, addCmds...)
			}
		}
	}

	config.Store.Runcmd = append(config.Store.Runcmd, cmds...)
}
