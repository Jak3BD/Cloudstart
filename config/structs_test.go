package config_test

import (
	"cloudstart/config"
	"reflect"
	"testing"
)

type Test struct {
	Name string
	Type reflect.Kind
	Tag  string
}

func runTests(t *testing.T, reflection reflect.Type, tests []Test) {
	for i := 0; i < reflection.NumField(); i++ {
		field := reflection.Field(i)
		name := field.Name

		t.Run(name, func(t *testing.T) {
			test, ok := checkExists(name, tests)
			if !ok {
				t.Errorf("field '%s' not found", name)
				t.FailNow()
			}

			if field.Type.Kind() != test.Type {
				t.Errorf("field '%s' type '%s' != '%s'", name, field.Type.Kind(), test.Type)
			}

			if tag := string(field.Tag); tag != test.Tag {
				t.Errorf("field '%s' tag '%s' != '%s'", name, tag, test.Tag)
			}
		})
	}
}

func checkExists(name string, tests []Test) (Test, bool) {
	for _, test := range tests {
		if test.Name == name {
			return test, true
		}
	}

	return Test{}, false
}

func TestConfig(t *testing.T) {
	tests := []Test{
		{"PackageUpdate", reflect.Bool, `yaml:"package_update,omitempty"`},
		{"PackageUpgrade", reflect.Bool, `yaml:"package_upgrade,omitempty"`},
		{"Packages", reflect.Slice, `yaml:"packages,omitempty"`},

		{"SSHPwauth", reflect.Bool, `yaml:"ssh_pwauth,omitempty"`},
		{"SSHDisableRoot", reflect.Bool, `yaml:"disable_root,omitempty"`},
		{"SSHNoFingerprints", reflect.Bool, `yaml:"no_ssh_fingerprints,omitempty"`},

		{"Users", reflect.Slice, `yaml:"users"`},

		{"Runcmd", reflect.Slice, `yaml:"runcmd,omitempty"`},
	}

	runTests(t, reflect.TypeOf(config.Config{}), tests)
}

func TestUser(t *testing.T) {
	tests := []Test{
		{"Name", reflect.String, `yaml:"name"`},
		{"Expiredate", reflect.String, `yaml:"expiredate,omitempty"`},
		{"Gecos", reflect.String, `yaml:"gecos,omitempty"`},
		{"Homedir", reflect.String, `yaml:"homedir,omitempty"`},
		{"PrimaryGroup", reflect.String, `yaml:"primary_group,omitempty"`},
		{"Groups", reflect.String, `yaml:"groups,omitempty"`},
		{"SelinuxUser", reflect.String, `yaml:"selinux_user,omitempty"`},
		{"LockPasswd", reflect.Bool, `yaml:"lock_passwd,omitempty"`},
		{"Inactive", reflect.String, `yaml:"inactive,omitempty"`},
		{"Passwd", reflect.String, `yaml:"passwd,omitempty"`},
		{"NoCreateHome", reflect.Bool, `yaml:"no_create_home,omitempty"`},
		{"NoUserGroup", reflect.Bool, `yaml:"no_user_group,omitempty"`},
		{"NoLogInit", reflect.Bool, `yaml:"no_log_init,omitempty"`},
		{"SshImportID", reflect.Slice, `yaml:"ssh_import_id,omitempty"`},
		{"SshAuthorizedKeys", reflect.Slice, `yaml:"ssh_authorized_keys,omitempty"`},
		{"SshRedirectUser", reflect.Bool, `yaml:"ssh_redirect_user,omitempty"`},
		{"Sudo", reflect.Slice, `yaml:"sudo,omitempty"`},
		{"System", reflect.Bool, `yaml:"system,omitempty"`},
		{"Snapuser", reflect.String, `yaml:"snapuser,omitempty"`},
	}

	runTests(t, reflect.TypeOf(config.User{}), tests)
}

func TestMetadata(t *testing.T) {
	tests := []Test{
		{"FilePath", reflect.String, ""},
		{"FileName", reflect.String, ""},
		{"UserCreds", reflect.Slice, ""},
	}

	runTests(t, reflect.TypeOf(config.Metadata{}), tests)
}

func TestUserCreds(t *testing.T) {
	tests := []Test{
		{"Name", reflect.String, ""},
		{"Passwd", reflect.String, ""},
	}

	runTests(t, reflect.TypeOf(config.UserCreds{}), tests)
}
