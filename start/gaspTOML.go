package start

import (
	"errors"
	"os"
	"github.com/BurntSushi/toml"
)

type tomlContent struct {
	Projectname string `toml:"projectname"`
	Framework string `toml:"framework"`
	Architecture string `toml:"architecture"`
	Version string `toml:"version"`
	Database string `toml:"database"`
}

func MakegaspTOML(projectname string, framework string, arch string, version string, database string) (string, error) {

	var content tomlContent

	content.Projectname = projectname
	content.Framework = framework
	content.Architecture = arch
	content.Version = version
	content.Database = database

	tomlFile, tomfileErr := os.Create("gasp.toml")

	if tomfileErr != nil {
		return "", errors.New(tomfileErr.Error())
	}

	defer tomlFile.Close()

	tomlData, tomlDataErr := toml.Marshal(content)

	if tomlDataErr != nil {
		return "", errors.New(tomlDataErr.Error())
	}

	_, writeErr := tomlFile.Write(tomlData)

	if writeErr != nil {
		return "", errors.New(writeErr.Error())
	}


	return "GASP config created !", nil
	
}