package imagestorage

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

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

	dst := filepath.Join(parentDir, "IMAGES", "testUser_avatar")
	destination, err := os.Create(dst)
	assert.Nil(suite.T(), err)

	defer destination.Close()
	_, err = io.Copy(image, destination)

}

func (suite *DownloadUserAvatarTestSuite) TearDownTest() {
	currDir, _ := os.Getwd()
	parentDir := filepath.Dir(currDir)

	imageToDeletePath := filepath.Join(parentDir, "IMAGES", "testUser_avatar.jpg")
	os.Remove(imageToDeletePath)
}

func (suite *DownloadUserAvatarTestSuite) TestDownloadingAvatarEndpoint() {
	fmt.Println("TESTING!")
	number := 25
	assert.Equal(suite.T(), number, 25)
}

func TestCommentsTestSuite(t *testing.T) {
	suite.Run(t, new(DownloadUserAvatarTestSuite))
}
