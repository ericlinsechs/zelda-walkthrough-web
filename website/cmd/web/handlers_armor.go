package main

import (
	"fmt"
	"net/http"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"github.com/gin-gonic/gin"
)

type armorTemplateData struct {
	Armor  models.Armor
	Armors []models.Armor
}

func (app *application) armorList(c *gin.Context) {

	// Get armor list from API
	var atd armorTemplateData
	app.infoLog.Println("Calling armor API...")
	app.getAPIContent(app.apis.armor, &atd.Armors)
	app.infoLog.Println(atd.Armors)

	// Load template files
	c.HTML(http.StatusOK, "armors/list", gin.H{
		"Armors": atd.Armors,
	})
}

func (app *application) armorView(c *gin.Context) {
	// Get id from incoming url
	armorID := c.Param("id")

	// Get armor list from API
	url := fmt.Sprintf("%s%s", app.apis.armor, armorID)
	app.infoLog.Printf("Calling api url: %s\n", url)

	var atd armorTemplateData
	app.getAPIContent(url, &atd.Armor)
	app.infoLog.Println(atd.Armor)

	// Load template files
	c.HTML(http.StatusOK, "armors/view", gin.H{
		"Name":     atd.Armor.Name,
		"Location": atd.Armor.Location,
		"Cost":     atd.Armor.Cost,
		"Uses":     atd.Armor.Uses,
		"ImageUrl": atd.Armor.ImageUrl,
	})
}
