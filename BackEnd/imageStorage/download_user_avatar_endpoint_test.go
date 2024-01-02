package imagestorage

import (
	"bytes"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestDownloadingUserAvatar(t *testing.T) {
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	var err error

	mockRouter := gin.Default()
	reqRecorder := httptest.NewRecorder()

	picDownloadRequest, _ := http.NewRequest(("GET"), "getUserAvatar", &b)
	picDownloadRequest.Body.Close()

	mockRouter.POST("/getUserAvatar")
}
