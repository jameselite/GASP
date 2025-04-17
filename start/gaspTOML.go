package start

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

type TomlContent struct {
	Projectname string `toml:"projectname"`
	Framework string `toml:"framework"`
	Architecture string `toml:"architecture"`
	Version string `toml:"version"`
	Database string `toml:"database"`
	Database_name string `toml:"database_name"`
	Database_user string `toml:"database_user"`
	Database_pass string `toml:"database_pass"`
}

func MakegaspTOML(framework string, arch string, version string, database string, database_pass string, database_name string, database_user string) (string, error) {

	cwd, err := os.Getwd()

	if err != nil {
		return "", errors.New(err.Error())
	}

	currentDir := filepath.Base(cwd)

	var content TomlContent

	content.Projectname = currentDir
	content.Framework = framework
	content.Architecture = arch
	content.Version = version
	content.Database = database
	content.Database_name = database_name
	content.Database_pass = database_pass
	content.Database_user = database_user

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