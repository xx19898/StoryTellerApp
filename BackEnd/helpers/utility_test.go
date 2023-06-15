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
