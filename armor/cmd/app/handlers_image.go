package main

import (
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"github.com/gin-gonic/gin"
)

func (app *application) createImage(c *gin.Context) {
	// get the uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get image file"})
		return
	}

	// read the file contents
	fileData, err := file.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not read image file"})
		return
	}
	defer fileData.Close()

	// write the file to disk
	data, err := ioutil.ReadAll(fileData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not read image file"})
		return
	}

	imageName := file.Filename
	imageName = strings.ReplaceAll(imageName, ".png", "")
	imageName = strings.ToLower(imageName) // convert to lowercase
	imageName = strings.Title(imageName)   // capitalize first letter of each word

	// Create a new Image struct and store it in MongoDB
	newImage := &models.ArmorImage{
		Name: imageName,
		Data: data,
	}

	insertResult, err := app.armorImage.InsertImage(newImage)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Printf("inserted document with ID %v", insertResult.InsertedID)

	// err = ioutil.WriteFile(app.imageRoot+"/"+imageName, fileContents, 0644)
	// if err != nil {
	// 	c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not write image file"})
	// 	return
	// }
}

func (app *application) findImage(c *gin.Context) {
	// Get id from incoming url
	id := c.Param("id")

	result, err := app.armorImage.FindImage(id)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.clientError(c, http.StatusBadRequest)
			return
		}
		// Any other error will send an internal server error
		app.serverError(c, err)
	}

	app.infoLog.Printf("Founded document with ID %v\n", id)

	// // Send response
	// c.JSON(http.StatusOK, result)

	// Set the content type to the appropriate image format
	c.Header("Content-Type", "image/jpeg")

	// Write the image data to the response
	c.Writer.Write(result.Data)
}
