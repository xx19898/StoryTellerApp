package imagestorage

import (
	imagestoragehelpers "StoryTellerAppBackend/image_storage_helpers"
	"StoryTellerAppBackend/middleware"
	"fmt"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

//TODO: endpoint test uploading story image, DONE
//      implement and download story image downloading,
//      implement and download story image deletion,

func TestUploadingStoryImage(t *testing.T) {
	err, picUploadRequest := WriteMultipartImageRequestForTest()
	if err != nil {
		t.Fatal("Could not create image upload request")
	}

	mockRouter := gin.Default()
	mockRouter.Use(middleware.UserInfoExtractionMiddleware())
	mockRouter.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))

	reqRecorder := httptest.NewRecorder()

	mockRouter.POST("/images/stories", UploadStoryImage)

	mockRouter.ServeHTTP(reqRecorder, picUploadRequest)

	if reqRecorder.Result().StatusCode != 202 {
		fmt.Println(reqRecorder.Body)
		t.Fatal("Request failed")
	}

	testImageUserFolderDeletionErr := imagestoragehelpers.DeleteUserDirInStoriesFolder("testuser_upload")
	if testImageUserFolderDeletionErr != nil {
		t.Errorf("Error when deleting user folder where image upload test image was placed: %s", testImageUserFolderDeletionErr.Error())
	}

}
