package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

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
	ImageData   string             `bson:"imagedata,omitempty"`
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

type ArmorImage struct {
	ID   primitive.ObjectID `bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Data []byte             `bson:"data"`
}

type armorTemplateData struct {
	ArmorSet    ArmorSet
	ArmorSets   []ArmorSet
	ArmorItem   ArmorItem
	ArmorItems  []ArmorItem
	ArmorImage  ArmorImage
	ArmorImages []ArmorImage
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

	for i, itemID := range atd.ArmorSet.Tag {
		// Get items
		url := fmt.Sprintf("%s%s", app.apis.armorItem, itemID)
		app.infoLog.Printf("Calling api url: %s\n", url)
		var temp armorTemplateData
		app.getAPIContent(url, &temp.ArmorItem)
		atd.ArmorItems = append(atd.ArmorItems, temp.ArmorItem)
		// Get image
		url = fmt.Sprintf("%s%s", app.apis.armorImage, convertNameFormat(atd.ArmorItems[i].Name))
		app.infoLog.Printf("Calling api url: %s\n", url)
		app.getAPIContent(url, &temp.ArmorImage)
		// atd.ArmorImages = append(atd.ArmorImages, temp.ArmorImage)
		atd.ArmorItems[i].ImageData = EncodeImageToBase64(temp.ArmorImage)
	}

	// imageData := EncodeImageToBase64(atd.ArmorImages)

	// Load template files
	c.HTML(http.StatusOK, "armors/view", gin.H{
		"SetName":    atd.ArmorItems[0].SetName,
		"ArmorItems": atd.ArmorItems,
		// "ArmorImages": imageData,
	})
}

func convertNameFormat(name string) string {
	name = strings.ReplaceAll(name, " ", "-")
	name = strings.ToLower(name) // convert to lowercase
	return name
}

func EncodeImageToBase64(src ArmorImage) (dst string) {
	// for _, image := range src {
	// 	// Encode the image data as a base64-encoded string.
	// 	encodedString := base64.StdEncoding.EncodeToString(image.Data)
	// 	dst = append(dst, encodedString)
	// }
	// Encode the image data as a base64-encoded string.
	dst = base64.StdEncoding.EncodeToString(src.Data)
	// dst = append(dst, encodedString)
	return dst
}
