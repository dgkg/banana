package main

import (
	"github.com/gin-gonic/gin"

	"banana/db"
	"banana/handler"
)

func main() {
	// init db and elements of my app
	myDb := db.NewDB()
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
