package auth

import (
	"StoryTellerAppBackend/configuration"
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/helpers"
	"StoryTellerAppBackend/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type AuthTestSuite struct {
	suite.Suite
}

func (suite *AuthTestSuite) SetupSuite() {
	godotenv.Load("../.env")
	configuration.ConfigureDatabaseForTest()
}

func (suite *AuthTestSuite) SetupTest() {
	configuration.ResetEverythingElseExceptRoles()
}

func (suite *AuthTestSuite) TestingBasicCreationOfUserEntityInDB() {
	newRole := models.Role{Name: "USER"}
	_ = databaselayer.CreateNewUser("NewUser", "password", "xx", []models.Role{newRole})
	//newUser2 := models.User{Name: "NewUser", Password: "password", Email: "xx", Roles: []models.Role{newRole}, ID: 0}
	//result := configuration.DB.Create(&newUser2)
	userFromDataBase := databaselayer.FindUserByName("NewUser")
	assert.Equal(suite.T(), "NewUser", userFromDataBase.Name)
}

func (suite *AuthTestSuite) TestingThatRegisterEndPointWorks() {
	router := gin.Default()
	router.POST("/register", Register)
	user := models.User{Name: "NewUser", Password: "PasswordZ", Email: "Email"}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	freshlyCreatedUser := databaselayer.FindUserByName("NewUser")
	assert.True(suite.T(), helpers.PasswordMatchesTheHash("PasswordZ", freshlyCreatedUser.Password))

}

func (suite *AuthTestSuite) TestThatLoginWithRightCredsWorks() {
	router := gin.Default()
	router.POST("/register", Register)
	user := models.User{Name: "testUser", Password: "testPassword", Email: "Email"}
	jsonValue, _ := json.Marshal(user)
	registerReq, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, registerReq)

	router.POST("/login", Login)
	userForm := models.User{Name: "testUser", Password: "testPassword"}
	userFormAsJson, _ := json.Marshal(userForm)
	loginReq, _ := http.NewRequest("POST", "/login", bytes.NewBuffer(userFormAsJson))
	v := httptest.NewRecorder()
	router.ServeHTTP(v, loginReq)
	assert.Equal(suite.T(), http.StatusAccepted, v.Code)
}

func TestAuthTestSuite(t *testing.T) {
	suite.Run(t, new(AuthTestSuite))
}
