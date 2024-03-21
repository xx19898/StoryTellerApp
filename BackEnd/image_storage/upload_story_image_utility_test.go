package imagestorage

import (
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
)

func TestFileExtractionUtility(t *testing.T) {
	type UploadedImageStoryId struct {
		ID uint
	}

	var b bytes.Buffer

	mw := multipart.NewWriter(&b)
	var err error

	hdr := make(textproto.MIMEHeader)
	secondHeader := make(textproto.MIMEHeader)

	cd := mime.FormatMediaType(
		"form-data",
		map[string]string{
			"name":     "newStoryPic",
			"filename": "test_image_sun.jpg",
		},
	)

	contentDispositionJsonPart := mime.FormatMediaType(
		"form-data",
		map[string]string{
			"name": "storyId",
		},
	)

	hdr.Set("Content-Disposition", cd)
	hdr.Set("Content-Type", "image/jpeg")

	secondHeader.Set("Content-Disposition", contentDispositionJsonPart)
	secondHeader.Set("Content-Type", "application/json")

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
		t.Errorf("Error with io.Copy: %s", err.Error())
	}

	jsonPart, err := mw.CreatePart(secondHeader)

	if err != nil {
		t.Errorf("Error when creating json form part : %s", err.Error())
	}

	picInfo := UploadedImageStoryId{
		ID: 1,
	}

	picInfoBytes, err := json.Marshal(picInfo)

	if err != nil {
		t.Errorf("Could not marshal uploaded pic info object into json string: %s", err.Error())
	}

	jsonPart.Write(picInfoBytes)

	mw.Close()

	if int64(n) != stat.Size() {
		t.Errorf("file size changed while writing")
	}

	picUploadRequest, _ := http.NewRequest("POST", "/test", &b)
	picUploadRequest.Body.Close()
	picUploadRequest.Header.Add("Content-Type", mw.FormDataContentType())

	r := gin.Default()
	reqRecorder := httptest.NewRecorder()

	r.POST("/test", func(c *gin.Context) {
		data, err := ExtractImageFileFromStoryImageUploadRequest(c)
		//storyId, err := ExtractStoryIdFromStoryImageUploadRequest(c)

		fmt.Println("-----")
		fmt.Println(c.Request)
		fmt.Println("-----")

		fmt.Println("-----")
		fmt.Println(err)
		fmt.Println("-----")

		if err != nil {
			c.JSON(
				http.StatusInternalServerError,
				gin.H{},
			)
			return
		}

		c.JSON(
			http.StatusAccepted,
			gin.H{
				"data": string(data),
			},
		)
	})

	r.ServeHTTP(reqRecorder, picUploadRequest)

	fmt.Println(reqRecorder.Result().Status)

	if reqRecorder.Result().StatusCode != 202 {
		t.Fatal("Extracting image file from image upload request failed")
	}
}
