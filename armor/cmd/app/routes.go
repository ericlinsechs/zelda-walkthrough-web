package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.Engine) *gin.Engine {
	// Armor item
	r.GET("/api/armor/item/", app.getAllItem)
	r.POST("/api/armor/item/", app.createItem)
	r.GET("/api/armor/item/:id", app.findItem)
	r.POST("/api/armor/item/many/", app.createManyItem)
	r.DELETE("/api/armor/item/:id", app.deleteItem)

	// Armor set
	r.GET("/api/armor/set/", app.getAllSet)
	r.POST("/api/armor/set/", app.createSet)
	r.POST("/api/armor/set/many/", app.createManySet)
	r.GET("/api/armor/set/:id", app.findSet)
	r.DELETE("/api/armor/set/:id", app.deleteSet)

	// Armor Image
	r.GET("/api/armor/image/", app.getAllImage)
	r.POST("/api/armor/image/", app.createImage)
	r.GET("/api/armor/image/:id", app.findImage)
	r.GET("/api/armor/image/name/:name", app.findImageByName)
	r.POST("/api/armor/image/name/update", app.updateImageName)
	r.DELETE("/api/armor/image/:id", app.deleteImage)
	return r
}
