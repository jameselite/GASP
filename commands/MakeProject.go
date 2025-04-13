package commands

import (
	"errors"
	"os"
	"os/exec"
	"path/filepath"
)

func MakeProject(projectname string) (string, error) {

	projectDir := os.Mkdir(projectname, 0755)

	if projectDir != nil {
		return "", errors.New(projectDir.Error())
	}

	initCmd := exec.Command("go", "mod", "init", projectname)

	initCmd.Dir = filepath.Join(".", projectname)

	cmdErr := initCmd.Run()

	if cmdErr != nil {
		return "", errors.New(cmdErr.Error())
	}

	return "Project " + projectname + " created!", nil

}