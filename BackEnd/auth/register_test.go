package auth

import (
	"StoryTellerAppBackend/configuration"
	"StoryTellerAppBackend/models"
	"fmt"
	"strconv"
	"testing"

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

func (suite *RegisterTestSuite) TestingBasicInjection() {
	newRole := models.Role{Name: "USER"}
	newUser := models.User{Name: "NewUser", Password: "password", Email: "xx", Roles: []models.Role{newRole}}

	result := configuration.DB.Create(&newUser)
	fmt.Println("Rows affected: " + strconv.Itoa(int(result.RowsAffected)))
	if result.Error != nil {
		panic("Something went wrong when inserting a new user")
	}
	var userFromDataBase models.User
	configuration.DB.Where(models.User{Name: "NewUser"}).First(&userFromDataBase)
	fmt.Println("Username: " + userFromDataBase.Name)
}

func TestRegisterTestSuite(t *testing.T) {
	suite.Run(t, new(RegisterTestSuite))
}
