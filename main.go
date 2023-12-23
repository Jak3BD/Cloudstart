package main

import (
	"cloudstart/config"
	"cloudstart/defaults"
	"cloudstart/generator"
	"cloudstart/version"
	"fmt"
)

func init() {
	fmt.Printf(`
░█▀▀░█░░░█▀█░█░█░█▀▄░█▀▀░▀█▀░█▀█░█▀▄░▀█▀
░█░░░█░░░█░█░█░█░█░█░▀▀█░░█░░█▀█░█▀▄░░█░
░▀▀▀░▀▀▀░▀▀▀░▀▀▀░▀▀░░▀▀▀░░▀░░▀░▀░▀░▀░░▀░ v%s (%s/%s)
%s`, version.Version, version.OS, version.Arch, "\n")

	fmt.Println("loading default config...")
	defaults.Init()

	fmt.Println()
}

func main() {
	file := generator.File{
		Path: defaults.Store.File.Path,
		Name: defaults.Store.File.Name,
	}
	file.Generate()

	packages := generator.Packages{
		Update:   defaults.Store.Packages.Update,
		Upgrade:  defaults.Store.Packages.Upgrade,
		Packages: defaults.Store.Packages.Packages,
	}
	packages.Generate()

	ssh := generator.SSH{
		Pwauth:         defaults.Store.SSH.Pwauth,
		DisableRoot:    defaults.Store.SSH.DisableRoot,
		NoFingerprints: defaults.Store.SSH.NoFingerprints,
	}
	ssh.Generate()

	users := generator.Users{
		Users:       defaults.Store.Users,
		UserOptions: defaults.Store.UserOptions,
	}
	users.Generate()

	runcmd := generator.Runcmd{
		Cmds: defaults.Store.Runcmds,
	}
	runcmd.Generate()

	config.Write()
}
