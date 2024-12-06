package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

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

	var srvAPI *http.Server
	// start the api
	go func() {
		// create handler
		myHandler := handler.NewHandler(myDb)
		// init routes
		api := gin.Default()
		myHandler.InitRoutes(api)
		srvAPI = &http.Server{
			Addr:    ":8000",
			Handler: api.Handler(),
		}
		srvAPI.ListenAndServe()
	}()

	var srvConcert *http.Server
	go func() {
		sdkConcert := concert.New("CVPUbWJa4ItbkVQDmExWnyBdUKkKwMpx2Vbn")
		concertAPI := gin.Default()
		concertAPI.GET("/artists", func(ctx *gin.Context) {
			artists, err := sdkConcert.Artists.Search(nil)
			if err != nil {
				ctx.JSON(500, gin.H{"error": err.Error()})
				return
			}
			ctx.JSON(200, artists)
		})
		srvConcert = &http.Server{
			Addr:    ":8001",
			Handler: concertAPI.Handler(),
		}
		srvConcert.ListenAndServe()
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal, 1)
	// kill (no param) default send syscall.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go func() {
		if err := srvConcert.Shutdown(ctx); err != nil {
			log.Fatal("Server Concert Shutdown:", err)
		}
	}()
	if err := srvAPI.Shutdown(ctx); err != nil {
		log.Fatal("Server API Shutdown:", err)
	}
	// catching ctx.Done(). timeout of 5 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 5 seconds.")
	}
	log.Println("Server exiting")
}
