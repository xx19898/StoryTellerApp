package middleware

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func TestThatNotAuthorizedIfNoAuthTokenAttached(t *testing.T) {
	godotenv.Load("../.test.env")
	router := gin.Default()
	router.Use(AuthorizationMiddleware())
	router.POST("/test", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	mockRequest, _ := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("MOCK")))
	router.ServeHTTP(w, mockRequest)

	if w.Code != 403 {
		t.Error("Http code is wrong. Should be 403, as no auth token is attached to the request")
	}

}
