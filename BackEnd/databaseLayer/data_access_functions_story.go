package databaselayer

import (
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/models"
)

func CreateNewStory(story models.Story) (models.Story, error) {
	result := configuration.DB.Create(&story)
	return story, result.Error
}

func UpdateStoryContentById(ID uint, newContent string) (models.Story, error) {
	storyToUpdate, err := FindStoryById(ID)
	if err != nil {
		return models.Story{}, err
	}
	storyToUpdate.Content = newContent
	configuration.DB.Save(&storyToUpdate)
	return storyToUpdate, nil
}

func FindStoryById(ID uint) (models.Story, error) {
	var story models.Story
	result := configuration.DB.Where(models.Story{ID: ID}).First(&story)
	if result.Error != nil {
		return models.Story{}, result.Error
	}
	return story, nil
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
