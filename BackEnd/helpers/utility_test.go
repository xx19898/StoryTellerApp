package helpers

import (
	"fmt"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnvVariablesGetLoaded(t *testing.T) {
	godotenv.Load("../.env")
	secret, ok := GetEnv("JWT_SECRET")
	if !ok {
		t.Fatal("Could not find the jwt_secret environment variable")
	}
	fmt.Println("Hello" + secret)
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
