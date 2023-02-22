package main

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"github.com/gin-gonic/gin"
)

func (app *application) getAll(c *gin.Context) {
	armor, err := app.armor.All()
	if err != nil {
		app.serverError(c, err)
	}
	app.infoLog.Println("Armors have been listed")

	// Send response
	c.JSON(http.StatusOK, armor)
}

func (app *application) findByID(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	movie, err := app.armor.FindByID(id)
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
	newArmor := new(models.Armor)

	err := json.NewDecoder(c.Request.Body).Decode(&newArmor)
	if err != nil {
		app.serverError(c, err)
	}

	newArmor.LastTimeEdit = time.Now().Format(time.UnixDate)

	// Insert new user
	insertResult, err := app.armor.Insert(newArmor)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("New armor have been created, %s", insertResult.InsertedID)

	// Send response back
	// c.JSON(http.StatusOK, users)
}

func (app *application) delete(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	// Delete user by id
	deleteResult, err := app.armor.Delete(id)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("%d armor(s) have been eliminated", deleteResult.DeletedCount)
}
