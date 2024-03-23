package imagestorage

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DownloadStoryImage(ctx *gin.Context) {
	fmt.Println("Hi guys")

	ctx.String(http.StatusOK, "")

}
