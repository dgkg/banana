package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"banana/db"
	"banana/handler"
)

func main() {
	myDb := db.NewDB()
	myHandler := handler.NewHandler(myDb)

	r := gin.Default()
	r.POST("/register", myHandler.Register)
	r.POST("/login", HandlerTest, myHandler.Login)
	r.GET("/users/:uuid", myHandler.GetUserByID)
	r.GET("/users", myHandler.SearchUser)
	r.Run(":8000") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func HandlerTest(ctx *gin.Context) {
	log.Println("test")
}
