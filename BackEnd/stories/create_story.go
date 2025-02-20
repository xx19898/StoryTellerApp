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

	res, errWhenCreatingStory := databaselayer.CreateNewStory(
		c.GetString("LOGGED_USER_NAME"),
		newStoryUserInput.Content,
		newStoryUserInput.Title,
	)

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
//TODO: implement and test
func UpdateStory(c *gin.Context) {

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
