package main

import (
	"StoryTellerAppBackend/auth"
	"StoryTellerAppBackend/configuration"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	authGroup := r.Group("/auth")

	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)

	godotenv.Load()
	configuration.ConnectDb(&gorm.Config{})
	r.Run()
}
