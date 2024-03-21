package imagestorage

import (
	"io"

	"github.com/gin-gonic/gin"
)

/*
func ExtractStoryIdFromStoryImageUploadRequest(ctx *gin.Context) ([]byte,error){
	var req = ctx.Request
}
*/

func ExtractImageFileFromStoryImageUploadRequest(ctx *gin.Context) ([]byte, error) {
	var data []byte
	var err error

	newStoryPicFileHeader, err := ctx.FormFile("newStoryPic")

	if err != nil {
		return data, err
	}

	newStoryPicFile, err := newStoryPicFileHeader.Open()

	if err != nil {
		return data, err
	}

	data, err = io.ReadAll(newStoryPicFile)

	if err != nil {
		return data, err
	}

	return data, err
}
