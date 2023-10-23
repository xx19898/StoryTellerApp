package comments

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

type CommentsTestSuite struct {
	suite.Suite
}

func (suite *CommentsTestSuite) SetupSuite() {
	godotenv.Load("../.env")
	configuration.ConfigureDatabaseForTest()
	fmt.Println("Setting up Comments Test Suite")
	newRole := models.Role{Name: "User"}
	databaselayer.CreateNewUser("testuser", "testpassword", "testEmail@gmail.com", []models.Role{newRole})
	databaselayer.CreateNewStory(
		models.Story{
			ID:       1,
			Content:  "Test Story Content",
			Title:    "Test Story Title",
			Username: "testuser",
		})
}

func (suite *CommentsTestSuite) TearDownTest() {
	fmt.Println("Cleaning Up Comments after test")

}

func (suite *CommentsTestSuite) TestCommentCreation() {
	commentToCreate := models.Comment{
		TextContent: "Test Comment",
		Username:    "testuser",
		StoryID:     1,
	}
	newComment, err := databaselayer.CreateNewComment(commentToCreate)

	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Test Comment", newComment.TextContent)
}

func (suite *CommentsTestSuite) TestCommentRetrievalById() {
	createdComment := models.Comment{
		ID:          1,
		TextContent: "Test Comment",
		Username:    "testuser",
		StoryID:     1,
	}

	_, err := databaselayer.CreateNewComment(createdComment)
	assert.Nil(suite.T(), err)

	commentFromDb, err := databaselayer.GetCommentById(1)
	assert.Equal(suite.T(), "Test Comment", commentFromDb.TextContent)
}

func (suite *CommentsTestSuite) TestCommentCreationEndpoint() {
	//TODO: finish this
}

func TestCommentsTestSuite(t *testing.T) {
	suite.Run(t, new(CommentsTestSuite))
}
