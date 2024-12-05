package db

import "banana/apifront/model"

type DB interface {
	SetUser(u *model.User) error
	GetUserByID(uuid string) (*model.User, error)
	GetUserByEmail(email string) (*model.User, error)
	GetUserByName(name string) ([]model.User, error)
	GetAllUser() ([]model.User, error)
}
