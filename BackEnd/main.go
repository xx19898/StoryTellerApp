package main

import (
	"StoryTellerAppBackend/configuration"
	"fmt"

	"gorm.io/gorm"
)

func main() {
	fmt.Println("hello")
	configuration.ConnectDb(&gorm.Config{})
}
