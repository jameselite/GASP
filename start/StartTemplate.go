package start

var StartTemplateGin string = `package main

import (
	%s
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadENV()

	app := gin.Default()

	port := os.Getenv("PORT")

	_, dberr := config.DBconnection()

	if dberr != nil {
		log.Fatalf("There is a problem in DB: %s", dberr.Error())
		return
	}

	if port == "" {
		log.Fatalf("There is a problem in loading .env")
		return
	}

	// GASP: register your routers here if you want to do it manually, example:
	// routers.AuthRouter(app)

	app.Run(port)
}
`

var StartTemplateFiber string = `package main

import (
	%s
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
)

func main() {

	config.LoadENV()

	app := fiber.New()

	port := os.Getenv("PORT")

	_, dberr := config.DBconnection()

	if dberr != nil {
		log.Fatalf("There is a problem in DB: %s", dberr.Error())
		return
	}

	if port == "" {
		log.Fatalf("There is a problem in loading .env")
		return
	}

	// GASP: register your routers here if you want to do it manually, example:
	// routers.AuthRouter(app)

	app.Listen(port)
}
`