package stories

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UpdateStoryContent(c *gin.Context) {
	var updatedStoryUserInput StoryDTO

	if err := c.ShouldBindJSON(&updatedStoryUserInput); err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed story object",
		})
		return
	}

	storyToUpdate, err := databaselayer.FindStoryById(updatedStoryUserInput.ID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "Could not find story with that id",
		})
		return
	}

	if storyToUpdate.Username != c.GetString("LOGGED_USER_NAME") {
		c.JSON(http.StatusMethodNotAllowed, gin.H{
			"error": "Trying to update story belonging to another user",
		})
		return
	}

	updatedStory, err := databaselayer.UpdateStoryContentById(storyToUpdate.ID, updatedStoryUserInput.Content)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updatedStory": updatedStory,
	})
	return
}
