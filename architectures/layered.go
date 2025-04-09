package architectures

import (
	"errors"
	"os"
)

func MakeLayered() (string, error) {

	var layeredDirs []string = []string{
		"internal/controllers",
		"internal/services",
		"internal/routers",
		"internal/middlewares",
		"internal/db",
	}

	for _, dir := range layeredDirs {
		var dirErr error = os.Mkdir(dir, 0755)
		if dirErr != nil {
			return "", errors.New(dirErr.Error())
		}
	}

	return "layered base files created !", nil
}