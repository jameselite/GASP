package commands

var ControllerTemplate string = `
package controllers

import (
	"github.com/gin-gonic/gin"
)

func %s(ctx *gin.context) {



}
`