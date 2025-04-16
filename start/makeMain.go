package start

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func MakeStart(projectName string, framework int64) (string, error) {

	var frameworkInstall *exec.Cmd

	mainFile, makeMainErr := os.Create("cmd/main.go")
	if makeMainErr != nil {
		return "", errors.New(makeMainErr.Error())
	}
	
	defer mainFile.Close()

	var mainContent string

	var FrameworkInstallErr error

	switch framework {

	case 1: // gin
		frameworkInstall = exec.Command("go", "get", "github.com/gin-gonic/gin")
		fmt.Println("Getting gin framework from : github.com/gin-gonic/gin")
		
		mainContent = fmt.Sprintf(StartTemplateGin, projectName + "/config")

		FrameworkInstallErr = frameworkInstall.Run()
	
		if FrameworkInstallErr != nil {
			return "", errors.New(FrameworkInstallErr.Error())
		}
	case 2: // fiber
		frameworkInstall = exec.Command("go", "get", "github.com/gofiber/fiber/v2")
		fmt.Println("Getting Fiber framework from : github.com/gofiber/fiber/v2")

		mainContent = fmt.Sprintf(StartTemplateFiber, projectName + "/config")

		FrameworkInstallErr = frameworkInstall.Run()
	
		if FrameworkInstallErr != nil {
			return "", errors.New(FrameworkInstallErr.Error())
		}

	}
	
	_, writeMainErr := mainFile.WriteString(mainContent)
	
	if writeMainErr != nil {
		return "", errors.New(writeMainErr.Error())
	}

	return "Main.go created !", nil
	
} 