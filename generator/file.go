package generator

import (
	"cloudstart/config"

	"github.com/AlecAivazis/survey/v2"
)

type File struct {
	Path string
	Name string
}

func (f *File) Generate() {
	f.SetPath()
	f.SetName()
}

func (f *File) SetPath() {
	if err := survey.AskOne(&survey.Input{
		Message: "Config file path",
		Default: f.Path,
	}, &config.Meta.FilePath); err != nil {
		panic(err)
	}
}

func (f *File) SetName() {
	if err := survey.AskOne(&survey.Input{
		Message: "Config file name",
		Default: f.Name,
	}, &config.Meta.FileName); err != nil {
		panic(err)
	}
}
