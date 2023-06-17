package auth

import (
	"StoryTellerAppBackend/configuration"
	databaselayer "StoryTellerAppBackend/databaseLayer"
	"StoryTellerAppBackend/models"
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type RegisterTestSuite struct {
	suite.Suite
}

func (suite *RegisterTestSuite) SetupSuite() {
	configuration.ConfigureDatabaseForTest()
}

func (suite *RegisterTestSuite) SetupTest() {
	fmt.Println("Before each test")
	configuration.ResetEverythingElseExceptRoles()
}

func (suite *RegisterTestSuite) TestingBasicCreationOfUserEntityInnD() {
	newRole := models.Role{Name: "USER"}
	_ = databaselayer.CreateNewUser("NewUser", "password", "xx", []models.Role{newRole})
	//newUser2 := models.User{Name: "NewUser", Password: "password", Email: "xx", Roles: []models.Role{newRole}, ID: 0}
	//result := configuration.DB.Create(&newUser2)
	userFromDataBase := databaselayer.FindUserByName("NewUser")
	assert.Equal(suite.T(), "NewUser", userFromDataBase.Name)
	fmt.Println("Username: " + userFromDataBase.Name + " id: " + strconv.FormatUint(uint64(userFromDataBase.ID), 10))
}

func (suite *RegisterTestSuite) TestingThatEndPointWorks() {
	router := gin.Default()
	router.POST("/register", Register)
	user := models.User{Name: "NewUser", Password: "PasswordZ", Email: "Email"}
	jsonValue, _ := json.Marshal(user)
	req, _ := http.NewRequest("POST", "/register", bytes.NewBuffer(jsonValue))
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)
	assert.Equal(suite.T(), http.StatusCreated, w.Code)
	freshlyCreatedUser := databaselayer.FindUserByName("NewUser")
	assert.Equal(suite.T(), freshlyCreatedUser.Password, "PasswordZ")
	fmt.Println("freshlyCreatedUserId: " + strconv.FormatUint(uint64(freshlyCreatedUser.ID), 10))
}

func TestRegisterTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}
