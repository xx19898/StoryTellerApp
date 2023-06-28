package main

import (
	"StoryTellerAppBackend/configuration"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	router := r.Group("/auth")

	router.POST("/register")
	router.POST("/login")
	fmt.Println("hello")
	godotenv.Load()
	configuration.ConnectDb(&gorm.Config{})
	r.Run()
}
