package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.Engine) *gin.Engine {

	r.LoadHTMLGlob("../../ui/html/**/*.tmpl")

	// Register handler functions.
	r.GET("/", app.home)
	r.GET("/armor/list", app.armorList)
	r.GET("/armor/view/:id", app.armorView)

	// static path
	r.Static("/static/", "../../ui/static")
	return r
}
