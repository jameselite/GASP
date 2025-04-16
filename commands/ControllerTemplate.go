package commands

var ControllerTemplateGin string = `
package controllers

import (
	"github.com/gin-gonic/gin"
)

func %s(ctx *gin.context) {



}
`

var ControllerTemplateFiber string = `
package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func %s(ctx *fiber.Ctx) {



}
`