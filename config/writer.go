package config

import (
	"bytes"
	"cloudstart/version"
	"fmt"
	"os"
	"path"
	"text/template"
	"time"

	"gopkg.in/yaml.v2"
)

func Write() {
	resolvedPath := os.ExpandEnv(Meta.FilePath)

	data, err := yaml.Marshal(&Store)
	if err != nil {
		panic(err)
	}

	fileName := fmt.Sprintf("%s.yaml", Meta.FileName)
	filePath := path.Join(resolvedPath, fileName)

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	header := header(Meta.UserCreds)
	if err != nil {
		panic(err)
	}

	_, err = file.WriteString(header)
	if err != nil {
		panic(err)
	}

	_, err = file.Write(data)
	if err != nil {
		panic(err)
	}
}

func header(userCreds []UserCreds) string {
	tmpl, err := template.New("example").Parse(`#cloud-config

############################## Cloudstart ##############################



#########################################
############### REMOVE ME ###############
{{range .UserCreds}}
# {{printf "%-16s: %s" .Name .Passwd}}{{end}}

############### REMOVE ME ###############
#########################################



# Cloudstart: v{{.Version}} ({{.OS}}/{{.Arch}})
# Created:    {{.Time}}
# License:    Apache 2.0 (https://www.apache.org/licenses/LICENSE-2.0.txt)

############################## Cloudstart ##############################

`)
	if err != nil {
		panic(err)
	}

	data := struct {
		UserCreds []UserCreds
		Time      string
		Version   string
		OS        string
		Arch      string
	}{
		UserCreds: userCreds,
		Time:      time.Now().Format(time.RFC3339),
		Version:   version.Version,
		OS:        version.OS,
		Arch:      version.Arch,
	}

	var result bytes.Buffer
	if err := tmpl.Execute(&result, data); err != nil {
		panic(err)
	}

	return result.String()
}
