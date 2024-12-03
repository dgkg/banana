package model

import (
	"context"
	"fmt"

	"github.com/wearemojo/mojo-public-go/lib/ksuid"
)

// User represents a user in the system.
type User struct {
	ID        string // Unique identifier for the user.
	Name      string // Name of the user.
	BirthDate string // Birth date of the user.
	CreatedAt string // Timestamp when the user was created.
	UpdatedAt string // Timestamp when the user was last updated.
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
