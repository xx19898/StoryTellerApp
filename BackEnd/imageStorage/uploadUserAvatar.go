package imagestorage

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UploadUserAvatar(ctx *gin.Context) {

	newAvatarFileHeader, err := ctx.FormFile("avatarPic")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	newAvatarFile, err := newAvatarFileHeader.Open()

	fmt.Println("-------")
	fmt.Println(newAvatarFile)
	fmt.Println("-------")

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Error while reading the file associated with the newAvatar key in the request",
		})
		return
	}

	//Test if newAvatar is infact image
	ctx.JSON(http.StatusOK, gin.H{})

}
