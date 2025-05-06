package stories

import (
	"StoryTellerAppBackend/configuration"
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/models"
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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
	storyUpdate := StoryDTO{Content: "<h1>Updated Title</h1>",ID: 1} 
	router := gin.Default()
	router.PUT("/updateStory", UpdateStoryContent)

	jsonStoryUpdate, _ := json.Marshal(storyUpdate)

	//TODO: gotta first obtain an jwt token
	req, _ := http.NewRequest("PUT","/update")
}