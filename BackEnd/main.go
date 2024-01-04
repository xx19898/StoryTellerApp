package main

import (
	"StoryTellerAppBackend/auth"
	"StoryTellerAppBackend/comments"
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/helpers"
	imagestorage "StoryTellerAppBackend/imageStorage"
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

	authGroup.POST("/register", auth.Register)
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
	imageGroup.POST("/avatar", imagestorage.UploadUserAvatar)
	imageGroup.GET("/avatar", imagestorage.DownloadUserAvatar)

	if !portIsFound {
		panic("could not retrieve env variable with port")
	}
	r.Run(fmt.Sprintf(":%s", PORT))
}
