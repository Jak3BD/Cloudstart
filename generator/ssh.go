package generator

import (
	"cloudstart/config"
	"cloudstart/utils"

	"github.com/AlecAivazis/survey/v2"
)

type SSH struct {
	Pwauth         bool `yaml:"pwauth,omitempty"`
	DisableRoot    bool `yaml:"disable_root,omitempty"`
	NoFingerprints bool `yaml:"no_fingerprints,omitempty"`
}

func (s *SSH) Generate() {
	s.SetPwauth()
	s.SetDisableRoot()
	s.SetNoFingerprints()
}

func (s *SSH) SetPwauth() {
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("SSH password authentication", s.Pwauth),
		Default: s.Pwauth,
	}, &config.Store.SSHPwauth); err != nil {
		panic(err)
	}
}

func (s *SSH) SetDisableRoot() {
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("SSH disable root", s.DisableRoot),
		Default: s.DisableRoot,
	}, &config.Store.SSHDisableRoot); err != nil {
		panic(err)
	}
}

func (s *SSH) SetNoFingerprints() {
	if err := survey.AskOne(&survey.Confirm{
		Message: utils.PrintDefault("SSH no fingerprints", s.NoFingerprints),
		Default: s.NoFingerprints,
	}, &config.Store.SSHNoFingerprints); err != nil {
		panic(err)
	}
}
