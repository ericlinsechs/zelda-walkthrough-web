package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type apis struct {
	users     string
	movies    string
	showtimes string
	bookings  string
}

type application struct {
	errorLog *log.Logger
	infoLog  *log.Logger
	apis     apis
}

func main() {
	serverAddr := flag.String("serverAddr", "", "HTTP server network address")
	serverPort := flag.Int("serverPort", 8000, "HTTP server network port")
	usersAPI := flag.String("usersAPI", "http://localhost:3000/api/users/", "Users API")
	moviesAPI := flag.String("moviesAPI", "http://localhost:3000/api/movies/", "Movies API")
	flag.Parse()

	// Create logger for writing information and error messages.
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	// Initialize a new instance of application containing the dependencies.
	app := &application{
		infoLog:  infoLog,
		errorLog: errorLog,
		apis: apis{
			users:  *usersAPI,
			movies: *moviesAPI,
		},
	}

	serverURI := fmt.Sprintf("%s:%d", *serverAddr, *serverPort)
	server := gin.Default()
	server = app.routes(server)
	if err := server.Run(serverURI); err != nil {
		errorLog.Fatal(err)
	}
}
