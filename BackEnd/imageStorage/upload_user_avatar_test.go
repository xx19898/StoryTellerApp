package imagestorage

import (
	"bytes"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUploadUserAvatar(t *testing.T) {
	var b bytes.Buffer
	var err error

	mockRouter := gin.Default()
	reqRecorder := httptest.NewRecorder()

	file, err := os.Open("/backend/testAssets/test_imageLA.jpg")

	if err != nil {
		t.Errorf("Could not find the test file in the filesystem")
	}

	w := multipart.NewWriter(&b)

	var fw io.Writer

	if fw, err = w.CreateFormFile("avatarPic", "avatarPic"); err != nil {
		t.Errorf("Error creating writer: %v", err)
	}

	if _, err = io.Copy(fw, file); err != nil {
		t.Errorf("Error with io.Copy: %v", err)
	}
	w.Close()

	picUploadRequest, _ := http.NewRequest("POST", "/setUserAvatar", &b)
	picUploadRequest.Header.Add("Content-Type", "multipart/form-data")
	picUploadRequest.Header.Add("boundary", w.Boundary())
	mockRouter.POST("/setUserAvatar", UploadUserAvatar)

	mockRouter.ServeHTTP(reqRecorder, picUploadRequest)

	if reqRecorder.Code != 200 {
		fmt.Println(reqRecorder.Body)
		panic("Code : " + strconv.Itoa(reqRecorder.Code))
	}
}
