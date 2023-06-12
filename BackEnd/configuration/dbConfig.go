package configuration

import (
	"StoryTellerAppBackend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(config *gorm.Config) {
	var err error
	dsn := "host=localhost user=admin password=verySecurePassword dbname=tellastory_db port=1338 sslmode=disable TimeZone=Europe/Berlin"
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")
	DB.AutoMigrate(&models.User{})
}
