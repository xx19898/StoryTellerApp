package stories

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context) {
	var newStoryUserInput StoryDTO

	//TODO: /For later/ probably should sanitize the html input in the new stories

	if err := c.ShouldBindJSON(&newStoryUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed story",
		})
		return
	}

	newStory := models.Story{Username: c.GetString("LOGGED_USER_NAME"), Content: newStoryUserInput.Content, Title: newStoryUserInput.Title}

	res, errWhenCreatingStory := databaselayer.CreateNewStory(newStory)

	if errWhenCreatingStory != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error when creating new story",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"newStory": res,
	})
	return
}

func UpdateStory(c *gin.Context) {
	//TODO: make this an endpoint which updates an already created thestory

	var newStory models.Story

	if err := c.ShouldBindJSON(&newStory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed story",
		})
		return
	}

	res, errWhenCreatingStory := databaselayer.CreateNewStory(newStory)

	if errWhenCreatingStory != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error when creating new story",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "New Story created",
		"story":   res,
	})
	return
}
