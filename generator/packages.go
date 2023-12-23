package generator

import (
	"cloudstart/config"
	"cloudstart/defaults"
	"cloudstart/utils"
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

type Packages struct {
	Update   bool
	Upgrade  bool
	Packages []string
}

func (p *Packages) Generate() {
	defaultsSet := false
	if p.Update || p.Upgrade || len(p.Packages) > 0 {
		defaultsSet = true
	}

	var setUpdateUpgrade bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Set update/upgrade", defaultsSet),
		Default: defaultsSet,
	}, &setUpdateUpgrade); err != nil {
		panic(err)
	}

	if setUpdateUpgrade {
		p.SetUpdate()
		p.SetUpgrade()
	}

	p.SetPackages()
}

func (p *Packages) SetUpdate() {
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Update packages", p.Update),
		Default: p.Update,
	}, &config.Store.PackageUpdate); err != nil {
		panic(err)
	}
}

func (p *Packages) SetUpgrade() {
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Upgrade packages", p.Upgrade),
		Default: p.Upgrade,
	}, &config.Store.PackageUpgrade); err != nil {
		panic(err)
	}
}

func (p *Packages) SetPackages() {
	if len(p.Packages) > 0 {
		var addDefaultPackages bool
		if err := survey.AskOne(&survey.Confirm{
			Message: utils.PrintDefault(fmt.Sprintf("Add default packages %s", defaults.Store.Packages.Packages), true),
			Default: true,
		}, &addDefaultPackages); err != nil {
			panic(err)
		}

		if addDefaultPackages {
			config.Store.Packages = append(config.Store.Packages, defaults.Store.Packages.Packages...)
		}
	}

	var addPackages bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Add packages", false),
		Default: false,
	}, &addPackages); err != nil {
		panic(err)
	}

	if addPackages {
		var packages string
		if err := survey.AskOne(&survey.Input{
			Message: "Packages (space separated)",
		}, &packages); err != nil {
			panic(err)
		}

		config.Store.Packages = append(config.Store.Packages, strings.Split(packages, " ")...)
	}
}
