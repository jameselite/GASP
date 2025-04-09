package architectures

import (
	"errors"
	"fmt"
	configfile "goTmp/config_file"
	"os"
	"os/exec"
)

func MakeSqlc(dbnumber int64, dbdir string, projectname string) (string, error) {

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

	switch dbnumber {

	case 1:
		
		var sqlcYamlTemp string = fmt.Sprintf(configfile.SqlcYamlPG, projectname, dbdir)

		_, writeErr := sqlcFile.WriteString(sqlcYamlTemp)
		if writeErr != nil {
			return "", errors.New(writeErr.Error())
		}

	case 2:
		
		var sqlcYamlTemp string = fmt.Sprintf(configfile.SqlcYamlPG, projectname, dbdir)

		_, writeErr := sqlcFile.WriteString(sqlcYamlTemp)
		if writeErr != nil {
			return "", errors.New(writeErr.Error())
		}
	}

	return "Sqlc installed and ready to use !", nil

}