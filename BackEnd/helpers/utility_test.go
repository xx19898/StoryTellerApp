package helpers

import (
	"StoryTellerAppBackend/models"
	"testing"
	"time"

	"github.com/joho/godotenv"
)

func TestEnvVariablesGetLoaded(t *testing.T) {
	godotenv.Load("../.env")

	_, ok := GetEnv("JWT_SECRET")

	if !ok {
		t.Fatal("Could not find the jwt_secret environment variable")
	}
}

func TestThatPasswordGetsEncryptedAndCanThenBeComparedWithAString(t *testing.T) {
	testPassword := "testPassword2023"
	var encryptedTestPassword, err = EncryptPassword(testPassword)
	if err != nil {
		t.Fatal(err.Error())
	}
	passwordsMatch := PasswordMatchesTheHash(testPassword, string(encryptedTestPassword))
	if !passwordsMatch {
		t.Fatal("Given password and hash of the encrypted password do not match")
	}
}

func TestThatJWTTokenGetsCreated(t *testing.T) {
	godotenv.Load("../.env")
	token, err := CreateToken("testUser", 1, []string{"User"}, time.Now().Add(time.Hour*time.Duration(2)).Unix(), time.Now().Unix())

	if err != nil {
		t.Fatal("Occured error in the process of password creation " + err.Error())
	}

	claims, errr := ParseJWTToken(token)

	if errr != nil {
		t.Fatal(errr.Error())
	}

	username := claims.Issuer
	if username != "testUser" {
		t.Fatal("Incorrect username parsed")
	}
}

func TestThatPasswordReturnsNegativeWhenPasswordsDontMatch(t *testing.T) {
	testPassword := "testPassword2023"
	var encryptedTestPassword, err = EncryptPassword(testPassword)
	if err != nil {
		t.Fatal(err.Error())
	}
	passwordsMatch := PasswordMatchesTheHash("wrongPassword", string(encryptedTestPassword))
	if passwordsMatch {
		t.Fatal("Password should not match the hash")
	}
}

func TestThatNamesGetExtractedFromRoleObjects(t *testing.T) {
	userRole := models.Role{Name: "USER", ID: 0}
	adminRole := models.Role{Name: "ADMIN", ID: 1}

	roles := []models.Role{userRole, adminRole}
	roleNames := RolesToString(roles)

	roleNameMap := make(map[string]bool)

	if len(roleNames) != 2 {
		t.Fatal("Incorrect number of role names, should be 3 and returned ", len(roleNames))
	}

	for i := 0; i < len(roleNames); i++ {
		roleNameMap[roleNames[i]] = true
	}

	if roleNameMap["USER"] != true || roleNameMap["ADMIN"] != true {
		t.Fatal("Incorrect role names returned")
	}
}
