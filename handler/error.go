package handler

import (
	"banana/db"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func respError(ctx *gin.Context, entity string, err error) {
	log.Println(err.Error())
	switch err := err.(type) {
	case *ErrorHandlerValidation:
		ctx.JSON(err.StatusCode, gin.H{
			"success": false,
			"error":   err.Message,
		})
	case *db.ErrorDB:
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
