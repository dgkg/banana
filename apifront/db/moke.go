package db

import (
	"github.com/agnivade/levenshtein"

	"banana/apifront/model"
)

var _ DB = &Moke{}

type Moke struct {
	Users map[string]model.User
}

func NewMoke() *Moke {
	return &Moke{
		Users: make(map[string]model.User),
	}
}

func (db *Moke) SetUser(u *model.User) error {
	_, err := db.GetUserByEmail(u.Email)
	if err == nil {
		return ErrUserAllreadExists
	}
	db.Users[u.UUID] = *u
	return nil
}

func (db *Moke) GetUserByID(uuid string) (*model.User, error) {
	u, ok := db.Users[uuid]
	if !ok {
		return nil, ErrUserNotFound
	}
	return &u, nil
}

func (db *Moke) GetUserByEmail(email string) (*model.User, error) {
	for _, u := range db.Users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, ErrUserNotFound
}

func (db *Moke) GetUserByName(name string) ([]model.User, error) {
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

func (db *Moke) GetAllUser() ([]model.User, error) {
	if len(db.Users) == 0 {
		return nil, ErrUserNotFound
	}
	var res []model.User
	for _, u := range db.Users {
		res = append(res, u)
	}
	return res, nil
}
