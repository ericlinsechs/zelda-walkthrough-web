package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArmorSet struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SetName     string             `bson:"setname,omitempty"`
	Effect      string             `bson:"effect,omitempty"`
	SetBonus    string             `bson:"setbonus,omitempty"`
	HowToObtain []string           `bson:"howtoobtain,omitempty"`
	Tag         *Item              `bson:"tag,omitempty"`
}
type Item struct {
	HeadGear string `bson:"headgear,omitempty"`
	BodyGear string `bson:"bodygear,omitempty"`
	LegGear  string `bson:"leggear,omitempty"`
}

type armorTemplateData struct {
	ArmorSet  ArmorSet
	ArmorSets []ArmorSet
}

func (app *application) armorList(c *gin.Context) {

	// Get armor list from API
	var atd armorTemplateData
	app.infoLog.Println("Calling armor API...")
	app.getAPIContent(app.apis.armorSet, &atd.ArmorSets)
	app.infoLog.Println(atd.ArmorSets)

	// Load template files
	c.HTML(http.StatusOK, "armors/list", gin.H{
		"ArmorSet": atd.ArmorSets,
	})
}

func (app *application) armorView(c *gin.Context) {
	// Get id from incoming url
	armorID := c.Param("id")

	// Get armor list from API
	url := fmt.Sprintf("%s%s", app.apis.armorSet, armorID)
	app.infoLog.Printf("Calling api url: %s\n", url)

	var atd armorTemplateData
	app.getAPIContent(url, &atd.ArmorSet)
	app.infoLog.Println(atd.ArmorSet)

	// Load template files
	c.HTML(http.StatusOK, "armors/view", gin.H{
		"ArmorSet": atd.ArmorSet,
		"HeadGear": atd.ArmorSet.HowToObtain[0],
		"BodyGear": atd.ArmorSet.HowToObtain[1],
		"LegGear":  atd.ArmorSet.HowToObtain[2],
	})
}
