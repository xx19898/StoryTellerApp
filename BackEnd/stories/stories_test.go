package stories

import (
	"StoryTellerAppBackend/configuration"
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/middleware"
	"StoryTellerAppBackend/models"
	"bytes"
	"encoding/json"
	"fmt"
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
	story := models.Story{Title: "Test Title", Content: "Test Content", Username: "testuser"}
	databaselayer.CreateNewStory(story)
	storyFromDb := databaselayer.FindStoryByTitle("Test Title")
	assert.Equal(suite.T(), "Test Title", storyFromDb.Title)
}

func (suite *StoriesTestSuite) TestThatStoryCreatingPipelineWorks() {
	mockRouter := gin.Default()
	mockRouter.Use(middleware.UserInfoExtractionMiddleware())
	mockRouter.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))
	secret, _ := helpers.GetEnv("JWT_SECRET")
	accToken, _ := middleware.GenerateJWTToken("testuser", 1, []string{"ROLE_USER"}, secret, middleware.AccessToken)
	mockRouter.POST("/testCreatingNewStory", CreateStory)
	w := httptest.NewRecorder()
	mockStory := StoryDTO{Content: "<p>Test</p>", Title: "Test Title"}
	mockStoryJson, err := json.Marshal(mockStory)
	if err != nil {
		panic("Could not marshal mockStory")
	}
	mockRequest, _ := http.NewRequest("POST", "/testCreatingNewStory", bytes.NewBuffer(mockStoryJson))
	mockRequest.Header.Set("Authorization", accToken)
	mockRouter.ServeHTTP(w, mockRequest)
	fmt.Printf(w.Body.String())
	assert.Equal(suite.T(), 202, w.Code)
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
