package start

import (
	"errors"

	"github.com/BurntSushi/toml"
)

func ParseTOML() (*TomlContent, error) {

	var parsedTOML TomlContent

	if _, parseErr := toml.Decode("gasp.toml", &parsedTOML); parseErr != nil {
		
		return nil, errors.New(parseErr.Error())
	}

	return &parsedTOML, nil

}