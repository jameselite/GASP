package commands

import (
	"errors"
	"os"
	"os/exec"
)

func MakeGit() (string, error) {

	var gitCommand *exec.Cmd = exec.Command("git", "init")

	var gitErr error = gitCommand.Run()

	if gitErr != nil {
		return "", errors.New(gitErr.Error())
	}

	var gitignoreFile *os.File
	var gitignoreErr error
	
	defer gitignoreFile.Close()

	gitignoreFile, gitignoreErr = os.Create(".gitignore")

	if gitignoreErr != nil {
		return "", errors.New(gitignoreErr.Error())
	}

	_, writeErr := gitignoreFile.WriteString(".env")

	if writeErr != nil {
		return "", errors.New(writeErr.Error())
	}

	return "Git init was successfull !", nil

}