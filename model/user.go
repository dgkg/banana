package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/wearemojo/mojo-public-go/lib/ksuid"
)

// User represents a user in the system.
type User struct {
	ID   string `json:"id"`   // Unique identifier for the user.
	Name string `json:"name"` // Name of the user.
	Human
	LastConnected time.Time `json:"last_connected"`
	CreatedAt     time.Time `json:"created_at"` // Timestamp when the user was created.
	UpdatedAt     time.Time `json:"updated_at"` // Timestamp when the user was last updated.
	Errs          []error
}

func NewUser(fn, ln string) *User {
	ctx := context.Background()
	sk := ksuid.Generate(ctx, "user")
	return &User{
		ID:        sk.String(),
		Name:      fn + " " + ln,
		CreatedAt: time.Now(),
	}
}

func (u *User) SetBirthDate(bd time.Time) *User {
	if bd.Unix() > time.Now().Unix() {
		u.Errs = append(u.Errs, errors.New("try to set a birthdate that is not allowed"))
	}
	u.BirthDate = bd
	return u
}

func (u *User) SetLastConnected(lastDate time.Time) *User {
	if len(u.Errs) > 0 {
		return u
	}
	u.LastConnected = lastDate
	return u
}

func (u *User) Error() []error {
	return u.Errs
}

type Human struct {
	BirthDate time.Time // Birth date of the user.
}

var u User = User{
	ID:   "1",
	Name: "toto",
	Human: Human{
		BirthDate: time.Now(),
	},
}

// CreateUserWithPtr creates a new user with the given name and returns a pointer to the User.
// It generates a unique ID for the user using ksuid.
func CreateUserWithPtr(name string) *User {
	ctx := context.Background()
	sk := ksuid.Generate(ctx, "user")
	return &User{
		ID:   sk.String(),
		Name: name,
	}
}

// CreateUserByValue creates a new user with the given name and returns the User by value.
func CreateUserByValue(name string) User {
	ctx := context.Background()
	sk := ksuid.Generate(ctx, "user")
	return User{
		ID:   sk.String(),
		Name: name,
	}
}

func GenerateUsers(n int) map[string]User {
	users := make(map[string]User)
	for i := 0; i < n; i++ {
		ctx := context.Background()
		sk := ksuid.Generate(ctx, "user")
		u := User{
			ID:   sk.String(),
			Name: "User-" + fmt.Sprintf("%v", i),
		}
		users[sk.String()] = u
	}
	return users
}
