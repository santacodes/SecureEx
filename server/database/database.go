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

var DB *gorm.DB

func DBMigrate() *gorm.DB {
	DB, err := gorm.Open(sqlite.Open("web.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to db", err)
	}
	log.Println("Connected to DB")

	DB.AutoMigrate(&Website{})
	return DB
}

func AlreadyCached(domain string) (isChecked bool, isSafe bool) {
	var website Website
	DB.Where("website = ?", domain).First(&website)
	if website.Website == "" {
		return false, true
	}
	return true, website.IsSafe
}
