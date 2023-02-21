package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ericlinsechs/go-mongodb-microservices/movies/pkg/models"
	"github.com/gin-gonic/gin"
)

func (app *application) getAll(c *gin.Context) {
	movies, err := app.movies.All()
	if err != nil {
		app.serverError(c, err)
	}
	app.infoLog.Println("Users have been listed")

	// Send response
	c.JSON(http.StatusOK, movies)
}

func (app *application) findByID(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	movie, err := app.movies.FindByID(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.clientError(c, http.StatusBadRequest)
			return
		}
		// Any other error will send an internal server error
		app.serverError(c, err)
	}

	app.infoLog.Printf("Movie(id:%v) has been found!\n", id)

	// Send response
	c.JSON(http.StatusOK, movie)
}

func (app *application) create(c *gin.Context) {
	newMovie := new(models.Movie)

	err := json.NewDecoder(c.Request.Body).Decode(&newMovie)
	if err != nil {
		app.serverError(c, err)
	}

	newMovie.ReleaseTime = time.Now().Format(time.UnixDate)

	// Insert new user
	insertResult, err := app.movies.Insert(newMovie)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("New movie have been created, %s", insertResult.InsertedID)

	// Send response back
	// c.JSON(http.StatusOK, users)
}

func (app *application) delete(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	// Delete user by id
	deleteResult, err := app.movies.Delete(id)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("%d movie(s) have been eliminated", deleteResult.DeletedCount)
}
