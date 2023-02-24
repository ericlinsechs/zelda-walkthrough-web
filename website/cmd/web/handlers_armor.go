package main

import (
	"fmt"
	"net/http"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"github.com/gin-gonic/gin"
)

type armorTemplateData struct {
	ArmorItem  models.ArmorItem
	ArmorItems []models.ArmorItem
}

func (app *application) armorList(c *gin.Context) {

	// Get armor list from API
	var atd armorTemplateData
	app.infoLog.Println("Calling armor API...")
	app.getAPIContent(app.apis.armor, &atd.ArmorItems)
	app.infoLog.Println(atd.ArmorItems)

	// Load template files
	c.HTML(http.StatusOK, "armors/list", gin.H{
		"Armors": atd.ArmorItems,
	})
}

func (app *application) armorView(c *gin.Context) {
	// Get id from incoming url
	// armorID := c.Param("id")
	armorID := "63f8573d49b281506bdedae5"

	// Get armor list from API
	url := fmt.Sprintf("%s%s", app.apis.armor, armorID)
	app.infoLog.Printf("Calling api url: %s\n", url)

	var atd armorTemplateData
	// app.getAPIContent(url, &atd.ArmorItem)
	app.getAPIContent(app.apis.armor, &atd.ArmorItem)
	app.infoLog.Println(atd.ArmorItem)

	// Load template files
	c.HTML(http.StatusOK, "armors/view", gin.H{
		"Name":           atd.ArmorItem.Name,
		"Characteristic": atd.ArmorItem.Characteristic,
		"Location":       atd.ArmorItem.Location,
		"Upgrade":        atd.ArmorItem.Upgrade,
	})
}
