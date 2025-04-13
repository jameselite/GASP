package start

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

func ParseTOML() (*TomlContent, error) {

	tomlFile, tomlReadErr := os.ReadFile("gasp.toml")

	if tomlReadErr != nil {
		return nil, errors.New(tomlReadErr.Error())
	}

	var parsedTOML TomlContent

	if _, parseErr := toml.Decode(string(tomlFile), &parsedTOML); parseErr != nil {

		return nil, errors.New(parseErr.Error())
	}

	return &parsedTOML, nil

}