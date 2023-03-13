package main

import (
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (app *application) createImage(c *gin.Context) {
	// get the uploaded file
	file, err := c.FormFile("image")
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "could not get image file"})
		return
	}

	imageName := c.Request.Header.Get("Image-Name")

	// fmt.Println(imageName)

	// read the file contents
	fileData, err := file.Open()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not read image file"})
		return
	}
	defer fileData.Close()

	// write the file to disk
	fileContents, err := ioutil.ReadAll(fileData)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not read image file"})
		return
	}

	err = ioutil.WriteFile(app.imageRoot+"/"+imageName, fileContents, 0644)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "could not write image file"})
		return
	}
}
