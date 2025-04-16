package commands

import (
	"errors"
	"fmt"
	"goTmp/start"
	"os"
)

func MakeMiddleware(name string) (string, error) {

	middlewareFile, makeErr := os.Create("internal/middlewares/"+name+".go")

	project, tomlErr := start.ParseTOML()

	if tomlErr != nil {
		return "", errors.New(tomlErr.Error())
	}

	if makeErr != nil {
		return "", errors.New(makeErr.Error())
	}

	defer middlewareFile.Close()

	var Content string

	if project.Framework == "gin" {

		Content = fmt.Sprintf(MiddlewareTemplateGin, name)

	}

	if project.Framework == "fiber" {

		Content = fmt.Sprintf(MiddlewareTemplateFiber, name)

	}

	_, writeErr := middlewareFile.WriteString(Content)

	if writeErr != nil {
		return "", errors.New(writeErr.Error())
	}

	return "Middleware " + name + " Created successfully !", nil

}