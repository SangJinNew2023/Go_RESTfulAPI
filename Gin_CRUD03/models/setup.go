package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open("root:1111@tcp(localhost:3306)/go_crud"))
	if err != nil {
		panic(err)
	}
	database.AutoMigrate(&Product{})

	DB = database
}
