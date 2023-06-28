package auth

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginInfo models.User
	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed login request",
		})
	}

	passwordsMatch, error := verifyPassword(
		helpers.PasswordMatchesTheHash,
		databaselayer.FindUserPasswordHashByName,
		loginInfo.Name,
		loginInfo.Password)
	if error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Received error while verifying the password",
		})
	}
}
