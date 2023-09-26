package middleware

import (
	"fmt"
	"testing"
)

var testSecret = "my-32-character-ultra-secure-and-ultra-long-secret"

func TestGenerateJWTToken(t *testing.T) {
	fmt.Println("STARTING!")
	testedToken, err := GenerateJWTToken("testUser", 1, []string{"ROLE_USER"}, testSecret, AccessToken)
	if err != nil {
		t.Fatal("Could not generate the token", err.Error())
	}
	fmt.Println("GENERATED THE TOKEN")
	testedTokenIsValid := ValidateJWTToken(testedToken, testSecret)
	fmt.Println("VALIDATED THE TOKEN")
	if !testedTokenIsValid {
		t.Fatal("Returned token is not valid")
	}
}

func TestDecidingValidTimeInHours(t *testing.T) {
	time, err := decideTimeValidInHours(RefreshToken)
	if err != nil {
		t.Fatal("Unknown type of token given as parameter")
	}
	if time != 24 {
		t.Fatal("Incorrect time returned for the refreshToken")
	}
}

func TestEncodingAndDecodingClaims(t *testing.T) {
	testedToken, errTokenGen := GenerateJWTToken("testUser", 1, []string{"ROLE_USER", "ROLE_ADMIN"}, testSecret, AccessToken)
	username, roles, errRolesExtract := ExtractRolesAndUsername(testedToken, testSecret)
	if errRolesExtract != nil {
		t.Fatal(errRolesExtract.Error())
	}
	if errTokenGen != nil {
		t.Fatal("Could not generate JWT token")
	}
	if username != "testUser" {
		t.Fatal("Wrong username")
	}
	if roles[0] != "ROLE_USER" {
		fmt.Println(roles[0])
		t.Fatal("Wrong role")
	}
}
