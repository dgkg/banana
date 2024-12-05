package db

import (
	"gorm.io/gorm"
)

type SQLite struct {
	db *gorm.DB
}

// func NewSQLite(fileName string) *SQLite {
// 	gormDB, err := gorm.Open("sqlite3", fileName)
// 	return &SQLite{}
// }
