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

type StoriesEndpointTestSuite struct {
	suite.Suite
}

func (suite *StoriesEndpointTestSuite) SetupSuite() {
	godotenv.Load("../.env")
	configuration.ConfigureDatabaseForTest()
	newRole := models.Role{Name: "USER"}
	databaselayer.CreateNewUser("testuser", "testpassword", "xx", []models.Role{newRole})
}

func (suite *StoriesEndpointTestSuite) TestThatStoryCreatingPipelineWorks() {
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

	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *StoriesEndpointTestSuite) TestStoryUpdatingEndpoint() {
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

	var objmap map[string]json.RawMessage
	var createdStory models.Story

	json.Unmarshal(w.Body.Bytes(), &objmap)
	json.Unmarshal(objmap["newStory"], &createdStory)

	assert.Equal(suite.T(), 200, w.Code)
	assert.Equal(suite.T(), "Test Title", "Test Title")

	mockRouter.PATCH("/updateStory", UpdateStoryContent)
	w2 := httptest.NewRecorder()
	storyUpdate := StoryDTO{Content: "<p>Updated</p>", Title: "Test Title", ID: createdStory.ID}
	storyUpdateJson, _ := json.Marshal(storyUpdate)

	mockRequest2, _ := http.NewRequest("PATCH", "/updateStory", bytes.NewBuffer(storyUpdateJson))
	mockRequest2.Header.Set("Authorization", accToken)
	mockRouter.ServeHTTP(w2, mockRequest2)

	var objmap2 map[string]json.RawMessage
	var updatedStory models.Story

	json.Unmarshal(w2.Body.Bytes(), &objmap2)
	json.Unmarshal(objmap2["updatedStory"], &updatedStory)

	assert.Equal(suite.T(), 200, w2.Code)
	assert.Equal(suite.T(), "<p>Updated</p>", updatedStory.Content)
}

func (suite *StoriesEndpointTestSuite) TestThatStoryDeletionPipelineWorks() {
	newStory, err := databaselayer.CreateNewStory(
		"testuser",
		"Test Content",
		"Story To Delete",
	)

	if err != nil {
		panic("Error when creating new story to test deletion")
	}

	newStoryInDb, findCreatedStoryErr := databaselayer.FindStoryById(newStory.ID)
	assert.Equal(suite.T(), nil, findCreatedStoryErr)
	assert.Equal(suite.T(), "Story To Delete", newStoryInDb.Title)

	mockRouter := gin.Default()
	mockRouter.Use(middleware.UserInfoExtractionMiddleware())
	mockRouter.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))

	secret, _ := helpers.GetEnv("JWT_SECRET")
	accToken, _ := middleware.GenerateJWTToken("testuser", 1, []string{"ROLE_USER"}, secret, middleware.AccessToken)

	mockRouter.DELETE("/testDeletingStory", DeleteStory)

	//Deleting non-existing story
	nonExistingStory := StoryDTO{ID: 333399888}
	nonExistingStoryDeletionRequestRecorder := httptest.NewRecorder()
	nonExistStoryToDeleteJSON, errMarshallingNonExistStory := json.Marshal(&nonExistingStory)
	if errMarshallingNonExistStory != nil {
		panic("Error when marshalling non-existing story to delete")
	}
	nonExistStoryDeletionRequest, _ := http.NewRequest("DELETE", "/testDeletingStory", bytes.NewBuffer(nonExistStoryToDeleteJSON))
	nonExistStoryDeletionRequest.Header.Set("Authorization", accToken)
	mockRouter.ServeHTTP(nonExistingStoryDeletionRequestRecorder, nonExistStoryDeletionRequest)
	assert.Equal(suite.T(), 404, nonExistingStoryDeletionRequestRecorder.Code)

	//Trying to delete right story but with acc token with wrong username
	story := StoryDTO{ID: newStoryInDb.ID}
	notAuthrzdUserStoryDelReqRecorder := httptest.NewRecorder()
	storyJSON, errMarshallingNonExistStory := json.Marshal(&story)
	if errMarshallingNonExistStory != nil {
		panic("Error when marshalling non-existing story to delete")
	}
	notAuthrzdUserStoryDelReq, _ := http.NewRequest("DELETE", "/testDeletingStory", bytes.NewBuffer(storyJSON))
	accTokenWithAnotherUser, _ := middleware.GenerateJWTToken("wrongUser", 1, []string{"ROLE_USER"}, secret, middleware.AccessToken)
	notAuthrzdUserStoryDelReq.Header.Set("Authorization", accTokenWithAnotherUser)
	mockRouter.ServeHTTP(notAuthrzdUserStoryDelReqRecorder, notAuthrzdUserStoryDelReq)
	assert.Equal(suite.T(), 403, notAuthrzdUserStoryDelReqRecorder.Code)

	//Deleting right story with right user encoded into JWT
	w := httptest.NewRecorder()
	storyToDeleteJSON, err := json.Marshal(&newStoryInDb)
	fmt.Println("XXXXXXXXXXXXXXX")
	fmt.Println(newStoryInDb.ID)
	fmt.Println("XXXXXXXXXXXXXXX")
	if err != nil {
		panic("Error when marshalling story to delete")
	}
	mockRequest, _ := http.NewRequest("DELETE", "/testDeletingStory", bytes.NewBuffer(storyToDeleteJSON))
	mockRequest.Header.Set("Authorization", accToken)
	mockRouter.ServeHTTP(w, mockRequest)

	_, errSameStoryLookup := databaselayer.FindStoryById(newStoryInDb.ID)

	assert.NotNil(suite.T(), errSameStoryLookup)
	assert.Equal(suite.T(), 200, w.Code)
}

func TestStoriesEndpointsSuite(t *testing.T) {
	suite.Run(t, new(StoriesEndpointTestSuite))
}
