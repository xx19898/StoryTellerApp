package stories

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

type StoriesTestSuite struct {
	suite.Suite
}

func (suite *StoriesTestSuite) SetupSuite() {
	godotenv.Load("../.env")
	configuration.ConfigureDatabaseForTest()
	newRole := models.Role{Name: "USER"}
	databaselayer.CreateNewUser("testuser", "testpassword", "xx", []models.Role{newRole})
}

func (suite *StoriesTestSuite) TestThatCreatingNewStoryAndRetrievingItWorks() {
	user := databaselayer.FindUserByName("testuser")
	story := models.Story{Title: "Test Title", Content: "Test Content", Owner: models.User{ID: uint(user.ID)}}
	databaselayer.CreateNewStory(story)
	allStories := databaselayer.FindAllStories()
	fmt.Println(allStories)
	fmt.Println("Got here")
	storyFromDb := databaselayer.FindStoryByTitle("Test Title")
	assert.Equal(suite.T(), "Test Title", storyFromDb.Title)
}

//1) Implement some kind of role based authorization
//1) Test  that updating the story works
//2) Test that deleting the story works
//3) Test that adding a comment works
//4) Test that deleting a comment works
// 5) Test that comments reply to association works
// 6) Test that comments sender association works
// 7) Implement authorization and only allow changing users own stories


//TODO: test that the user is not able to update the story if he does not own it

func (suite *StoriesTestSuite)

func TestStoriesTestSuite(t *testing.T) {
	suite.Run(t, new(StoriesTestSuite))
}
