package users

import (
	"StoryTellerAppBackend/configuration"
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/models"
	"fmt"
	"testing"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type UsersTestSuite struct {
	suite.Suite
}

func (suite *UsersTestSuite) SetupSuite() {
	godotenv.Load("../.env")
	configuration.ConfigureDatabaseForTest()
	fmt.Println("Setting up Users Test Suite")
	newRole := models.Role{Name: "User"}
	databaselayer.CreateNewUser("testuser2", "testpassword", "testEmail@gmail.com", []models.Role{newRole})
}

func (suite *UsersTestSuite) TestRetrievingCreatedUser() {
	user, err := databaselayer.FindUserByName("testuser2")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "testEmail@gmail.com", user.Email)
}

func (suite *UsersTestSuite) TestRetrievingNonExistantUser() {
	_, err := databaselayer.FindUserByName("nonexistantuser")
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "record not found", err.Error())
}

func (suite *UsersTestSuite) TestRetrievingHashedPasswordForUser() {
	hash, err := databaselayer.FindUserPasswordHashByName("testuser2")
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "testpassword", hash)
}

func (suite *UsersTestSuite) TestRetrievingHashForNonExistantUser() {
	hash, err := databaselayer.FindUserPasswordHashByName("NonExistantUser")
	assert.NotNil(suite.T(), err)
	assert.Equal(suite.T(), "record not found", err.Error())
	assert.Equal(suite.T(), "", hash)
}

func TestUsersTestSuite(t *testing.T) {
	suite.Run(t, new(UsersTestSuite))
}
