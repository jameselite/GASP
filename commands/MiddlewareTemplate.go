package commands

var MiddlewareTemplateGin string = 

`
package middlewares 

import (
	"github.com/gin-gonic/gin"
)

func %s() gin.HandlerFunc {
	return func (ctx *gin.Context) {


		ctx.Next()
	}
}

`

var MiddlewareTemplateFiber string = 

`
package middlewares 

import (
	"github.com/gofiber/fiber/v2"
)

func %s() fiber.Handler {
	return func(ctx *fiber.Ctx) error {

		return ctx.Next()
	}
}

`