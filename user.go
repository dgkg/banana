package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserLoginPayload struct {
	Email    string   `json:"email" form:"email"`
	Password Password `json:"pass" form:"pass"`
}

type UserRegister struct {
	FirstName string   `json:"first_name" form:"first_name"`
	LastName  string   `json:"last_name" form:"last_name"`
	Email     string   `json:"email" form:"email" validate:"required,email"`
	Password  Password `json:"pass" form:"pass"`
}

type Password string

func (p *Password) UnmarshalJSON(b []byte) error {
	aux := ""
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return err
	}
	h := sha256.New()
	h.Write([]byte(aux))
	*p = Password(fmt.Sprintf("%x", h.Sum(nil)))
	return nil
}

func (p Password) String() string {
	return "****"
}

func (p Password) MarshalJSON() ([]byte, error) {
	var s string = "****"
	return json.Marshal(s)
}

type User struct {
	UUID      string    `json:"uuid"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  Password  `json:"pass"`
	CreatedAt time.Time `json:"-"`
	UpdatedAt time.Time `json:"-"`
	DeleteAt  time.Time `json:"-"`
}

func NewUser(fn, ln, email string, pass Password) *User {
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

var db DB

func init() {
	db = DB{
		Users: make(map[string]User),
	}
}

type DB struct {
	Users map[string]User
}

func (db *DB) SetUser(u *User) error {
	_, err := db.GetUserByEmail(u.Email)
	if err == nil {
		return errors.New("user email allready exists")
	}
	db.Users[u.UUID] = *u
	return nil
}

func (db *DB) GetUserByID(uuid string) (*User, error) {
	u, ok := db.Users[uuid]
	if !ok {
		return nil, errors.New("user not found")
	}
	return &u, nil
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	for _, u := range db.Users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
