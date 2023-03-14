package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/ericlinsechs/zelda-walkthrough-web/armor/pkg/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
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

	imageName := convertImageNameFormat(file.Filename)

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

func (app *application) findImageByName(c *gin.Context) {
	// Get id from incoming url
	name := c.Param("name")

	result, err := app.armorImage.FindImageByName(name)
	if err != nil {
		if err.Error() == "ErrNoDocuments" {
			app.clientError(c, http.StatusBadRequest)
			return
		}
		// Any other error will send an internal server error
		app.serverError(c, err)
	}

	app.infoLog.Printf("Founded document with Name %v\n", name)

	// Send response
	c.JSON(http.StatusOK, result)

	// // Set the content type to the appropriate image format
	// c.Header("Content-Type", "image/jpeg")

	// // Write the image data to the response
	// c.Writer.Write(result.Data)
}

func (app *application) updateImageName(c *gin.Context) {
	newImage := new(models.ArmorImage)
	err := json.NewDecoder(c.Request.Body).Decode(&newImage)
	if err != nil {
		app.serverError(c, err)
	}

	app.infoLog.Println(newImage)

	// Define the filter and update operations.
	filter := bson.M{"_id": newImage.ID}
	update := bson.M{"$set": bson.M{"name": newImage.Name}}

	// Update the document.
	result, err := app.armorImage.Collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		// Handle the error.
	}

	// Print the number of documents that were modified.
	fmt.Printf("Updated %v documents.\n", result.ModifiedCount)
}

func convertImageNameFormat(name string) string {
	name = strings.ReplaceAll(name, ".png", "")
	name = strings.ToLower(name) // convert to lowercase
	return name
}
