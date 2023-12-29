package imagestorage

import (
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/middleware"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"path/filepath"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type DownloadUserAvatarTestSuite struct {
	suite.Suite
}

func (suite *DownloadUserAvatarTestSuite) SetupSuite() {

	currDir, err := os.Getwd()
	parentDir := filepath.Dir(currDir)

	fmt.Println("*************")
	imagePath := filepath.Join(parentDir, "testAssets", "test_imageLA.jpg")
	image, err := os.Open(imagePath)
	assert.Nil(suite.T(), err)

	defer image.Close()

	dst := filepath.Join(parentDir, "IMAGES", "testUser_avatar.jpg")
	destination, err := os.Create(dst)
	assert.Nil(suite.T(), err)

	defer destination.Close()
	_, err = io.Copy(destination, image)

}

func (suite *DownloadUserAvatarTestSuite) TearDownTest() {
	currDir, _ := os.Getwd()
	parentDir := filepath.Dir(currDir)

	imageToDeletePath := filepath.Join(parentDir, "IMAGES", "testUser_avatar")
	os.Remove(imageToDeletePath)
}

func (suite *DownloadUserAvatarTestSuite) TestDownloadingAvatarEndpoint() {

	mockRouter := gin.Default()
	secret, _ := helpers.GetEnv("JWT_SECRET")
	accToken, _ := middleware.GenerateJWTToken("testuser", 1, []string{"ROLE_USER"}, secret, middleware.AccessToken)

	recorder := httptest.NewRecorder()
	mockRouter.GET("/getUserAvatar", DownloadUserAvatar)

	avatarDownloadingRequest, _ := http.NewRequest("GET", "/getUserAvatar", bytes.NewBuffer([]byte{}))
	avatarDownloadingRequest.Header.Set("Authorization", accToken)

	mockRouter.ServeHTTP(recorder, avatarDownloadingRequest)

	receivedFile := recorder.Body.Bytes()

	fmt.Println("-----------")
	fmt.Println(receivedFile)
	fmt.Println("-----------")

	/*

		parentDir := filepath.Dir(currDir)
		imagePath := filepath.Join(parentDir, "testAssets", "test_imageLA.jpg")
		image, err := os.Open(imagePath)
		assert.Nil(suite.T(), err)


		result	 := os.SameFile(fil1,fil2)
		assert.True(suite.T(),result)

	*/
}

func TestCommentsTestSuite(t *testing.T) {
	suite.Run(t, new(DownloadUserAvatarTestSuite))
}
