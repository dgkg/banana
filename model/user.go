package model

import (
	"context"

	"github.com/wearemojo/mojo-public-go/lib/ksuid"
)

type User struct {
	ID        string
	Name      string
	BirthDate string
	CreatedAt string
	UpdatedAt string
}

func CreateUserWithPtr(name string) *User {
	ctx := context.Background()
	sk := ksuid.Generate(ctx, "user")
	return &User{
		ID:   sk.String(),
		Name: name,
	}
}

func CreateUserByValue(name string) User {
	return User{
		Name: name,
	}
}
