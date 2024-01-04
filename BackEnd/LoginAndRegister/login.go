package auth

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var loginInfo models.User

	if err := c.ShouldBindJSON(&loginInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Malformed login request",
		})
		return
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
		return
	}

	if !passwordsMatch {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Either password or username are not correct. Please check entered credentials",
		})
		return
	}

	user := databaselayer.FindUserByName(loginInfo.Name)

	jwtToken, err := helpers.CreateToken(user.Name, int64(user.ID), helpers.RolesToString(user.Roles), time.Now().Add(time.Hour*time.Duration(2)).Unix(), time.Now().Unix())

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusAccepted, gin.H{
			"token": jwtToken,
		})
		return
	}
}
