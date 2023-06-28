package auth

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	userRole := databaselayer.FindRoleByName("User")
	encryptedPassword, err := helpers.EncryptPassword(newUser.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Error when trying to encrypt the password",
		})
	}
	databaselayer.CreateNewUser(newUser.Name, string(encryptedPassword), newUser.Email, []models.Role{userRole})
	c.JSON(http.StatusCreated, "User created!")
}
