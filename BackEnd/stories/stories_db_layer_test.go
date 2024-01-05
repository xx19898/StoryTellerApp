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

func (suite *StoriesDBLayerTestSuite) TearDownSuite() {
	configuration.ResetCommentsAndStories()
	fmt.Println("RESETTING STORIES AND COMMENTS!")
}

func (suite *StoriesDBLayerTestSuite) TestRetrievingStoryFromDbById() {
	fmt.Println("TEST ONE")
	story, err := databaselayer.FindStoryById(1)
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Test Story Title", story.Title)
}

func (suite *StoriesDBLayerTestSuite) TestThatCreatingNewStoryAndRetrievingItInDatabaseLayerWorks() {
	story := models.Story{Title: "Test Title", Content: "Test Content", Username: "testuser"}
	databaselayer.CreateNewStory(story.Username, story.Content, story.Title)
	storyFromDb := databaselayer.FindStoryByTitle("Test Title")
	assert.Equal(suite.T(), "Test Title", storyFromDb.Title)
}

func (suite *StoriesDBLayerTestSuite) TestRetrievingAllStories() {
	fmt.Println("TEST 2")
	stories := databaselayer.FindAllStories()
	assert.Equal(suite.T(), 2, len(stories))
}

func (suite *StoriesDBLayerTestSuite) TestChangingContentOfAStory() {
	fmt.Println("TEST 3")

}

func (suite *StoriesDBLayerTestSuite) TestThatSearchingAfterStoryWithNonExistingIDGetsHandledProperly() {
	_, err := databaselayer.FindStoryById(9999)
	assert.NotEqual(suite.T(), err, nil)
}

func (suite *StoriesDBLayerTestSuite) TestThatUpdatingStoryWithNonExistingIDGetsHandledProperly() {
	_, err := databaselayer.UpdateStoryContentById(999, "wsww")
	assert.Equal(suite.T(), err.Error(), "record not found")
}

func (suite *StoriesDBLayerTestSuite) TestThatUpdatingStoryInDatalayerWorksProperly() {
	createdStory, _ := databaselayer.CreateNewStory("testuser", "content1", "Test Title")
	databaselayer.UpdateStoryContentById(createdStory.ID, "content2")
	updatedStory, _ := databaselayer.FindStoryById(createdStory.ID)
	assert.Equal(suite.T(), updatedStory.Content, "content2")
}

func (suite *StoriesDBLayerTestSuite) TestThatDatabaselayersDeleteStoryFunctionWorks() {
	newStory, err := databaselayer.CreateNewStory("testuser", "Test Content", "Story To Delete")

	if err != nil {
		panic("Error when creating new story to test deletion")
	}

	newStoryInDb, findCreatedStoryErr := databaselayer.FindStoryById(newStory.ID)
	assert.Equal(suite.T(), nil, findCreatedStoryErr)
	assert.Equal(suite.T(), "Story To Delete", newStoryInDb.Title)

	databaselayer.DeleteStory(newStory)
	_, findStoryErr := databaselayer.FindStoryById(newStory.ID)

	assert.NotNil(suite.T(), findStoryErr)
	assert.Equal(suite.T(), "record not found", findStoryErr.Error())
}

func TestStoriesDBLayerFuncsSuite(t *testing.T) {
	suite.Run(t, new(StoriesDBLayerTestSuite))
}
