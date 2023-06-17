package auth

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
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
	databaselayer.CreateNewUser(newUser.Name, newUser.Password, newUser.Email, []models.Role{userRole})
	c.JSON(http.StatusCreated, "User created!")
}

//TODO: first  create test for binding users data to go struct,
//TODO: then test that saving class to the database functions as it should,
//TODO:  then test it all as an integration
