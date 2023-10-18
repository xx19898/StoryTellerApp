package stories

import (
	"StoryTellerAppBackend/configuration"
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
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
	storyFromDb := databaselayer.FindStoryByTitle("Test Title")
	assert.Equal(suite.T(), "Test Title", storyFromDb.Title)
}

func (suite *StoriesTestSuite) TestThatStoryCreatingPipelineWorks() {
	mockRouter := gin.Default()
	//TODO: get jwt access token and put it into the mockrequest
	mockRouter.POST("/testCreatingNewStory", CreateStory)

	w := httptest.NewRecorder()
	mockRequest, _ := http.NewRequest("POST", "/testCreatingNewStory")
}

// 1) Implement some kind of role based authorization DONE
// 1) Test  that updating the story works TODO!!!
// 2) Test that deleting the story works
// 3) Test that adding a comment works
// 4) Test that deleting a comment works
// 5) Test that comments reply to association works
// 6) Test that comments sender association works
// 7) Implement authorization and only allow changing users own stories

// TODO: Test that the user is able to create a story

func TestStoriesTestSuite(t *testing.T) {
	suite.Run(t, new(StoriesTestSuite))
}
