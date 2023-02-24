package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.Engine) *gin.Engine {
	// Register handler functions.
	r.GET("/api/armor/item/", app.getAllItem)
	r.POST("/api/armor/item/", app.createItem)
	r.GET("/api/armor/item/:id", app.findItemByID)
	// r.POST("/api/armormany/", app.createMany)
	// r.DELETE("/api/armor/:id", app.delete)

	return r
}
