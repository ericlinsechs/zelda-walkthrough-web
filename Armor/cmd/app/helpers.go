package main

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func (app *application) serverError(c *gin.Context, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Output(2, trace)

	c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
	app.errorLog.Println(err)
}

func (app *application) clientError(c *gin.Context, status int) {
	c.JSON(status, http.StatusText(status))
}
