package defaults_test

import (
	"cloudstart/defaults"
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

func TestDefaults(t *testing.T) {
	tests := []Test{
		{"File", reflect.Struct, `yaml:"file,omitempty"`},
		{"Packages", reflect.Struct, `yaml:"packages,omitempty"`},
		{"SSH", reflect.Struct, `yaml:"ssh,omitempty"`},
		{"UserOptions", reflect.Struct, `yaml:"user_options,omitempty"`},

		{"Presets", reflect.Slice, `yaml:"presets,omitempty"`},
		{"Users", reflect.Slice, `yaml:"users,omitempty"`},
		{"Runcmds", reflect.Slice, `yaml:"runcmds,omitempty"`},
	}

	runTests(t, reflect.TypeOf(defaults.Defaults{}), tests)
}

func TestFile(t *testing.T) {
	tests := []Test{
		{"Path", reflect.String, `yaml:"path"`},
		{"Name", reflect.String, `yaml:"name"`},
	}

	runTests(t, reflect.TypeOf(defaults.File{}), tests)
}

func TestPackages(t *testing.T) {
	tests := []Test{
		{"Update", reflect.Bool, `yaml:"update,omitempty"`},
		{"Upgrade", reflect.Bool, `yaml:"upgrade,omitempty"`},
		{"Packages", reflect.Slice, `yaml:"packages,omitempty"`},
	}

	runTests(t, reflect.TypeOf(defaults.Packages{}), tests)
}

func TestSSH(t *testing.T) {
	tests := []Test{
		{"Pwauth", reflect.Bool, `yaml:"pwauth,omitempty"`},
		{"DisableRoot", reflect.Bool, `yaml:"disable_root,omitempty"`},
		{"NoFingerprints", reflect.Bool, `yaml:"no_fingerprints,omitempty"`},
	}

	runTests(t, reflect.TypeOf(defaults.SSH{}), tests)
}

func TestUserOptions(t *testing.T) {
	tests := []Test{
		{"Expiredate", reflect.String, `yaml:"expiredate,omitempty"`},
		{"Groups", reflect.Slice, `yaml:"groups,omitempty"`},
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
	}

	runTests(t, reflect.TypeOf(defaults.UserOptions{}), tests)
}

func TestRuncmd(t *testing.T) {
	tests := []Test{
		{"Name", reflect.String, `yaml:"name"`},
		{"Cmds", reflect.Slice, `yaml:"cmds"`},
	}

	runTests(t, reflect.TypeOf(defaults.Runcmd{}), tests)
}
