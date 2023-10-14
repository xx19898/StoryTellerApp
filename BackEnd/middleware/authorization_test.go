package middleware

import (
	"StoryTellerAppBackend/helpers"
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
	router.Use(UserInfoExtractionMiddleware())
	router.POST("/test", func(c *gin.Context) {})

	w := httptest.NewRecorder()
	mockRequest, _ := http.NewRequest("POST", "/test", bytes.NewBuffer([]byte("MOCK")))
	router.ServeHTTP(w, mockRequest)

	if w.Code != 403 {
		t.Error("Http code is wrong. Should be 403, as no auth token is attached to the request")
	}

}

func TestThatComparingRolesWorks(t *testing.T) {
	foundRoles := []string{"ROLE_USER", "ROLE_ADMIN"}
	neededRoles := []string{"ROLE_ADMIN"}

	firstResult := CompareRoles(neededRoles, foundRoles)
	if firstResult {
		t.Error("First comparison is not right, should be false")
	}

	neededRoles2 := []string{"ROLE_USER", "ROLE_ADMIN"}
	foundRoles2 := []string{"ROLE_ADMIN", "ROLE_USER"}

	secondResult := CompareRoles(neededRoles2, foundRoles2)
	if !secondResult {
		t.Error("Second comparison is not right, should be true")
	}
}

func TestThatAuthorizationPipelineWorks(t *testing.T) {
	neededRoles := []string{"ROLE_USER"}

	//create token

	godotenv.Load("../.test.env")
	secret, jwtSecretOk := helpers.GetEnv("JWT_SECRET")
	testedToken, err := GenerateJWTToken("testUser", 1, []string{"ROLE_USER"}, secret, AccessToken)

	if !jwtSecretOk {
		t.Fatal("Could not extract jwt secret from env vars")
	}
	if err != nil {
		t.Fatal("Failed when trying to create jwt token")
	}

	router := gin.Default()
	router.Use(UserInfoExtractionMiddleware())
	router.Use(AuthorizationMiddleware(CompareRoles, neededRoles))
	router.POST("/testAuth", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Everything OK",
		})
	})

	w := httptest.NewRecorder()
	mockRequest, _ := http.NewRequest("POST", "/testAuth", bytes.NewBuffer([]byte("MOCK")))
	mockRequest.Header.Set("Authorization", testedToken)
	router.ServeHTTP(w, mockRequest)

	if w.Code != 200 {
		t.Error("Http code is wrong. Should be 200")
	}
}
