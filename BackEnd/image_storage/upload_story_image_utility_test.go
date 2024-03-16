package imagestorage

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestFileExtractionUtility(t *testing.T) {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

}
