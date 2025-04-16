package commands

var MiddlewareTemplate string = 

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