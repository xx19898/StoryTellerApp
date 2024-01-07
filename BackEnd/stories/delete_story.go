package stories

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteStory(c *gin.Context) {
	userBehindRequest := c.GetString("LOGGED_USER_NAME")

	fmt.Println(userBehindRequest)

	var storyToDelete StoryDTO

	if errUnmarshalling := c.ShouldBindJSON(&storyToDelete); errUnmarshalling != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed delete request",
		})
		return
	}

	storyFromDb, errStoryLookup := databaselayer.FindStoryById(storyToDelete.ID)

	if errStoryLookup != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"error": "No story to delete with that ID",
		})
		return
	}

	if storyFromDb.Username != userBehindRequest {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Not allowed. You can only delete your own stories",
		})
		return
	}

	_, errStoryDeletion := databaselayer.DeleteStory(storyFromDb)

	if errStoryDeletion != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Error when trying to delete the story",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Story deleted",
	})
}
