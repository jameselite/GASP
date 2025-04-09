package architectures

import (
	"errors"
	"os"
)

func MakeBase() (string, error) {

	var baseDirs []string = []string{ 
		"cmd",
		"config",
		"internal",
	}

	for _, dir := range baseDirs {
		var dirErr error = os.Mkdir(dir, 0755)

		if dirErr != nil {
			return "", errors.New(dirErr.Error())
		}
	}

	return "Base directories created successfully", nil
}