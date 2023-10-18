package stories

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context) {
	var newStoryUserInput StoryDTO
	var newStory models.Story
	//TODO: /For later/ probably should sanitize the html input in the new stories

	if err := c.ShouldBindJSON(&newStoryUserInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed story",
		})
	}

	res, errWhenCreatingStory := databaselayer.CreateNewStory(newStory)

	if errWhenCreatingStory != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Error when creating new story",
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"newStory": res,
	})
}

func UpdateStory(c *gin.Context) {
	//TODO: make this an endpoint which creates/updates the story

	var newStory models.Story

	if err := c.ShouldBindJSON(&newStory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed story",
		})
	}

	res, errWhenCreatingStory := databaselayer.CreateNewStory(newStory)

	if errWhenCreatingStory != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error when creating new story",
		})
	}

	c.JSON(http.StatusAccepted, gin.H{
		"message": "New Story created",
		"story":   res,
	})
}
