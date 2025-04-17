package architectures

import (
	"errors"
	"fmt"
	configfile "github.com/jameselite/GASP/config_file"
	"github.com/jameselite/GASP/start"
	"os"
	"os/exec"
)

func MakeSqlc(dbdir string) (string, error) {

	project, tomlErr := start.ParseTOML()

	if tomlErr != nil {
		return "", errors.New(tomlErr.Error())
	}

	var cmdSqlc *exec.Cmd = exec.Command("go", "install", "github.com/sqlc-dev/sqlc/cmd/sqlc@latest")

	fmt.Println("Getting sqlc ...")

	var cmdSqlcErr error = cmdSqlc.Run()
	if cmdSqlcErr != nil {
		return "", errors.New(cmdSqlcErr.Error())
	}

	sqlcFile, sqlcfileErr := os.Create("sqlc.yaml")
	if sqlcfileErr != nil {
		return "", errors.New(sqlcfileErr.Error())
	}

	_, schemaErr := os.Create("schema.sql")
	if schemaErr != nil {
		return "", errors.New(schemaErr.Error())
	}

	_, queriesErr := os.Create("query.sql")
	if queriesErr != nil {
		return "", errors.New(queriesErr.Error())
	}

	switch project.Database {

	case "postgres":
		
		var sqlcYamlTemp string = fmt.Sprintf(configfile.SqlcYamlPG, project.Projectname, dbdir)

		_, writeErr := sqlcFile.WriteString(sqlcYamlTemp)
		if writeErr != nil {
			return "", errors.New(writeErr.Error())
		}

	case "mysql":
		
		var sqlcYamlTemp string = fmt.Sprintf(configfile.SqlcYamlPG, project.Projectname, dbdir)

		_, writeErr := sqlcFile.WriteString(sqlcYamlTemp)
		if writeErr != nil {
			return "", errors.New(writeErr.Error())
		}

	default:
		return "", errors.New("sorry, your database is not supported")
	}

	return "Sqlc installed and ready to use !", nil

}