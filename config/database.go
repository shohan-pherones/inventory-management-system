package config

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabse() {
	database, err := gorm.Open(sqlite.Open("inventory.db"), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	DB = database

	log.Println("Database connection estiblished.")
}
