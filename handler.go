package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
)

func HandlerLogin(ctx *gin.Context) {
	// bin from body payload
	var payload UserLoginPayload
	err := ctx.Bind(&payload)
	if err != nil {
		err := NewErrorValidation("user", "invalid payload", err)
		respError(ctx, "user", err)
		return
	}
	// get user in db
	u, err := db.GetUserByEmail(payload.Email)
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	// verify if the user has the good credentials
	if u.Password != payload.Password {
		err := NewErrorAutorization("user", u.UUID)
		respError(ctx, "user", err)
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
	err = db.SetUser(usr)
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	ctx.JSON(http.StatusOK, usr)
}

func HandlerGetUserByID(ctx *gin.Context) {
	uuidParam := ctx.Param("uuid")
	log.Println("HandlerGetUserByID: uuidParam", uuidParam)
	_, err := uuid.Parse(uuidParam)
	log.Println("HandlerGetUserByID: err", err)
	if err != nil {
		log.Println("validation uuid", err)
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	u, err := db.GetUserByID(uuidParam)
	log.Println("HandlerGetUserByID: GetUserByID:", u, err)
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func HandlerSearchUser(ctx *gin.Context) {
	name := ctx.Query("name")
	if len(name) > 0 {
		u, err := db.GetUserByName(name)
		if err != nil {
			respError(ctx, "user", err)
			return
		}
		ctx.JSON(http.StatusOK, u)
		return
	}
	u, err := db.GetAllUser()
	if err != nil {
		respError(ctx, "user", err)
		return
	}
	ctx.JSON(http.StatusOK, u)
}

func respNotfound(ctx *gin.Context, entity string) {
	ctx.JSON(http.StatusNotFound, gin.H{
		"success": false,
		"error":   entity + " not found",
	})
}

func respError(ctx *gin.Context, entity string, err error) {
	log.Println(err.Error())
	switch err := err.(type) {
	case *ErrorHandlerValidation:
		ctx.JSON(err.StatusCode, gin.H{
			"success": false,
			"error":   err.Message,
		})
	case *ErrorDB:
		ctx.JSON(err.StatusCode, gin.H{
			"success": false,
			"error":   err.Message,
		})
	default:
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"error":   entity + " " + err.Error(),
		})
	}
}

type ErrorHandlerValidation struct {
	Err        error
	Message    string
	Entity     string
	StatusCode int
}

func (e *ErrorHandlerValidation) Error() string {
	return e.Message
}

func NewErrorValidation(entity, message string, errOrigin error) *ErrorHandlerValidation {
	return &ErrorHandlerValidation{
		Err:        errOrigin,
		Message:    message,
		Entity:     entity,
		StatusCode: 400,
	}
}

func NewErrorAutorization(entity, uuidUser string) *ErrorHandlerValidation {
	return &ErrorHandlerValidation{
		Err:        nil,
		Message:    "user not authorized",
		Entity:     entity + " " + uuidUser,
		StatusCode: 401,
	}
}
