package imagestorage

import (
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ImageInfo struct {
	StoryId    uint   `form:"id"`
	Identifier string `form:"identifier"`
}

func uploadImage(ctx *gin.Context) {
	var imageInfo ImageInfo

	newStoryPic, err := ctx.FormFile("newStoryPic")
	currPath, err := os.Getwd()

	parent := filepath.Dir(currPath)
	folderWithSameUsername := filepath.Join(parent, "IMAGES", ctx.GetString("LOGGED_USER_NAME"))
	story, err := databaselayer.FindStoryById(imageInfo.StoryId)
	storyIdString := strconv.FormatUint(uint64(imageInfo.StoryId), 10)

	if ctx.ShouldBind(&imageInfo) != nil {
		ctx.String(400, "Unsuccessful")
	}

	// 1) Form new imagefile name from storyid and identifier
	imageName := story.Title + "-" + imageInfo.Identifier
	dest := filepath.Join(folderWithSameUsername, storyIdString, imageName)

	// 2) Check if folder with same username exists already, if not create one

	if err != nil {
		ctx.String(http.StatusInternalServerError, "Could not get current directory")
	}

	result := CheckThatDirectoryExists(folderWithSameUsername)

	if !result {
		err := os.Chdir("../IMAGES/")

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not get to IMAGES folder with chdir",
			})
		}
		os.Mkdir(ctx.GetString("LOGGED_USER_NAME"), os.ModePerm)
		os.Chdir(ctx.GetString("LOGGED_USER_NAME"))
		os.Mkdir(storyIdString, os.ModePerm)
		ctx.SaveUploadedFile(newStoryPic, dest)
		ctx.JSON(http.StatusOK, gin.H{})
	}

	// IMAGES/storyWithSameUsername exists
	folderWithSameStoryId := filepath.Join(folderWithSameUsername, strconv.FormatUint(uint64(imageInfo.StoryId), 10))
	result = CheckThatDirectoryExists(folderWithSameStoryId)

	if !result {
		err := os.Chdir("../IMAGES/")
		secErr := os.Chdir("/" + storyIdString)

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Could not get to IMAGES folder with chdir",
			})
		}

		if secErr != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": fmt.Sprintf("Could not get to IMAGES folder with chdir"),
			})
		}
		ctx.JSON(http.StatusOK, gin.H{})
	}

	ctx.SaveUploadedFile(newStoryPic, dest)
	ctx.JSON(http.StatusOK, gin.H{})

	// !!! check that directory exists,
	// if not create, then check if directory with same story id exists,
	// if not create, then check if file with same name exists,
	// if exists update, if not create

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Could not get current directory",
		})
	}

	// 3) In the folder with same username as the one that sent the request save
	// 		   the image as a file with the same name a in first part of instructions
}
