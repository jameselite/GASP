package architectures

import (
	"errors"
	"os"
)

func MakeClean() (string, error) {
	var cleanDirs []string = []string{
		"delivery",
		"delivery/routers",
		"delivery/handlers",
		"internal/entity",
		"internal/infrastructure",
		"internal/middlewares",
		"internal/usecases",
		"internal/utils",
		"internal/db",
	}

	for _, dir := range cleanDirs {

		var dirErr error = os.Mkdir(dir, 0755)
		if dirErr != nil {
			return "", errors.New(dirErr.Error())
		}

	}
	return "Clean architecture created !", nil
}