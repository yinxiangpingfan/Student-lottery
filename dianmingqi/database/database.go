package database

import (
	//gorm
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func OpenDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("student.db"))
	if err != nil {
		panic("failed to connect database")
	}
}
