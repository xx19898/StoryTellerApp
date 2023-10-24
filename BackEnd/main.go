package main

import (
	"StoryTellerAppBackend/auth"
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/helpers"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()

	authGroup := r.Group("/auth")

	authGroup.POST("/register", auth.Register)
	authGroup.POST("/login", auth.Login)

	configuration.ConnectDb(&gorm.Config{})

	PORT, ok := helpers.GetEnv("PORT")
	fmt.Println("This is secret " + PORT)
	if !ok {
		panic("could not retrieve env variable with port")
	}
	r.Run(fmt.Sprintf(":%s", PORT))
}
