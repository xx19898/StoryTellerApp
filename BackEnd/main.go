package main

import (
	"StoryTellerAppBackend/configuration"
	"fmt"

	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	godotenv.Load()
	fmt.Println("hello")
	configuration.ConnectDb(&gorm.Config{})
}
