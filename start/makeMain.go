package start

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

func MakeStart() (string, error) {

	project, tomlErr := ParseTOML()

	if tomlErr != nil {
		return "", errors.New(tomlErr.Error())
	}

	var frameworkInstall *exec.Cmd

	mainFile, makeMainErr := os.Create("cmd/main.go")
	if makeMainErr != nil {
		return "", errors.New(makeMainErr.Error())
	}
	
	defer mainFile.Close()

	var mainContent string

	var FrameworkInstallErr error

	switch project.Framework {

	case "gin": // gin
		frameworkInstall = exec.Command("go", "get", "github.com/gin-gonic/gin")
		fmt.Println("Getting gin framework from : github.com/gin-gonic/gin")
		
		mainContent = fmt.Sprintf(StartTemplateGin, project.Projectname + "/config")

		FrameworkInstallErr = frameworkInstall.Run()
	
		if FrameworkInstallErr != nil {
			return "", errors.New(FrameworkInstallErr.Error())
		}
	case "fiber": // fiber
		frameworkInstall = exec.Command("go", "get", "github.com/gofiber/fiber/v2")
		fmt.Println("Getting Fiber framework from : github.com/gofiber/fiber/v2")

		mainContent = fmt.Sprintf(StartTemplateFiber, project.Projectname + "/config")

		FrameworkInstallErr = frameworkInstall.Run()
	
		if FrameworkInstallErr != nil {
			return "", errors.New(FrameworkInstallErr.Error())
		}

	default:
		return "", errors.New("sorry, your framework is not supported")
	}
	
	_, writeMainErr := mainFile.WriteString(mainContent)
	
	if writeMainErr != nil {
		return "", errors.New(writeMainErr.Error())
	}

	return "Main.go created !", nil
	
} 