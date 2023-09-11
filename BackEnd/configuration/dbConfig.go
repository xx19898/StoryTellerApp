package configuration

import (
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDb(config *gorm.Config) {
	var err error
	host, _ := helpers.GetEnv("host")
	db_port, _ := helpers.GetEnv("db_port")
	user, _ := helpers.GetEnv("user")
	password, _ := helpers.GetEnv("password")
	dbname, _ := helpers.GetEnv("dbname")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin", host, user, password, dbname, db_port)
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
	host, _ := helpers.GetEnv("test_db_host")
	port, _ := helpers.GetEnv("test_db_PORT")
	user, _ := helpers.GetEnv("test_db_user")
	password, _ := helpers.GetEnv("test_db_password")
	dbname, _ := helpers.GetEnv("test_db_name")
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Europe/Berlin", host, user, password, dbname, port)
	DB, err = gorm.Open(postgres.Open(dsn), config)
	if err != nil {
		log.Fatal("Failed to connect to the Database")
	}
	fmt.Println("Connected Successfully to the Database")
	DB.AutoMigrate(&models.User{}, &models.Role{})
	fmt.Println("Connected Successfully to the Database")
}
