package config

import (
	"fmt"
	"log"
	"os"
	"warlock-backend/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		host, user, password, dbname, port)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	database.AutoMigrate(&models.User{})
	database.AutoMigrate(&models.Role{})
	database.AutoMigrate(&models.Qa{})
	database.AutoMigrate(&models.Document{})
	database.AutoMigrate(&models.Subject{})
	database.AutoMigrate(&models.Topic{})
	SeedRoles(database)
	SeedSubjects(database)

	DB = database

}
