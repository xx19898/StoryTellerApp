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

type StoriesDBLayerTestSuite struct {
	suite.Suite
}

func (suite *StoriesDBLayerTestSuite) SetupSuite() {
	godotenv.Load("../.env")
	configuration.ConfigureDatabaseForTest()
	fmt.Println("Setting up Stories Test Suite")
	newRole := models.Role{Name: "User"}
	databaselayer.CreateNewUser("testuser", "testpassword", "testEmail@gmail.com", []models.Role{newRole})
	databaselayer.CreateNewStoryCustomId("testuser", 1, "Test Story Content", "Test Story Title")
	databaselayer.CreateNewStoryCustomId("testuser", 2, "Test Story Content2", "Test Story Title2")
}

func (suite *StoriesDBLayerTestSuite) TearDownTest() {
	configuration.ResetCommentsAndStories()
}

func (suite *StoriesDBLayerTestSuite) TestRetrievingStoryFromDbById() {

	stories := databaselayer.FindAllStories()

	fmt.Println("********")
	fmt.Println(stories)
	fmt.Println("********")

	story, err := databaselayer.FindStoryById(1)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Test Story Title", story.Title)
}

func (suite *StoriesDBLayerTestSuite) TestRetrievingAllStories() {
	stories := databaselayer.FindAllStories()
	assert.Equal(suite.T(), 2, len(stories))
}

func TestStoriesDBLayerFuncsSuite(t *testing.T) {
	suite.Run(t, new(StoriesDBLayerTestSuite))
}
