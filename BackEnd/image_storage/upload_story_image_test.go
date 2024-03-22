package imagestorage

import (
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/middleware"
	"bytes"
	"encoding/json"
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
	"github.com/joho/godotenv"
)

//TODO: endpoint test uploading story image,
//      implement and download story image downloading,
//      implement and download story image deletion,

func TestUploadingStoryImage(t *testing.T) {
	var b bytes.Buffer
	mw := multipart.NewWriter(&b)
	var err error

	hdr := make(textproto.MIMEHeader)

	cd := mime.FormatMediaType(
		"form-data",
		map[string]string{
			"name":     "newStoryPic",
			"filename": "test_image_sun.jpg",
		},
	)

	hdr.Set("Content-Disposition", cd)
	hdr.Set("Content-Type", "image/jpeg")

	mockRouter := gin.Default()
	mockRouter.Use(middleware.UserInfoExtractionMiddleware())
	mockRouter.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))

	reqRecorder := httptest.NewRecorder()

	file, err := os.Open("/backend/IMAGES/test_image_sun.jpg")

	if err != nil {
		t.Errorf("Could not find the test image in the filesystem")
	}
	stat, err := file.Stat()
	if err != nil {
		t.Errorf("Could not get stats for test image file")
	}

	part, err := mw.CreatePart(hdr)

	if err != nil {
		t.Errorf("Error when creating new form part : %s", err.Error())
	}

	n, err := io.Copy(part, file)

	if err != nil {
		t.Errorf(fmt.Sprintf("Error with io.Copy: %s", err.Error()))
	}

	jsonPartHeader := make(textproto.MIMEHeader)

	contentDispositionJsonPart := mime.FormatMediaType(
		"form-data",
		map[string]string{
			"name": "storyId",
		},
	)

	jsonPartHeader.Set("Content-Disposition", contentDispositionJsonPart)
	jsonPartHeader.Set("Content-Type", "application/json")

	jsonPart, err := mw.CreatePart(jsonPartHeader)

	if err != nil {
		t.Errorf("Could not create jsonpart of the multipart request to upload a story image")
	}

	picInfo := UploadedImageStoryId{
		ID: 1,
	}

	picInfoBytes, err := json.Marshal(picInfo)

	jsonPart.Write(picInfoBytes)

	mw.Close()

	if int64(n) != stat.Size() {
		t.Errorf("file size changed while writing")
	}

	godotenv.Load("../.env")
	secret, _ := helpers.GetEnv("JWT_SECRET")

	accToken, _ := middleware.GenerateJWTToken("testuser_upload", 1, []string{"ROLE_USER"}, secret, middleware.AccessToken)

	picUploadRequest, _ := http.NewRequest("POST", "/images/stories", &b)
	picUploadRequest.Body.Close()
	picUploadRequest.Header.Add("Content-Type", mw.FormDataContentType())
	picUploadRequest.Header.Set("Authorization", accToken)

	mockRouter.POST("/images/stories", UploadStoryImage)

	mockRouter.ServeHTTP(reqRecorder, picUploadRequest)

	if reqRecorder.Result().StatusCode != 202 {
		fmt.Println(reqRecorder.Body)
		t.Fatal("Request failed")
	}
}
