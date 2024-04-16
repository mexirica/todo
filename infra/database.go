package infra

import (
	"github.com/mexirica/todo/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() {
	database, erro := gorm.Open(sqlite.Open("todo.db"), &gorm.Config{})

	if erro != nil {
		panic("Cannot connect to database!")
	}

	err := database.AutoMigrate(&models.Task{}, &models.User{})
	if err != nil {
		panic("Error migrating database!")
	}

	DB = database

}
