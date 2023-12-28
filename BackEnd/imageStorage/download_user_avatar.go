package imagestorage

import (
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/gin-gonic/gin"
)

func DownloadUserAvatar(ctx *gin.Context) {
	username := ctx.GetString("LOGGED_USER_NAME")

	currPath, err := os.Getwd()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"error": "Internal error",
		})
	}

	parent := filepath.Dir(currPath)

	var sb strings.Builder

	sb.WriteString(username)
	sb.WriteString("Avatar")

	finalFilepathToAvatar := filepath.Join(parent, "IMAGES", sb.String())

	avatarFile, err := os.ReadFile(finalFilepathToAvatar)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find avatar for chosen user",
		})
	}

	ctx.Data(
		http.StatusOK,
		"image/x-cion",
		avatarFile,
	)
}
