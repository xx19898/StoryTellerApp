package helpers

import (
	"fmt"
	"strconv"
	"testing"

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
	fmt.Println("passwords should match: " + strconv.FormatBool(passwordsMatch))
	if !passwordsMatch {
		t.Fatal("Given password and hash of the encrypted password do not match")
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
