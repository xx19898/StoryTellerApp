package comments

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
	configuration.ResetComments()
}

func setupRouterAndGetToken() (*gin.Engine, string) {
	mockRouter := gin.Default()
	mockRouter.Use(middleware.UserInfoExtractionMiddleware())
	mockRouter.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))

	secret, _ := helpers.GetEnv("JWT_SECRET")
	accToken, _ := middleware.GenerateJWTToken("testuser", 1, []string{"ROLE_USER"}, secret, middleware.AccessToken)

	return mockRouter, accToken
}

func commentToJson(comment models.Comment) []byte {
	commentAsJSON, errMarshalling := json.Marshal(&comment)
	if errMarshalling != nil {
		panic("Could not marshal comment to json")
	}
	return commentAsJSON
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
	assert.Nil(suite.T(), err)
	assert.Equal(suite.T(), "Test Comment", commentFromDb.TextContent)
}

func (suite *CommentsTestSuite) TestCommentCreationEndpoint() {
	commentToCreate := models.Comment{
		ID:          2,
		TextContent: "Test Comment2",
		Username:    "testuser",
		StoryID:     1,
	}

	mockRouter, accToken := setupRouterAndGetToken()
	commentJSON := commentToJson(commentToCreate)
	comCreatingRecorder := httptest.NewRecorder()

	mockRouter.POST("/postComment", CreateComment)
	commentCreationRequest, _ := http.NewRequest("POST", "/postComment", bytes.NewBuffer(commentJSON))
	commentCreationRequest.Header.Set("Authorization", accToken)

	mockRouter.ServeHTTP(comCreatingRecorder, commentCreationRequest)

	assert.Equal(suite.T(), 200, comCreatingRecorder.Code)

	var commentFromServer models.Comment

	var objmap map[string]json.RawMessage

	json.Unmarshal(comCreatingRecorder.Body.Bytes(), &objmap)
	json.Unmarshal(objmap["newComment"], &commentFromServer)

	assert.Equal(suite.T(), "Test Comment2", commentFromServer.TextContent)
}

func (suite *CommentsTestSuite) TestCommentsRetrievalByStoryId() {
	commentToCreate1 := models.Comment{
		ID:          2,
		TextContent: "Test Comment2",
		Username:    "testuser",
		StoryID:     1,
	}

	commentToCreate2 := models.Comment{
		ID:          3,
		TextContent: "Test Comment3",
		Username:    "testuser",
		StoryID:     1,
	}

	databaselayer.CreateNewComment(commentToCreate1)
	databaselayer.CreateNewComment(commentToCreate2)

	mockRouter := gin.Default()

	commentsReceivingRecorder := httptest.NewRecorder()

	mockRouter.GET("/getComments", GetCommentsByStoryId)
	commentRetrievingRequest, _ := http.NewRequest("GET", "/getComments?storyId=1", bytes.NewBuffer([]byte{}))

	mockRouter.ServeHTTP(commentsReceivingRecorder, commentRetrievingRequest)

	assert.Equal(suite.T(), 200, commentsReceivingRecorder.Code)

	var commentsFromServer []CommentDTO

	var objmap map[string]json.RawMessage

	json.Unmarshal(commentsReceivingRecorder.Body.Bytes(), &objmap)

	json.Unmarshal(objmap["comments"], &commentsFromServer)

	fmt.Println("----------")
	fmt.Println(commentsFromServer)
	fmt.Println("----------")

	var firstComment CommentDTO
	var secondComment CommentDTO

	if commentsFromServer[0].ID == 2 {
		firstComment = commentsFromServer[0]
		secondComment = commentsFromServer[1]
	} else {
		firstComment = commentsFromServer[1]
		secondComment = commentsFromServer[0]
	}

	assert.Equal(suite.T(), uint(2), firstComment.ID)
	assert.Equal(suite.T(), uint(3), secondComment.ID)
}

func TestCommentsTestSuite(t *testing.T) {
	suite.Run(t, new(CommentsTestSuite))
}
