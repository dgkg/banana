package db

import (
	"banana/model"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var _ DB = &SQLite{}

type SQLite struct {
	con *gorm.DB
}

func NewSQLite(fileName string) *SQLite {
	dbconn, err := gorm.Open(sqlite.Open(fileName), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	dbconn.AutoMigrate(&model.User{})

	return &SQLite{
		con: dbconn,
	}
}

func (db *SQLite) SetUser(u *model.User) error {
	return db.con.Debug().Create(u).Error
}

func (db *SQLite) GetUserByID(uuid string) (*model.User, error) {
	var u model.User
	err := db.con.Debug().Model(&u).Where("uuid = ?", uuid).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *SQLite) GetUserByEmail(email string) (*model.User, error) {
	var u model.User
	err := db.con.Debug().Table("users").Where("email = ?", email).First(&u).Error
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (db *SQLite) GetUserByName(name string) ([]model.User, error) {
	var us []model.User
	err := db.con.Debug().Table("users").Where("last_name LIKE ?", name).Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}

func (db *SQLite) GetAllUser() ([]model.User, error) {
	var us []model.User
	err := db.con.Debug().Table("users").Find(&us).Error
	if err != nil {
		return nil, err
	}
	return us, nil
}
