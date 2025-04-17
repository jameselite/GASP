package configfile

import (
	"errors"
	"fmt"
	"github.com/jameselite/GASP/start"
	"os"

	"os/exec"
)

func MakeConfig() (string, error) {

	project, tomlErr := start.ParseTOML()

	if tomlErr != nil {
		return "", errors.New(tomlErr.Error())
	}

	var dsn string

	var cmdEnv *exec.Cmd = exec.Command("go", "get", "github.com/joho/godotenv")

	fmt.Println("installing dotenv...")

	var cmdEnvErr error = cmdEnv.Run()


	if cmdEnvErr != nil {
		return "", errors.New(cmdEnvErr.Error())
	}

	_, envFileErr := os.Create(".env")
	if envFileErr != nil {
		return "", errors.New(envFileErr.Error())
	}

	switch project.Database {

	case "postgres":

		var cmdDriver *exec.Cmd = exec.Command("go", "get", "github.com/jackc/pgx/v5/pgxpool")
	
		fmt.Println("installing pgx/V5/pgxpool...")

		var cmdDriverErr error = cmdDriver.Run()
	
		if cmdDriverErr != nil {
			return "", errors.New(cmdDriverErr.Error())
		}

		dsn = fmt.Sprintf("postgres://%s:%s@localhost:5432/%s", project.Database_user, project.Database_pass, project.Database_name)

		PostgreConfig = fmt.Sprintf(PostgreConfig, "github.com/jackc/pgx/v5/pgxpool",
		"*pgxpool.Pool",
		"*pgxpool.Pool", 
		dsn,
		"pgxpool.New(context.Background(), dsn)",
		)

		file, fileErr := os.Create("config/config.go")

		if fileErr != nil {
			return "", errors.New(fileErr.Error())
		}

		defer file.Close()

		_, writeErr := file.WriteString(PostgreConfig)

		if writeErr != nil {
			return "", errors.New(writeErr.Error())
		}

	case "mysql":

		var cmdDriver *exec.Cmd = exec.Command("go", "get", "github.com/go-sql-driver/mysql")
	
		fmt.Println("installing go-sql-driver/mysql...")

		var cmdDriverErr error = cmdDriver.Run()
	
		if cmdDriverErr != nil {
			return "", errors.New(cmdDriverErr.Error())
		}
		
		dsn = fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s", project.Database_user, project.Database_pass, project.Projectname)

		MySQLConfig = fmt.Sprintf(MySQLConfig, "database/sql", 
		"github.com/go-sql-driver/mysql",
		"*sql.DB",
		"*sql.DB",
		dsn,
		`sql.Open("mysql", dsn)`,
		)

		file, fileErr := os.Create("config/config.go")

		if fileErr != nil {
			return "", errors.New(fileErr.Error())
		}

		defer file.Close()

		_, writeErr := file.WriteString(MySQLConfig)

		if writeErr != nil {
			return "", errors.New(writeErr.Error())
		}

	default:
		return "", errors.New("sorry, your database is not in our supported list")
	}

	return "Creating config file...", nil
}