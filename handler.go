package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
)

func HandlerLogin(ctx *gin.Context) {
	// bin from body payload
	var payload UserLoginPayload
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	// get user in db
	u, err := db.GetUserByEmail(payload.Email)
	if err != nil {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// verify if the user has the good credentials
	if u.Password != payload.Password {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	// user is authorized
	ctx.JSON(200, gin.H{"ok": true})
}

func HandleRegister(ctx *gin.Context) {
	var payload UserRegister
	err := ctx.Bind(&payload)
	// data, err := io.ReadAll(ctx.Request.Body)
	// if err != nil {
	// 	ctx.AbortWithStatus(http.StatusBadRequest)
	// 	return
	// }
	// log.Println("data in body:", string(data))
	// err = json.Unmarshal(data, &payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	err = validator.New().Struct(payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	log.Printf("UserRegisterPayload : %#v", payload)
	usr := NewUser(payload.FirstName, payload.LastName, payload.Email, payload.Password)
	db.SetUser(usr)
	ctx.JSON(http.StatusOK, usr)
}

func HandlerGetUserByID(ctx *gin.Context) {
	uuid := ctx.Param("uuid")
	u, err := db.GetUserByID(uuid)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error": ErrUserNotFound,
		})
		return
	}
	ctx.JSON(http.StatusOK, u)
}
