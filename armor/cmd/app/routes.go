package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.Engine) *gin.Engine {
	// Register handler functions.
	r.GET("/api/armor/", app.getAll)
	r.GET("/api/armor/:id", app.findByID)
	r.POST("/api/armor/", app.create)
	r.POST("/api/armormany/", app.createMany)
	r.DELETE("/api/armor/:id", app.delete)

	return r
}
