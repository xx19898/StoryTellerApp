package stories

import (
	"StoryTellerAppBackend/configuration"
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/middleware"
	"StoryTellerAppBackend/models"
	"bytes"
	"encoding/json"
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

func (suite *StoriesTestSuite) TestThatCreatingNewStoryAndRetrievingItInDatabaseLayerWorks() {
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

	assert.Equal(suite.T(), 200, w.Code)
}

func (suite *StoriesTestSuite) TestThatSearchingAfterStoryWithNonExistingIDGetsHandledProperly() {
	_, err := databaselayer.FindStoryById(9999)
	assert.NotEqual(suite.T(), err, nil)
}

func (suite *StoriesTestSuite) TestThatUpdatingStoryWithNonExistingIDGetsHandledProperly() {
	_, err := databaselayer.UpdateStoryContentById(999, "wsww")
	assert.Equal(suite.T(), err.Error(), "record not found")
}

func (suite *StoriesTestSuite) TestThatUpdatingStoryInDatalayerWorksProperly() {
	createdStory, _ := databaselayer.CreateNewStory(
		models.Story{
			Username: "testuser",
			Title:    "Test Title",
			Content:  "content1",
		})

	databaselayer.UpdateStoryContentById(createdStory.ID, "content2")

	updatedStory, _ := databaselayer.FindStoryById(createdStory.ID)

	assert.Equal(suite.T(), updatedStory.Content, "content2")
}

func (suite *StoriesTestSuite) TestThatStoryUpdatingPipelineWorks() {
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

// 3) Test that adding a comment works
// 4) Test that deleting a comment works
// 5) Test that comments reply to association works
// 6) Test that comments sender association works
// 7) Implement authorization and only allow changing user's own stories

func (suite *StoriesTestSuite) TestThatDatabaselayersDeleteStoryFunctionWorks() {
	newStory, err := databaselayer.CreateNewStory(models.Story{Content: "Test Content", Title: "Story To Delete", Username: "testuser"})

	if err != nil {
		panic("Error when creating new story to test deletion")
	}

	newStoryInDb, findCreatedStoryErr := databaselayer.FindStoryById(newStory.ID)
	assert.Equal(suite.T(), nil, findCreatedStoryErr)
	assert.Equal(suite.T(), "Story To Delete", newStoryInDb.Title)

	databaselayer.DeleteStory(newStory)
	_, findStoryErr := databaselayer.FindStoryById(newStory.ID)

	assert.NotEqual(suite.T(), nil, findStoryErr)
	assert.Equal(suite.T(), "record not found", findStoryErr.Error())
}

func (suite *StoriesTestSuite) TestThatStoryDeletionPipelineWorks() {
	newStory, err := databaselayer.CreateNewStory(models.Story{Content: "Test Content", Title: "Story To Delete", Username: "testuser"})

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
	if err != nil {
		panic("Error when marshalling story to delete")
	}
	mockRequest, _ := http.NewRequest("DELETE", "/testDeletingStory", bytes.NewBuffer(storyToDeleteJSON))
	mockRequest.Header.Set("Authorization", accToken)
	mockRouter.ServeHTTP(w, mockRequest)

	_, errSameStoryLookup := databaselayer.FindStoryById(newStory.ID)

	assert.NotNil(suite.T(), errSameStoryLookup)
	assert.Equal(suite.T(), 200, w.Code)
}

func TestStoriesTestSuite(t *testing.T) {
	suite.Run(t, new(StoriesTestSuite))
}
