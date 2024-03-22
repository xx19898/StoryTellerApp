package imagestorage

import (
	imagestoragehelpers "StoryTellerAppBackend/image_storage_helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadStoryImage(ctx *gin.Context) {

	userBehindRequest := ctx.GetString("LOGGED_USER_NAME")

	data, picFilename, storyPicFileExtractionErr := ExtractImageFileFromStoryImageUploadRequest(ctx)
	storyId, storyIdExtractionErr := ExtractStoryIdFromStoryImageUploadRequest(ctx)

	if storyPicFileExtractionErr != nil {
		ctx.String(http.StatusInternalServerError, "Could not extract the story picture from the request")
		return
	}

	if storyIdExtractionErr != nil {
		ctx.String(http.StatusInternalServerError, "Could not extract the story id from the request")
		return
	}

	err := imagestoragehelpers.CreateNewStoryImage(userBehindRequest, storyId, data, picFilename)

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Could not save the uploaded picture file to the filesystem on the server")
		return
	}

	ctx.String(http.StatusAccepted, "Story uploaded correctly")
}
