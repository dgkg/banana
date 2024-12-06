package db

import (
	"errors"
	"fmt"
)

var (
	ErrUserAllreadExists = NewErroNotAuthorized("user", "email allready exists")
	ErrUserNotFound      = NewErroNotFound("user", "not found")
)

type ErrorDB struct {
	Err        error
	Entity     string
	Message    string
	StatusCode int
}

func (e *ErrorDB) Error() string {
	return fmt.Sprintf("Error: %v, Entity: %s, Message: %s, StatusCode: %d", e.Err, e.Entity, e.Message, e.StatusCode)
}

func NewErroNotFound(entity, message string) *ErrorDB {
	if len(message) == 0 {
		message = fmt.Sprintf("%s not found", entity)
	}
	return &ErrorDB{
		Err:        errors.New("not found"),
		Entity:     entity,
		Message:    message,
		StatusCode: 404,
	}
}

func NewErrorInternal(entity, message string, errOrigin error) *ErrorDB {
	if len(message) == 0 {
		message = fmt.Sprintf("%s internal serveur error", entity)
	}
	return &ErrorDB{
		Err:        errOrigin,
		Entity:     entity,
		Message:    message,
		StatusCode: 500,
	}
}

func NewErroNotAuthorized(entity, message string) *ErrorDB {
	if len(message) == 0 {
		message = fmt.Sprintf("%s not authorized", entity)
	}
	return &ErrorDB{
		Err:        errors.New("not authorized"),
		Entity:     entity,
		Message:    message,
		StatusCode: 401,
	}
}
