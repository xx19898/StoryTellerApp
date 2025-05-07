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

type StoryUpdateTestSuite struct {
	suite.Suite
}

func (suite *StoryUpdateTestSuite) SetupSuite(){
	godotenv.Load("../.env")
	configuration.ConfigureDatabaseForTest()
}

func (suite *StoryUpdateTestSuite) SetupTest() {
	configuration.ResetEverythingElseExceptRoles()
	databaselayer.CreateNewUser("TestUser","testPassword","testemail@gmail.com",[]models.Role{{Name:"USER"}})
	databaselayer.CreateNewStory("TestUser","<h1>Hello World</h1>","Test Story")
}

func (suite *StoryUpdateTestSuite) AfterTest(){
	configuration.ResetEverythingElseExceptRoles()
}

func (suite *StoryUpdateTestSuite) TestUpdatingCorrectStory(){
	story := databaselayer.FindStoryByTitle("Test Story")
	storyUpdate := StoryDTO{Content: "<h1>Updated Content</h1>",ID: story.ID} 
	
	router := gin.Default()
	router.Use(middleware.UserInfoExtractionMiddleware())
	router.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))
	router.PUT("/updateStory", UpdateStoryContent)

	reqRecorder := httptest.NewRecorder()

	jsonStoryUpdate, _ := json.Marshal(storyUpdate)

	godotenv.Load("../.env")
	secret, _ := helpers.GetEnv("JWT_SECRET")

	accToken, _ := middleware.GenerateJWTToken("TestUser", 1, []string{"ROLE_USER"}, secret, middleware.AccessToken)

	storyUpdateReq, _ := http.NewRequest("PUT","/updateStory",bytes.NewBuffer(jsonStoryUpdate))
	storyUpdateReq.Header.Add("Authorization", accToken)

	router.ServeHTTP(reqRecorder,storyUpdateReq)

	updatedStory := databaselayer.FindStoryByTitle("Test Story")
	
	assert.Equal(suite.T(),reqRecorder.Code,200)
	assert.Equal(suite.T(),updatedStory.Content,"<h1>Updated Content</h1>")
}

//TODO: testing updating story with nonexistant story,
// updating story which user is not author of 
// updating story with wrong content

func TestUpdateStoryTestSuite(t *testing.T){
	suite.Run(t, new(StoryUpdateTestSuite))
}