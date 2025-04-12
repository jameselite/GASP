package start

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func MakeStart(projectName string) (string, error) {

	var ginInstall *exec.Cmd = exec.Command("go", "get", "github.com/gin-gonic/gin")

	var ginInstallErr error = ginInstall.Run()
	if ginInstallErr != nil {
		return "", errors.New(ginInstallErr.Error())
	}

	mainFile, makeMainErr := os.Create("cmd/main.go")
	if makeMainErr != nil {
		return "", errors.New(makeMainErr.Error())
	}
	
	defer mainFile.Close()
	
	var mainContent string = fmt.Sprintf(StartTemplate, projectName + "/config")

	_, writeMainErr := mainFile.WriteString(mainContent)
	if writeMainErr != nil {
		return "", errors.New(writeMainErr.Error())
	}

	return "Main.go created !", nil
	
} 