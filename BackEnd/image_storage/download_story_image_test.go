package imagestorage

import (
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/middleware"
	"bytes"
	"encoding/json"
	"errors"
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

func WriteMultipartImageRequestForTest() (error, *http.Request) {
	var b bytes.Buffer
	var picUploadRequest *http.Request

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
/*
	mockRouter := gin.Default()
	mockRouter.Use(middleware.UserInfoExtractionMiddleware())
	mockRouter.Use(middleware.AuthorizationMiddleware(middleware.CompareRoles, []string{"ROLE_USER"}))
*/
	file, err := os.Open("/backend/IMAGES/test_image_sun.jpg")
	if err != nil {
		return fmt.Errorf("could not find the test image in the filesystem"), picUploadRequest
	}

	stat, err := file.Stat()
	if err != nil {
		return fmt.Errorf("could not get stats for test image file"), picUploadRequest
	}

	part, err := mw.CreatePart(hdr)
	if err != nil {
		return fmt.Errorf("error when creating new form part : %s", err.Error()), picUploadRequest
	}

	n, err := io.Copy(part, file)

	if err != nil {
		return fmt.Errorf("error with io.Copy: %s", err.Error()), picUploadRequest
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
		return fmt.Errorf("сould not create jsonpart of the multipart request to upload a story image"), picUploadRequest
	}

	picInfo := UploadedImageStoryId{
		ID: 1,
	}

	picInfoBytes, err := json.Marshal(picInfo)
	if err != nil {
		return errors.New("сould not marshal the picInfo Object"), picUploadRequest
	}

	jsonPart.Write(picInfoBytes)

	mw.Close()

	if int64(n) != stat.Size() {
		return errors.New("file size changed while writing"), picUploadRequest
	}

	godotenv.Load("../.env")
	secret, _ := helpers.GetEnv("JWT_SECRET")

	accToken, _ := middleware.GenerateJWTToken("testuser_upload", 1, []string{"ROLE_USER"}, secret, middleware.AccessToken)

	picUploadRequest, _ = http.NewRequest("POST", "/images/stories", &b)
	picUploadRequest.Body.Close()
	picUploadRequest.Header.Add("Content-Type", mw.FormDataContentType())
	picUploadRequest.Header.Set("Authorization", accToken)

	return err, picUploadRequest
}

func TestDownloadingStoryImage(t *testing.T) {
	var emptyBuffer bytes.Buffer
	downloadReqRecorder := httptest.NewRecorder()

	mockRouter := gin.Default()
	downloadRequest, _ := http.NewRequest("GET", "/static/stories/testUser/1/test_image_1.jpg", &emptyBuffer)

	mockRouter.Static("/static", "../IMAGES")

	mockRouter.ServeHTTP(downloadReqRecorder, downloadRequest)

	if downloadReqRecorder.Result().StatusCode != 200 {
		t.Fatal("Wrong status code")
	}

	var body []byte
	downloadRequest.Body.Read(body)

	fmt.Println(body)
}
