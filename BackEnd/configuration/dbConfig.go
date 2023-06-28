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
	DB.AutoMigrate(&models.User{}, &models.Role{})

	var count int64
	DB.Find(&models.Role{}).Count(&count)
	if count == 0 {
		userRole := models.Role{Name: "User"}
		adminRole := models.Role{Name: "Admin"}
		DB.Model(&models.Role{}).Create(&userRole)
		DB.Model(&models.Role{}).Create(&adminRole)
	}
}

func ConnectTestDb(config *gorm.Config) {
	var err error
	dsn := "host=localhost user=test_admin password=testDBPassword dbname=tellastory_test_db port=1339 sslmode=disable TimeZone=Europe/Berlin"
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")
	DB.AutoMigrate(&models.User{}, &models.Role{})
	fmt.Println("Connected Successfully to the Database")
}
