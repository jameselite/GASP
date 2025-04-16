package routers

var RouterTemplateGin string = `

package routers

import (
	"github.com/gin-gonic/gin"
)

func %s(app *gin.Engine) {

	{
		router := app.Group("%s")

	}

}
	
`

var RouterTemplateFiber string = `

package routers

import (
	"github.com/gofiber/fiber/v2"
)

func %s(app *fiber.App) {

	{
		router := app.Group("%s")

	}
}

	
`