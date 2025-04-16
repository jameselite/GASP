package commands

import (
	"errors"
	"fmt"
	"os"
)

func MakeMiddleware(name string) (string, error) {

	middlewareFile, makeErr := os.Create("internal/middlewares/"+name+".go")

	if makeErr != nil {
		return "", errors.New(makeErr.Error())
	}

	defer middlewareFile.Close()

	Content := fmt.Sprintf(MiddlewareTemplate, name)

	_, writeErr := middlewareFile.WriteString(Content)

	if writeErr != nil {
		return "", errors.New(writeErr.Error())
	}

	return "Middleware " + name + " Created successfully !", nil

}