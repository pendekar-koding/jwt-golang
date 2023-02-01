package database

import (
	"auth-golang/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

var Instance *gorm.DB
var dbError error

func Connect(connectionString string) {
	Instance, dbError = gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if dbError != nil {
		log.Fatal(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connect to DB")
}

func Migrate() {
	Instance.AutoMigrate(&models.User{})
	log.Println("Database migration Completed!")
}
