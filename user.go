package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type UserRegister struct {
	FirstName string `json:"first_name" form:"first_name"`
	LastName  string `json:"last_name" form:"last_name"`
	Email     string `json:"email" form:"email"`
	Password  string `json:"pass" form:"pass"`
}

type User struct {
	UUID      string    `json:"uuid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"pass"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeleteAt  time.Time `json:"-"`
}

func NewUser(fn, ln, email, pass string) *User {
	return &User{
		UUID:      uuid.New().String(),
		FirstName: fn,
		LastName:  ln,
		Email:     email,
		Password:  pass,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

func HandleRegister(ctx *gin.Context) {
	var payload UserRegister
	err := ctx.Bind(&payload)
	if err != nil {
		ctx.AbortWithStatus(http.StatusBadRequest)
		return
	}
	usr := NewUser(payload.FirstName, payload.LastName, payload.Email, payload.Password)
	ctx.JSON(http.StatusOK, usr)
}
