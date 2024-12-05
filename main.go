package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"banana/db"
	"banana/handler"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = "prod"
	}
	log.Println("ENV:", env)
	// init db and elements of my app
	var myDb db.DB
	if env == "local" {
		myDb = db.NewMoke()
	} else if env == "prod" {
		myDb = db.NewSQLite("banana.db")
	}
	if myDb == nil {
		panic("error creating db")
	}

	// create handler
	myHandler := handler.NewHandler(myDb)
	// init routes
	r := gin.Default()
	myHandler.InitRoutes(r)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8000")
}
