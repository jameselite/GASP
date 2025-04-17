package routers

import (
	"errors"
	"fmt"
	"github.com/jameselite/GASP/helper"
	"github.com/jameselite/GASP/start"
	"os"
	"strings"
)

func AddRouterToMain(group_path string) error {

	rawContent, readErr := os.ReadFile("cmd/main.go")
	
	if readErr != nil {
		return errors.New(readErr.Error())
	}

	splitedContent := strings.Split(string(rawContent), "\n")

	project, errToml := start.ParseTOML()
	if errToml != nil {
		return errToml
	}

	var updatedContent []string

	for _, line := range splitedContent {

		if strings.Contains(line, "import (") {
			updatedContent = append(updatedContent, line)
			updatedContent = append(updatedContent, "")

			switch project.Architecture {
			case "layered":
				importString := fmt.Sprintf(`"%s/internal/routers"`, project.Projectname)
				updatedContent = append(updatedContent, "	"+importString)

			case "clean":
				importString := fmt.Sprintf(`"%s/internal/routers"`, project.Projectname)
				updatedContent = append(updatedContent, "	" +importString)
			}
			continue
		}

		if strings.Contains(line, "gin.Default()") {
			updatedContent = append(updatedContent, line)
			updatedContent = append(updatedContent, "")

			updatedContent = append(updatedContent, "	router." + group_path + "(app)")
			continue
		}

		updatedContent = append(updatedContent, line)
	}

	finalContent := strings.Join(updatedContent, "\n")

	err := os.WriteFile("cmd/main.go", []byte(finalContent), 0644)
	if err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func MakeRouter(group_path string) (string, error) {

	project, tomlErr := start.ParseTOML()

	if tomlErr != nil {
		return "", errors.New(tomlErr.Error())
	}

	var routerContent string

	if project.Framework == "gin" {
		routerContent = fmt.Sprintf(RouterTemplateGin, helper.CapitalizeFirstLetter(group_path), group_path)
	}

	if project.Framework == "fiber" {
		routerContent = fmt.Sprintf(RouterTemplateFiber, helper.CapitalizeFirstLetter(group_path), group_path)
	}

	var haveBlank bool = strings.Contains(group_path, " ")

	if haveBlank {
		return "", errors.New("sorry, you can not have blank space as group path for your router")
	}

	switch project.Architecture {

	case "layered":

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


	case "clean":

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

	addMainErr := AddRouterToMain(group_path)
	if addMainErr != nil {
		return "", errors.New(addMainErr.Error())
	}

	return group_path + " router created !", nil
}