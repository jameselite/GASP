package routers

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

func MakeRouter(group_path string, arch int64) (string, error) {

	var routerContent string = fmt.Sprintf(RouterTemplate, group_path, group_path)

	var haveBlank bool = strings.Contains(group_path, " ")

	if haveBlank {
		return "", errors.New("sorry, you can not have blank space as group path for your router")
	}

	switch arch {

	case 1: // layered 

	var routerFilePath string = fmt.Sprintf("internal/routers/%s.go", group_path)
	routerFile, routerFileErr := os.Create(routerFilePath)
	if routerFileErr != nil {
		return "", errors.New(routerFileErr.Error())
	}

	defer routerFile.Close()

	_, writeRouterErr := routerFile.WriteString(routerContent)
	if writeRouterErr != nil {
		return "", errors.New(writeRouterErr.Error())
	}


	case 2: // clean

	var routerFilePath string = fmt.Sprintf("delivery/routers/%s.go", group_path)
	routerFile, routerFileErr := os.Create(routerFilePath)
	if routerFileErr != nil {
		return "", errors.New(routerFileErr.Error())
	}

	defer routerFile.Close()

	_, writeRouterErr := routerFile.WriteString(routerContent)
	if writeRouterErr != nil {
		return "", errors.New(writeRouterErr.Error())
	}

	}

	return group_path + " router created !", nil
}