package imagestorage

import (
	"net/http"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func UploadUserAvatar(ctx *gin.Context) {

	newAvatarFile, err := ctx.FormFile("avatarPic")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error when forming file": err.Error(),
		})
		return
	}

	currPath, err := os.Getwd()

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not get current directory",
		})
	}

	parent := filepath.Dir(currPath)
	dest := filepath.Join(parent, "IMAGES", "test", newAvatarFile.Filename)

	ctx.SaveUploadedFile(newAvatarFile, dest)

	ctx.JSON(http.StatusOK, gin.H{})
}
