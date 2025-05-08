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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find story with that id",
		}) 
		return
	}

	if storyToUpdate.Username != c.GetString("LOGGED_USER_NAME") {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Trying to update story belonging to another user",
		})
		return
	}

	if err = CheckStory(updatedStoryUserInput.Content); err != nil{
		fmt.Println(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{
			"error":"Story update is not approved as the story is malformed",
		})
	}

	updatedStory, err := databaselayer.UpdateStoryContentById(storyToUpdate.ID, updatedStoryUserInput.Content)

	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":"Error when trying to update story",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"updatedStory": updatedStory,
	})

	return
}
