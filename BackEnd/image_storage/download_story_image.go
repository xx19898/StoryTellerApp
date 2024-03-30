package imagestorage

import (
	imagestoragehelpers "StoryTellerAppBackend/image_storage_helpers"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

func DownloadStoryImage(ctx *gin.Context) {
	fmt.Println("Hi guys")

	username := ctx.Query("user")
	storyId := ctx.Query("id")
	filename := ctx.Query("filename")

	var missingQueryParam string

	for _, queryParam := range []string{username, storyId, filename} {
		if queryParam == "" {
			missingQueryParam = queryParam
			break
		}
	}

	if missingQueryParam != "" {
		ctx.String(http.StatusBadRequest, fmt.Sprintf("No %s found in request queries", missingQueryParam))
	}

	fmt.Printf("username: %s, userId: %s,filename: %s", username, storyId, filename)

	storyIdAsUint, storyIdUintConvErr := strconv.ParseUint(storyId, 10, 32)

	if storyIdUintConvErr != nil {
		ctx.String(http.StatusBadRequest, "StoryId is not in correct format")
	}

	if !imagestoragehelpers.CheckIfStoryImageFileExists(username, uint(storyIdAsUint), filename) {
		ctx.String(http.StatusInternalServerError, "Could not find the sought image")
	}

	filepath := filepath.Join("backend", "IMAGES", "stories", username, storyId, filename)

	data, fileErr := os.ReadFile(filepath)

	if fileErr != nil {
		ctx.String(http.StatusInternalServerError, "Error when trying to open story image file for reading")
	}

	ctx.JSON(http.StatusOK, data)
}
