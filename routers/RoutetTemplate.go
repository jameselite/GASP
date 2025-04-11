package routers

var RouterTemplate string = `

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