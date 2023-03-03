package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArmorSet struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	SetName  string             `bson:"setname,omitempty"`
	Effect   string             `bson:"effect,omitempty"`
	SetBonus string             `bson:"setbonus,omitempty"`
	Tag      []string           `bson:"tag,omitempty"`
	// CreatedOn   string             `bson:"createdOn,omitempty"`
}
type ArmorItem struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	SetName     string             `bson:"setname,omitempty"`
	Name        string             `bson:"name,omitempty"`
	HowToObtain string             `bson:"howtoobtain,omitempty"`
	Url         string             `bson:"url,omitempty"`
	Upgrade     UpgradeLevel       `bson:"upgrade,omitempty"`
}

type UpgradeLevel struct {
	FirstUpgrade  UpgradeInfo `bson:"firstupgrade,omitempty"`
	SecondUpgrade UpgradeInfo `bson:"secondupgrade,omitempty"`
	ThirdUpgrade  UpgradeInfo `bson:"thirdupgrade,omitempty"`
	FinalUpgrade  UpgradeInfo `bson:"finalupgrade,omitempty"`
}
type UpgradeInfo struct {
	Bonus     string   `bson:"bonus,omitempty"`
	Materials []string `bson:"materials,omitempty"`
}

type armorTemplateData struct {
	ArmorSet   ArmorSet
	ArmorSets  []ArmorSet
	ArmorItem  ArmorItem
	ArmorItems []ArmorItem
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

	// Get the armorSet from API
	url := fmt.Sprintf("%s%s", app.apis.armorSet, armorID)
	app.infoLog.Printf("Calling api url: %s\n", url)

	var atd armorTemplateData
	app.getAPIContent(url, &atd.ArmorSet)
	app.infoLog.Println(atd.ArmorSet)

	for _, itemID := range atd.ArmorSet.Tag {
		url := fmt.Sprintf("%s%s", app.apis.armorItem, itemID)
		app.infoLog.Printf("Calling api url: %s\n", url)
		var temp armorTemplateData
		app.getAPIContent(url, &temp.ArmorItem)
		atd.ArmorItems = append(atd.ArmorItems, temp.ArmorItem)
	}
	// Load template files
	c.HTML(http.StatusOK, "armors/view", gin.H{
		"SetName":    atd.ArmorItems[0].SetName,
		"ArmorItems": atd.ArmorItems,
	})
}
