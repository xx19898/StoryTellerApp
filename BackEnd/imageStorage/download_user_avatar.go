package imagestorage

import (
	"fmt"
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
			"error": "Internal server error",
		})
		return
	}

	parent := filepath.Dir(currPath)

	var sb strings.Builder

	fmt.Println("***********")
	fmt.Println(username)
	fmt.Println("***********")

	sb.WriteString(username)
	sb.WriteString("_avatar.jpg")

	finalFilepathToAvatar := filepath.Join(parent, "IMAGES", sb.String())

	fmt.Println("?????????")
	fmt.Println(finalFilepathToAvatar)
	fmt.Println("?????????")

	avatarFile, err := os.ReadFile(finalFilepathToAvatar)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Could not find avatar for chosen user",
		})
		return
	}

	ctx.Data(
		http.StatusOK,
		"image/x-cion",
		avatarFile,
	)
}
