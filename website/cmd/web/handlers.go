package main

import (
	"encoding/json"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func (app *application) home(c *gin.Context) {
	c.HTML(http.StatusOK, "home", nil)
}

func (app *application) getAPIContent(url string, templateData interface{}) error {
	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	if err := json.NewDecoder(r.Body).Decode(templateData); err != nil {
		app.errorLog.Println(err)
	}
	// app.infoLog.Println(templateData)

	return nil
}

func (app *application) static(dir string) http.Handler {
	dirCleaned := filepath.Clean(dir)
	return http.StripPrefix("/static/", http.FileServer(http.Dir(dirCleaned)))
}
