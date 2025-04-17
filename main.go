package main

import (
	"fmt"
	"goTmp/architectures"
	"goTmp/commands"
	configfile "goTmp/config_file"
	"goTmp/routers"
	"goTmp/start"
	"os"

	"github.com/olekukonko/tablewriter"
)

func ShowHelp() {

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Command", "Arg", "Description"})
	table.SetBorder(true)
	table.SetRowLine(true)
	table.SetAutoFormatHeaders(true)

	data := [][]string{
		{"init", "{project_name}", "Makes a new project inside of the current directory"},
		{"start", "none", "Start prompting user for setup the project"},
		{"add", "git", "Adds git ( if not already created )"},
		{"add", "redis", "Adds redis ( if not already added )"},
		{"add", "config", "Adds main config of app if not exist in /config/config.go"},
		{"generate", "router {group_path}", "Adds a router. group path is the static path of the endpoints for this router"},
		{"generate", "controller {router_name} {endpoint_path} {method} {controller_name}", "Generates a new Controller based on the name, method, and url path, needs a router. "},
		{"generate", "middleware {name}", "Adds a middlware with the name you provided"},
		{"generate", "main", "Adds the start point of your app based on your framework name, gin or fiber"},	
	}

	for _, v := range data {
		table.Append(v)
	}

	table.Render()

}

func main() {

	args := os.Args

	switch args[1] {

	case "-h", "--help":
		ShowHelp()
	
	case "init":
		fmt.Printf("Enter the name of your project : ")
		var projectname string

		fmt.Scan(&projectname)		

		makeProject, projectErr := commands.MakeProject(projectname)
		if projectErr != nil {
			fmt.Println(projectErr.Error())
			return
		}

		fmt.Println(makeProject)
		return
	
	case "start":
		
		var database, database_name, database_pass, database_user, framework, arch, version string
	
		var wantGit, wantRedis string

		fmt.Println(`What database are you using?
		type "postgres" for PostgreSQL, type "mysql" for MySQL : `)
		fmt.Scan(&database)

		fmt.Println("Enter the name of your database: ")
		fmt.Scan(&database_name)

		fmt.Println("Enter your database password: ")
		fmt.Scan(&database_pass)

		fmt.Println("Enter the name of your database user: ")
		fmt.Scan(&database_user)

		fmt.Println("What framework do you want to use ? : \n Enter gin or fiber")
		fmt.Scan(&framework)

		fmt.Println("What architecture do you want to use? : \n Enter layered or clean ")
		fmt.Scan(&arch)

		fmt.Println("What is the version of your app ? \n Enter like this: 1.2.5 ")
		fmt.Scan(&version)

		fmt.Println("Do you want Git ? : \n Enter yes or no")
		fmt.Scan(&wantGit)

		fmt.Println("Do you want Redis ? : \n Enter yes or no")
		fmt.Scan(&wantRedis)

		makeGaspToml, tomlErr := start.MakegaspTOML(framework, arch, version, database, database_pass, database_name, database_user)
		if tomlErr != nil {
			fmt.Println(tomlErr.Error())
			return
		}

		fmt.Println(makeGaspToml)

		makeConfig, configErr := configfile.MakeConfig()
		if configErr != nil {
			fmt.Println(configErr.Error())
			return
		}

		fmt.Println(makeConfig)

		if wantGit == "yes" {

			makeGit, gitErr := commands.MakeGit()
			if gitErr != nil {
				fmt.Println(gitErr.Error())
				return
			}

			fmt.Println(makeGit)

		}

		makeBased, baseErr := architectures.MakeBase()
		if baseErr != nil {
			fmt.Println(baseErr.Error())
			return
		}

		fmt.Println(makeBased)

		switch arch {
		case "layered":

			makeLayered, layeredErr := architectures.MakeLayered()
			if layeredErr != nil {
				fmt.Println(layeredErr.Error())
				return
			}

			fmt.Println(makeLayered)

			makesqlc, sqlcErr := architectures.MakeSqlc("internal/db")
			if sqlcErr != nil {
				fmt.Println(sqlcErr.Error())
				return
			}

			fmt.Println(makesqlc)
		
		case "clean":

			makeClean, cleanErr := architectures.MakeClean()
			if cleanErr != nil {
				fmt.Println(cleanErr.Error())
				return
			}

			fmt.Println(makeClean)

			makesqlc, sqlcErr := architectures.MakeSqlc("internal/db")
			if sqlcErr != nil {
				fmt.Println(sqlcErr.Error())
				return
			}

			fmt.Println(makesqlc)

		default:
			fmt.Println("sorry, your architecture is not supported")
		}


		if wantRedis == "yes" {

			makeRedis, redisErr := configfile.AddRedis()
			if redisErr != nil {
				fmt.Println(redisErr.Error())
				return
			}

			fmt.Println(makeRedis)

		}

		fmt.Println("Done !! \n Ready to go")
		return

	case "add":

		if args[2] == "git" {

			makeGit, gitErr := commands.MakeGit()
			if gitErr != nil {
				fmt.Println(gitErr.Error())
				return
			}

			fmt.Println(makeGit)

		}

		if args[2] == "redis" {

			makeRedis, redisErr := configfile.AddRedis()
			if redisErr != nil {
				fmt.Println(redisErr.Error())
				return
			}

			fmt.Println(makeRedis)

		}

		if args[2] == "config" {

			makeConfig, configErr := configfile.MakeConfig()
			if configErr != nil {
				fmt.Println(configErr.Error())
				return
			}

			fmt.Println(makeConfig)

		}

	case "generate":

		if args[2] == "router" {

			if len(args) < 4 {
				fmt.Println("Not enough argumant for generating router.")
				return
			}

			group_path := args[3]

			makeRouter, routerErr := routers.MakeRouter(group_path)

			if routerErr != nil {
				fmt.Println(routerErr.Error())
				return
			}

			fmt.Println(makeRouter)
			return
		}

		if args[2] == "middleware" {

			if len(args) < 4 {
				fmt.Println("Not enough argumant for generating middleware.")
				return
			}

			name := args[3]

			makeMiddleware, middlewareErr := commands.MakeMiddleware(name)

			if middlewareErr != nil {
				fmt.Println(middlewareErr.Error())
				return
			}

			fmt.Println(makeMiddleware)
			return
		}

		if args[2] == "main" {

			if len(args) < 4 {
				fmt.Println("Not enough argumant for generating main file.")
				return
			}

			makeMain, mainErr := start.MakeStart()

			if mainErr != nil {
				fmt.Println(mainErr.Error())
				return
			}

			fmt.Println(makeMain)
			return
		}

		if args[2] == "controller" {

			if len(args) < 5 {
				fmt.Println("Not enough argumant for generating controller.")
				return
			}

			routerName := args[3]
			endpointPath := args[4]
			method := args[5]
			controllerName := args[6]

			makeController, controllerErr := commands.MakeController(routerName, endpointPath, controllerName, method)
			if controllerErr != nil {
				fmt.Println(controllerErr.Error())
				return
			}

			fmt.Println(makeController)
			return
		}
	}
	
}