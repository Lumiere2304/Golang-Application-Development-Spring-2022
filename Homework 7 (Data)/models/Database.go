package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDB() {
	connStr := "host = 127.0.0.1 port = 5432 user = postgres dbname = practice11 password = root sslmode = disable"
	db, err := gorm.Open("postgres", connStr)
	if err != nil {
		panic("Не удается подключится к базе")
	}

	db.AutoMigrate(&User{})
	DB = db
}
