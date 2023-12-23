package generator

import (
	"cloudstart/config"
	"cloudstart/defaults"
	"cloudstart/utils"
	"fmt"

	"github.com/AlecAivazis/survey/v2"
)

type Users struct {
	Users       []config.User
	UserOptions defaults.UserOptions
}

func (u *Users) Generate() {
	if len(u.Users) > 0 {
		var setPredefinedUser bool
		if err := survey.AskOne(&survey.Confirm{
			Message: utils.PrintDefault("Set predefined user", true),
			Default: true,
		}, &setPredefinedUser); err != nil {
			panic(err)
		}

		if setPredefinedUser {
			u.SetPredefinedUser()
		}
	}

	var addUsers bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Add users", true),
		Default: true,
	}, &addUsers); err != nil {
		panic(err)
	}

	if addUsers {
		addUser := AddUser{UserOptions: u.UserOptions}
		addUser.Generate()
		config.Store.Users = append(config.Store.Users, addUser.Users...)
	}
}

func (u *Users) SetPredefinedUser() {
	var userNames []string
	for _, user := range u.Users {
		userNames = append(userNames, user.Name)
	}

	var selected []string
	if err := survey.AskOne(&survey.MultiSelect{
		Message: "Select predefined user",
		Options: userNames,
	}, &selected); err != nil {
		panic(err)
	}

	var users []config.User
	for _, user := range u.Users {
		for _, name := range selected {
			if user.Name == name {
				var password string

				if user.Passwd == "assign" {
					if err := survey.AskOne(&survey.Password{
						Message: fmt.Sprintf("Set password for %s:", user.Name),
					}, &password); err != nil {
						panic(err)
					}
				} else {
					password = utils.CreateRandomPassword()
				}

				config.Meta.UserCreds = append(config.Meta.UserCreds, config.UserCreds{
					Name:   user.Name,
					Passwd: password,
				})

				user.Passwd = utils.HashPassword(password)
				users = append(users, user)
			}
		}
	}

	config.Store.Users = append(config.Store.Users, users...)
}
