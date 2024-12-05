package main

import (
	"errors"

	"github.com/agnivade/levenshtein"
)

var db DB

var (
	ErrUserAllreadExists = errors.New("user email allready exists")
	ErrUserNotFound      = errors.New("user not found")
)

func init() {
	db = DB{
		Users: make(map[string]User),
	}
}

type DB struct {
	Users map[string]User
}

func NewDB() *DB {
	return &DB{
		Users: make(map[string]User),
	}
}

func (db *DB) SetUser(u *User) error {
	_, err := db.GetUserByEmail(u.Email)
	if err == nil {
		return ErrUserAllreadExists
	}
	db.Users[u.UUID] = *u
	return nil
}

func (db *DB) GetUserByID(uuid string) (*User, error) {
	u, ok := db.Users[uuid]
	if !ok {
		return nil, ErrUserNotFound
	}
	return &u, nil
}

func (db *DB) GetUserByEmail(email string) (*User, error) {
	for _, u := range db.Users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (db *DB) GetUserByName(name string) ([]User, error) {
	var res []User
	for _, u := range db.Users {
		distance := levenshtein.ComputeDistance(u.FirstName+" "+u.LastName, name)
		if distance > 3 {
			res = append(res, u)
		}
	}
	if len(res) > 0 {
		return res, nil
	}
	return nil, ErrUserNotFound
}

func (db *DB) GetAllUser() ([]User, error) {
	if len(db.Users) == 0 {
		return nil, ErrUserNotFound
	}
	var res []User
	for _, u := range db.Users {
		res = append(res, u)
	}
	return res, nil
}
