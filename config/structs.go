package config

type Config struct {
	PackageUpdate  bool     `yaml:"package_update,omitempty"`
	PackageUpgrade bool     `yaml:"package_upgrade,omitempty"`
	Packages       []string `yaml:"packages,omitempty"`

	SSHPwauth         bool `yaml:"ssh_pwauth,omitempty"`
	SSHDisableRoot    bool `yaml:"disable_root,omitempty"`
	SSHNoFingerprints bool `yaml:"no_ssh_fingerprints,omitempty"`

	Users []User `yaml:"users"`

	Runcmd []string `yaml:"runcmd,omitempty"`
}

type User struct {
	Name              string   `yaml:"name"`
	Expiredate        string   `yaml:"expiredate,omitempty"`
	Gecos             string   `yaml:"gecos,omitempty"`
	Homedir           string   `yaml:"homedir,omitempty"`
	PrimaryGroup      string   `yaml:"primary_group,omitempty"`
	Groups            string   `yaml:"groups,omitempty"`
	SelinuxUser       string   `yaml:"selinux_user,omitempty"`
	LockPasswd        bool     `yaml:"lock_passwd,omitempty"`
	Inactive          string   `yaml:"inactive,omitempty"`
	Passwd            string   `yaml:"passwd,omitempty"`
	NoCreateHome      bool     `yaml:"no_create_home,omitempty"`
	NoUserGroup       bool     `yaml:"no_user_group,omitempty"`
	NoLogInit         bool     `yaml:"no_log_init,omitempty"`
	SshImportID       []string `yaml:"ssh_import_id,omitempty"`
	SshAuthorizedKeys []string `yaml:"ssh_authorized_keys,omitempty"`
	SshRedirectUser   bool     `yaml:"ssh_redirect_user,omitempty"`
	Sudo              []string `yaml:"sudo,omitempty"`
	System            bool     `yaml:"system,omitempty"`
	Snapuser          string   `yaml:"snapuser,omitempty"`
}

type Metadata struct {
	FilePath  string
	FileName  string
	UserCreds []UserCreds
}

type UserCreds struct {
	Name   string
	Passwd string
}
