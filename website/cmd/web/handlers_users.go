package main

import (
	"fmt"
	"net/http"

	"github.com/ericlinsechs/go-mongodb-microservices/users/pkg/models"
	"github.com/gin-gonic/gin"
)

type userTemplateData struct {
	User  models.User
	Users []models.User
}

func (app *application) usersList(c *gin.Context) {

	// Get users list from API
	utd := new(userTemplateData)
	err := app.getAPIContent(app.apis.users, &utd.Users)
	if err != nil {
		app.errorLog.Println(err.Error())
	}
	app.infoLog.Println(utd.Users)

	// Load template files
	c.HTML(http.StatusOK, "users/list", gin.H{
		"Users": utd.Users,
	})
}

func (app *application) usersView(c *gin.Context) {
	// Get id from incoming url
	userID := c.Param("id")

	// Get users list from API
	app.infoLog.Println("Calling users API...")
	url := fmt.Sprintf("%s%s", app.apis.users, userID)

	var utd userTemplateData
	app.getAPIContent(url, &utd.User)
	app.infoLog.Println(utd.User)

	// Load template files
	c.HTML(http.StatusOK, "users/view", gin.H{
		"Name":     utd.User.Name,
		"LastName": utd.User.LastName,
	})
}
