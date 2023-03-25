package database

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Website struct {
	Website string
	IsSafe  bool
}

func DBMigrate() {
	db, err := gorm.Open(sqlite.Open("web.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
	log.Println("Connected to DB")

	db.AutoMigrate(&Website{})
}
