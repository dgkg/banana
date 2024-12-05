package model

import (
	"time"

	"github.com/google/uuid"
)

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
