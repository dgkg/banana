package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"banana/apifront/db"
	"banana/apifront/handler"
	"banana/concert"
)

const (
	EnvProd    = "prod"
	EnvPreprod = "preprod"
	EnvDev     = "dev"
	EnvLocal   = "local"
)

func main() {
	env := os.Getenv("ENV")
	if env == "" {
		env = EnvLocal
	}
	sdkConcert := concert.New("CVPUbWJa4ItbkVQDmExWnyBdUKkKwMpx2Vbn")
	log.Println("ENV:", env)
	// init db and elements of my app depending of the env
	var myDb db.DB
	switch env {
	case EnvLocal:
		myDb = db.NewMoke()
	case EnvPreprod:
		myDb = db.NewMoke()
	case EnvDev:
		myDb = db.NewMoke()
	case EnvProd:
		myDb = db.NewSQLite("banana.db")
	default:
		panic("error creating db")
	}

	// create handler
	myHandler := handler.NewHandler(myDb, sdkConcert)
	// init routes
	r := gin.Default()
	myHandler.InitRoutes(r)
	// listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	r.Run(":8000")
}
