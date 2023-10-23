package comments

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetCommentsByStoryId(c *gin.Context) {
	storyIdString := c.Query("storyId")

	storyId, storyIdConvErr := strconv.Atoi(storyIdString)
	if storyIdConvErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Story id parameter is in wrong form",
		})
		return
	}

	story, storyRetrievalErr := databaselayer.FindStoryById(uint(storyId))
	if storyRetrievalErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find story with requested story id",
		})
		return
	}

	comments := story.Comments

	c.JSON(
		http.StatusOK,
		gin.H{
			"comments": comments,
		},
	)
	return
}
