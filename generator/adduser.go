package generator

import (
	"cloudstart/config"
	"cloudstart/defaults"
	"cloudstart/utils"
	"strings"

	"github.com/AlecAivazis/survey/v2"
)

type AddUser struct {
	UserOptions defaults.UserOptions
	Users       []config.User
}

func (a *AddUser) Generate() {
	var user config.User

	if err := survey.AskOne(&survey.Input{
		Message: "Username",
	}, &user.Name); err != nil {
		panic(err)
	}

	a.SetExpiredate(&user)
	a.SetGecos(&user)
	a.SetHomedir(&user)
	a.SetPrimaryGroup(&user)
	a.SetGroups(&user)
	a.SetSelinuxUser(&user)
	a.SetLockPasswd(&user)
	a.SetInactive(&user)
	a.SetPasswd(&user)
	a.SetNoCreateHome(&user)
	a.SetNoUserGroup(&user)
	a.SetNoLogInit(&user)
	a.SetSshImportID(&user)
	a.SetSshAuthorizedKeys(&user)
	a.SetSshRedirectUser(&user)
	a.SetSudo(&user)
	a.SetSystem(&user)
	a.SetSnapuser(&user)

	a.Users = append(a.Users, user)

	var addMore bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Add more users", false),
		Default: false,
	}, &addMore); err != nil {
		panic(err)
	}

	if addMore {
		a.Generate()
	}
}

func (a *AddUser) SetExpiredate(user *config.User) {
	defaultExpiredate := false
	if a.UserOptions.Expiredate != "" {
		defaultExpiredate = true
	}

	var setExpiredate bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Set expiredate", defaultExpiredate),
		Default: defaultExpiredate,
	}, &setExpiredate); err != nil {
		panic(err)
	}

	if setExpiredate {
		if err := survey.AskOne(&survey.Input{
			Message: utils.PrintDefault("Expiredate", a.UserOptions.Expiredate),
			Default: a.UserOptions.Expiredate,
		}, &user.Expiredate); err != nil {
			panic(err)
		}
	}
}

func (a *AddUser) SetGecos(user *config.User) {
	if err := survey.AskOne(&survey.Input{
		Message: "Gecos",
	}, &user.Gecos); err != nil {
		panic(err)
	}
}

func (a *AddUser) SetHomedir(user *config.User) {
	if err := survey.AskOne(&survey.Input{
		Message: "Homedir",
	}, &user.Homedir); err != nil {
		panic(err)
	}
}

func (a *AddUser) SetPrimaryGroup(user *config.User) {
	if err := survey.AskOne(&survey.Input{
		Message: "Primary group",
	}, &user.PrimaryGroup); err != nil {
		panic(err)
	}
}

func (a *AddUser) SetGroups(user *config.User) {
	var groups []string
	if len(a.UserOptions.Groups) > 0 {
		if err := survey.AskOne(&survey.MultiSelect{
			Message: "Set default groups",
			Options: a.UserOptions.Groups,
		}, &groups); err != nil {
			panic(err)
		}
	}

	var addGroups string
	if err := survey.AskOne(&survey.Input{
		Message: "Add groups (space separated)",
	}, &addGroups); err != nil {
		panic(err)
	}

	if addGroups != "" {
		groups = append(groups, strings.Split(addGroups, " ")...)
	}

	user.Groups = strings.Join(groups, ", ")
}

func (a *AddUser) SetSelinuxUser(user *config.User) {
	if err := survey.AskOne(&survey.Input{
		Message: "Selinux user",
	}, &user.SelinuxUser); err != nil {
		panic(err)
	}
}

func (a *AddUser) SetLockPasswd(user *config.User) {
	var lockPasswd bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Lock passwd", a.UserOptions.LockPasswd),
		Default: a.UserOptions.LockPasswd,
	}, &lockPasswd); err != nil {
		panic(err)
	}

	user.LockPasswd = lockPasswd
}

func (a *AddUser) SetInactive(user *config.User) {
	if err := survey.AskOne(&survey.Input{
		Message: utils.PrintDefault("Inactive", a.UserOptions.Inactive),
		Default: a.UserOptions.Inactive,
	}, &user.Inactive); err != nil {
		panic(err)
	}
}

