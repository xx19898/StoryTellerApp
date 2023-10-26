package main

import (
	"StoryTellerAppBackend/auth"
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/helpers"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	godotenv.Load(".env")
	authGroup := r.Group("/auth")

	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)

	configuration.ConnectDb(&gorm.Config{})

	PORT, portIsFound := helpers.GetEnv("PORT")

	if !portIsFound {
		panic("could not retrieve env variable with port")
	}
	r.Run(fmt.Sprintf(":%s", PORT))
}
