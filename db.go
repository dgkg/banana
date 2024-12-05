package main

import "errors"

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
