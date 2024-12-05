package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"time"

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
	return "SSSSS"
}

func (p Password) MarshalJSON() ([]byte, error) {
	var s string = "JJJJJJ"
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
