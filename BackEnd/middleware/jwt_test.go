package middleware

import "testing"

var testSecret = "TestSecret"

func TestGenerateJWTToken(t *testing.T) {
	testedToken := GenerateJWTToken("testUser", []string{"ROLE_USER"}, testSecret)
	testedTokenIsValid := ValidateJWTToken(testedToken, testSecret)
	if !testedTokenIsValid {
		t.Fatal("Returned token is not valid")
	}
}

func TestDecidingValidTimeInHours(t *testing.T) {
	time := dec
}
