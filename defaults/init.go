package defaults

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

const configFilePath = "$HOME/.config/cloudstart/defaults.yml"

var Store Defaults

func Init() {
	resolvedPath := os.ExpandEnv(configFilePath)

	if _, err := os.Stat(resolvedPath); os.IsNotExist(err) {
		fmt.Println("file or directory does not exist, creating '$HOME/.config/cloudstart/defaults.json'...")
		createFile(resolvedPath)
		return
	}

	file, err := os.Open(resolvedPath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := yaml.NewDecoder(file).Decode(&Store); err != nil {
		panic(err)
	}
}

func createFile(path string) {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, os.ModePerm); err != nil {
		panic(err)
	}

	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	if err := yaml.NewEncoder(file).Encode(Store); err != nil {
		panic(err)
	}
}
