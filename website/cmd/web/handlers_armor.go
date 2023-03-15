package main

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strings"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"github.com/gin-gonic/gin"
)

type armorTemplateData struct {
	ArmorSet    models.ArmorSet
	ArmorSets   []models.ArmorSet
	ArmorItem   models.ArmorItem
	ArmorItems  []models.ArmorItem
	ArmorImage  models.ArmorImage
	ArmorImages []models.ArmorImage
}

func (app *application) armorList(c *gin.Context) {
	// Get armor list from API
	var atd armorTemplateData
	app.infoLog.Println("Calling armor API...")
	app.getAPIContent(app.apis.armorSet, &atd.ArmorSets)
	// app.infoLog.Println(atd.ArmorSets)

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
		// Convert image to base64 format
		atd.ArmorItems[i].ImageData = EncodeImageToBase64(temp.ArmorImage)
	}

	// Load template files
	c.HTML(http.StatusOK, "armors/view", gin.H{
		"SetName":    atd.ArmorItems[0].SetName,
		"ArmorItems": atd.ArmorItems,
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
