package stories

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateStory(c *gin.Context) {
	var newStory models.Story

	if err := c.ShouldBindJSON(&newStory); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed story",
		})
	}

	//TODO: save userId and username in c object in auth middleware and test it

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
