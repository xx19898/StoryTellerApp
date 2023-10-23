package comments

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateComment(c *gin.Context) {
	var newCommentUserInput CommentDTO

	if errDeserializingUserInput := c.ShouldBindJSON(&newCommentUserInput); errDeserializingUserInput != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed comment object",
		})
		return
	}

	userBehindComment := c.GetString("LOGGED_USER_NAME")

	newComment := models.Comment{
		TextContent: newCommentUserInput.TextContent,
		Username:    userBehindComment,
		StoryID:     newCommentUserInput.StoryID,
	}

	createdComment, errComCreation := databaselayer.CreateNewComment(newComment)

	if errComCreation != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed comment object",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"newComment": createdComment,
	})
}
