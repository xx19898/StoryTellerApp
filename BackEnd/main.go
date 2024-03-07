package main

import (
	"StoryTellerAppBackend/comments"
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/login_and_register"
	auth "StoryTellerAppBackend/login_and_register"
	"StoryTellerAppBackend/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
)

func main() {
	r := gin.Default()
	godotenv.Load(".env")
	authGroup := r.Group("/auth")

	authGroup.POST("/register", login_and_register.Register)
	authGroup.POST("/login", auth.Login)

	configuration.ConnectDb(&gorm.Config{})

	PORT, portIsFound := helpers.GetEnv("PORT")

	storiesGroup := r.Group("/stories")
	storiesGroup.Use(middleware.UserInfoExtractionMiddleware())
	storiesGroup.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))

	commentGroup := r.Group("/comments")
	commentGroup.Use(middleware.UserInfoExtractionMiddleware())
	commentGroup.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))
	commentGroup.POST("/comment", comments.CreateComment)
	commentGroup.GET("/commentsByStoryId", comments.GetCommentsByStoryId)

	imageGroup := r.Group("/images")
	imageGroup.Use(middleware.UserInfoExtractionMiddleware())
	imageGroup.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))
	//TODO: fix imports
	imageGroup.POST("/avatar", imagestorage.UploadUserAvatar)
	imageGroup.GET("/avatar", imagestorage.DownloadUserAvatar)

	if !portIsFound {
		panic("could not retrieve env variable with port")
	}
	r.Run(fmt.Sprintf(":%s", PORT))
}
