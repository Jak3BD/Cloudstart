package defaults

import "cloudstart/config"

type Defaults struct {
	File        File        `yaml:"file,omitempty"`
	Packages    Packages    `yaml:"packages,omitempty"`
	SSH         SSH         `yaml:"ssh,omitempty"`
	UserOptions UserOptions `yaml:"user_options,omitempty"`

	Users   []config.User `yaml:"users,omitempty"`
	Runcmds []Runcmd      `yaml:"runcmds,omitempty"`
}

type File struct {
	Path string `yaml:"path"`
	Name string `yaml:"name"`
}

type Packages struct {
	Update   bool     `yaml:"update,omitempty"`
	Upgrade  bool     `yaml:"upgrade,omitempty"`
	Packages []string `yaml:"packages,omitempty"`
}

type SSH struct {
	Pwauth         bool `yaml:"pwauth,omitempty"`
	DisableRoot    bool `yaml:"disable_root,omitempty"`
	NoFingerprints bool `yaml:"no_fingerprints,omitempty"`
}

type UserOptions struct {
	Expiredate        string   `yaml:"expiredate,omitempty"`
	Groups            []string `yaml:"groups,omitempty"`
	LockPasswd        bool     `yaml:"lock_passwd,omitempty"`
	Inactive          string   `yaml:"inactive,omitempty"`
	Passwd            string   `yaml:"passwd,omitempty"` // set: "random" or "assign"
	NoCreateHome      bool     `yaml:"no_create_home,omitempty"`
	NoUserGroup       bool     `yaml:"no_user_group,omitempty"`
	NoLogInit         bool     `yaml:"no_log_init,omitempty"`
	SshImportID       []string `yaml:"ssh_import_id,omitempty"`
	SshAuthorizedKeys []string `yaml:"ssh_authorized_keys,omitempty"`
	SshRedirectUser   bool     `yaml:"ssh_redirect_user,omitempty"`
	Sudo              []string `yaml:"sudo,omitempty"`
}

type Runcmd struct {
	Name string   `yaml:"name"`
	Cmds []string `yaml:"cmds"`
}