func (a *AddUser) SetPasswd(user *config.User) {
	setPassword := true
	if a.UserOptions.Passwd == "random" {
		setPassword = false
	}

	var setPasswd bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Set password else random", setPassword),
		Default: setPassword,
	}, &setPasswd); err != nil {
		panic(err)
	}

	var password string
	if setPasswd {
		if err := survey.AskOne(&survey.Password{
			Message: "Password",
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
}

func (a *AddUser) SetNoCreateHome(user *config.User) {
	var noCreateHome bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("No create home", a.UserOptions.NoCreateHome),
		Default: a.UserOptions.NoCreateHome,
	}, &noCreateHome); err != nil {
		panic(err)
	}

	user.NoCreateHome = noCreateHome
}

func (a *AddUser) SetNoUserGroup(user *config.User) {
	var noUserGroup bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("No user group", a.UserOptions.NoUserGroup),
		Default: a.UserOptions.NoUserGroup,
	}, &noUserGroup); err != nil {
		panic(err)
	}

	user.NoUserGroup = noUserGroup
}

func (a *AddUser) SetNoLogInit(user *config.User) {
	var noLogInit bool
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("No log init", a.UserOptions.NoLogInit),
		Default: a.UserOptions.NoLogInit,
	}, &noLogInit); err != nil {
		panic(err)
	}

	user.NoLogInit = noLogInit
}

func (a *AddUser) SetSshImportID(user *config.User) {
	var sshImportID []string
	if len(a.UserOptions.SshImportID) > 0 {
		if err := survey.AskOne(&survey.MultiSelect{
			Message: "Set default ssh import id",
			Options: a.UserOptions.SshImportID,
		}, &sshImportID); err != nil {
			panic(err)
		}
	}

	var addSshImportID string
	if err := survey.AskOne(&survey.Input{
		Message: "Add ssh import id (space separated)",
	}, &addSshImportID); err != nil {
		panic(err)
	}

	if addSshImportID != "" {
		sshImportID = append(sshImportID, strings.Split(addSshImportID, " ")...)
	}

	user.SshImportID = sshImportID
}

func (a *AddUser) SetSshAuthorizedKeys(user *config.User) {
	var sshAuthorizedKeys []string
	if len(a.UserOptions.SshAuthorizedKeys) > 0 {
		if err := survey.AskOne(&survey.MultiSelect{
			Message: "Set default ssh authorized keys",
			Options: a.UserOptions.SshAuthorizedKeys,
		}, &sshAuthorizedKeys); err != nil {
			panic(err)
		}
	}

	var addSshAuthorizedKeys string
	if err := survey.AskOne(&survey.Input{
		Message: "Add ssh authorized keys (space separated)",
	}, &addSshAuthorizedKeys); err != nil {
		panic(err)
	}

	if addSshAuthorizedKeys != "" {
		sshAuthorizedKeys = append(sshAuthorizedKeys, strings.Split(addSshAuthorizedKeys, " ")...)
	}

	user.SshAuthorizedKeys = sshAuthorizedKeys
}

func (a *AddUser) SetSshRedirectUser(user *config.User) {
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("Ssh redirect user", a.UserOptions.SshRedirectUser),
		Default: a.UserOptions.SshRedirectUser,
	}, &user.SshRedirectUser); err != nil {
		panic(err)
	}
}

func (a *AddUser) SetSudo(user *config.User) {
	var sudo []string
	if len(a.UserOptions.Sudo) > 0 {
		if err := survey.AskOne(&survey.MultiSelect{
			Message: "Set default sudo",
			Options: a.UserOptions.Sudo,
		}, &sudo); err != nil {
			panic(err)
		}
	}

	var addSudo string
	if err := survey.AskOne(&survey.Input{
		Message: "Add sudo (space separated)",
	}, &addSudo); err != nil {
		panic(err)
	}

	if addSudo != "" {
		sudo = append(sudo, strings.Split(addSudo, " ")...)
	}

	user.Sudo = sudo
}

func (a *AddUser) SetSystem(user *config.User) {
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("System", false),
		Default: false,
	}, &user.System); err != nil {
		panic(err)
	}
}

func (a *AddUser) SetSnapuser(user *config.User) {
	if err := survey.AskOne(&survey.Input{
		Message: "Snapuser",
	}, &user.Snapuser); err != nil {
		panic(err)
	}
}
