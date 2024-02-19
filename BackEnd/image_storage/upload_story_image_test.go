package imagestorage

import (
	"bytes"
	"fmt"
	"io"
	"mime"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"net/textproto"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUploadingStoryImage(t *testing.T) {
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	var err error

	hdr := make(textproto.MIMEHeader)

	cd := mime.FormatMediaType(
		"form-data",
		map[string]string{
			"name":     "storyPic",
			"filename": "test_image_sun.jpg",
		},
	)

	hdr.Set("Content-Disposition", cd)
	hdr.Set("Content-Type", "image/jpeg")

	mockRouter := gin.Default()
	reqRecorder := httptest.NewRecorder()

	file, err := os.Open("/backend/IMAGES/test_image.sun.jpg")
	if err != nil {
		t.Errorf("Could not find the test image in the filesystem")
	}
	stat, err := file.Stat()
	if err != nil {
		t.Errorf("Could not get stats for test image file")
	}

	part, err := mw.CreatePart(hdr)
	if err != nil {
		t.Error(fmt.Sprintf("Error when creating new form part : %s", err.Error()))
	}

	n, err := io.Copy(part, file)

	if err != nil {
		t.Errorf(fmt.Sprintf("Error with io.Copy: %s", err.Error()))
	}

	mw.Close()

	if int64(n) != stat.Size() {
		t.Errorf("file size changed while writing")
	}

	picUploadRequest, _ := http.NewRequest("POST", "/images/stories", &b)
	picUploadRequest.Body.Close()
	picUploadRequest.Header.Add("Content-Type", mw.FormDataContentType())
	mockRouter.POST("/images/stories", UploadStoryImage)
}
