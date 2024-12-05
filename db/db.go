package db

import (
	"banana/model"

	"github.com/agnivade/levenshtein"
)

var (
	ErrUserAllreadExists = NewErroNotAuthorized("user", "email allready exists")
	ErrUserNotFound      = NewErroNotFound("user", "not found")
)

type DB struct {
	Users map[string]model.User
}

func NewDB() *DB {
	return &DB{
		Users: make(map[string]model.User),
	}
}

func (db *DB) SetUser(u *model.User) error {
	_, err := db.GetUserByEmail(u.Email)
	if err == nil {
		return ErrUserAllreadExists
	}
	db.Users[u.UUID] = *u
	return nil
}

func (db *DB) GetUserByID(uuid string) (*model.User, error) {
	u, ok := db.Users[uuid]
	if !ok {
		return nil, ErrUserNotFound
	}
	return &u, nil
}

func (db *DB) GetUserByEmail(email string) (*model.User, error) {
	for _, u := range db.Users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (db *DB) GetUserByName(name string) ([]model.User, error) {
	var res []model.User
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

func (db *DB) GetAllUser() ([]model.User, error) {
	if len(db.Users) == 0 {
		return nil, ErrUserNotFound
	}
	var res []model.User
	for _, u := range db.Users {
		res = append(res, u)
	}
	return res, nil
}
