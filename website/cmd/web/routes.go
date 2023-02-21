package main

import (
	"github.com/gin-gonic/gin"
)

func (app *application) routes(r *gin.Engine) *gin.Engine {

	r.LoadHTMLGlob("../../ui/html/**/*.tmpl")

	// Register handler functions.
	r.GET("/", app.home)
	r.GET("/movies/list", app.moviesList)
	r.GET("/movies/view/:id", app.moviesView)
	r.GET("/users/list", app.usersList)
	r.GET("/users/view/:id", app.usersView)

	// static path
	r.Static("/static/", "../../ui/static")
	return r
}
