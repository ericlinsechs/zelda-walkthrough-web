package main

import (
	"fmt"
	"net/http"

	"github.com/ericlinsechs/go-mongodb-microservices/movies/pkg/models"
	"github.com/gin-gonic/gin"
)

type movieTemplateData struct {
	Movie  models.Movie
	Movies []models.Movie
}

func (app *application) moviesList(c *gin.Context) {

	// Get movies list from API
	var mtd movieTemplateData
	app.infoLog.Println("Calling movies API...")
	app.getAPIContent(app.apis.movies, &mtd.Movies)
	app.infoLog.Println(mtd.Movies)

	// Load template files
	c.HTML(http.StatusOK, "movies/list", gin.H{
		"Movies": mtd.Movies,
	})
}

func (app *application) moviesView(c *gin.Context) {
	// Get id from incoming url
	movieID := c.Param("id")

	// Get movies list from API
	app.infoLog.Println("Calling movies API...")
	url := fmt.Sprintf("%s%s", app.apis.movies, movieID)

	var mtd movieTemplateData
	app.getAPIContent(url, &mtd.Movie)
	app.infoLog.Println(mtd.Movie)

	// Load template files
	c.HTML(http.StatusOK, "movies/view", gin.H{
		"Title":    mtd.Movie.Title,
		"Director": mtd.Movie.Director,
		"Rating":   mtd.Movie.Rating,
	})
}
