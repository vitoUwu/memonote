package database

import (
	"log"
	"memonote/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func CreateConnection() {
	db, err = gorm.Open(sqlite.Open("donotdelete.db"), &gorm.Config{})
	db.AutoMigrate(&models.User{}, &models.Note{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *gorm.DB {
	return db
}
