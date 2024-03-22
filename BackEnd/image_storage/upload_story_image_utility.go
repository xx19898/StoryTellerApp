package imagestorage

import (
	"encoding/json"
	"errors"
	"io"

	"github.com/gin-gonic/gin"
)

type UploadedImageStoryId struct {
	ID uint
}

func ExtractStoryIdFromStoryImageUploadRequest(ctx *gin.Context) (uint, error) {
	var storyId uint

	storyIdStringValue := ctx.Request.FormValue("storyId")

	if storyIdStringValue == "" {
		return storyId, errors.New("found no form value for storyId in the request")
	}

	var storyIdJson UploadedImageStoryId

	err := json.Unmarshal([]byte(storyIdStringValue), &storyIdJson)

	if err != nil {
		return storyId, errors.New("could not process the storyId key-value pair in the story upload request")
	}

	return storyIdJson.ID, err
}

func ExtractImageFileFromStoryImageUploadRequest(ctx *gin.Context) ([]byte, string, error) {
	var data []byte
	var err error

	newStoryPicFileHeader, err := ctx.FormFile("newStoryPic")

	if err != nil {
		return data, "", err
	}

	filename := newStoryPicFileHeader.Filename

	if len(filename) == 0 {
		return data, "", errors.New("length of the filename of the story pic being uploaded is 0")
	}

	newStoryPicFile, err := newStoryPicFileHeader.Open()

	if err != nil {
		return data, "", err
	}

	data, err = io.ReadAll(newStoryPicFile)

	if err != nil {
		return data, "", err
	}

	return data, filename, err
}
