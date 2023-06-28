package helpers

import (
	"fmt"
	"strconv"
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
	fmt.Println("passwords should match: " + strconv.FormatBool(passwordsMatch))
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
	date := time.Unix(claims.ExpiresAt, 0)
	fmt.Println("expiration date: ", date)

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
