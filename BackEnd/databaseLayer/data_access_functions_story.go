package databaselayer

import (
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/models"
)

func CreateNewStory(story models.Story) (models.Story, error) {
	result := configuration.DB.Create(&story)
	return story, result.Error
}

func FindStoryById(ID uint) models.Story {
	var story models.Story
	configuration.DB.Where(models.Story{ID: ID}).First(&story)
	return story
}

func FindAllStories() []models.Story {
	var stories []models.Story
	configuration.DB.Find(&stories)
	return stories
}

func FindStoryByTitle(title string) models.Story {
	var story models.Story
	configuration.DB.Where(models.Story{Title: title}).First(&story)
	return story
}
