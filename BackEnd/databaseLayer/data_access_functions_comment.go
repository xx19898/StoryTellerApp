package databaselayer

import (
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/models"
)

func CreateNewComment(newComment models.Comment) (models.Comment, error) {
	result := configuration.DB.Create(&newComment)
	return newComment, result.Error
}

func GetCommentById(ID int) (models.Comment, error) {
	var comment models.Comment
	result := configuration.DB.Where(models.Comment{ID: uint(ID)}).First(&comment)
	return comment, result.Error
}

func GetCommentsByStoryId(storyId int) ([]models.Comment, error) {
	var comments []models.Comment
	result := configuration.DB.Where(models.Comment{StoryID: uint(storyId)}).Find(&comments)
	return comments, result.Error
}
