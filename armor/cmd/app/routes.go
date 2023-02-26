package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.Engine) *gin.Engine {
	// Register handler functions.
	r.GET("/api/armor/item/", app.getAllItem)
	r.POST("/api/armor/item/", app.createItem)
	r.GET("/api/armor/item/:id", app.findItem)
	// r.POST("/api/armormany/", app.createMany)
	// r.DELETE("/api/armor/:id", app.delete)

	r.GET("/api/armor/set/", app.getAllSet)
	r.POST("/api/armor/set/", app.createSet)
	r.GET("/api/armor/set/:id", app.findSet)
	r.DELETE("/api/armor/set/:id", app.deleteSet)
	return r
}
