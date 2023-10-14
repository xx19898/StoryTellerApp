package middleware

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType struct {
	typeName string
}

var (
	RefreshToken = TokenType{"RefreshToken"}
	AccessToken  = TokenType{"AccessToken"}
)

type CustomClaims struct {
	Roles []string
	ID    int
	jwt.RegisteredClaims
}

func GenerateJWTToken(Username string, Id int, Roles []string, Secret string, TokenType TokenType) (string, error) {
	timeValidInHours, err := decideTimeValidInHours(TokenType)
	if err != nil {
		log.Fatal(err)
	}
	customClaims := CustomClaims{
		Roles,
		Id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour + time.Duration(timeValidInHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    Username,
			Subject:   "User",
			Audience:  []string{"StoryTellerApp"},
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	secretAsBytes := []byte(Secret)
	signedString, err := token.SignedString(secretAsBytes)
	return signedString, err
}

func decideTimeValidInHours(tokenType TokenType) (int, error) {
	if tokenType == RefreshToken {
		return 24, nil
	}
	if tokenType == AccessToken {
		return 2, nil
	}
	return 0, errors.New("unknown type of token: " + tokenType.typeName)
}

func ValidateJWTToken(token string, secret string) bool {
	_, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretAsBytes := []byte(secret)
		return secretAsBytes, nil
	})

	return err == nil
}

func ExtractUserInfo(token string, secret string) (string, uint, []string, error) {
	parsedToken, _ := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretAsBytes := []byte(secret)
		return secretAsBytes, nil
	})
	claims := parsedToken.Claims.(jwt.MapClaims)

	username, usernameNotFoundError := claims.GetIssuer()

	if usernameNotFoundError != nil {
		return "", 0, []string{}, errors.New("issuer not found in claims")
	}

	id, idNotFoundError := claims["ID"]

	if !idNotFoundError {
		return "", 0, []string{}, errors.New("id not found in claims")
	}

	idAsFloat, idTransformOk := id.(float64)

	if !idTransformOk {
		return "", 0, []string{}, errors.New("id is in improper format when parsing string to int")
	}

	idAsInt := int(idAsFloat)

	fmt.Println("Id:")
	fmt.Println(idAsInt)

	roles, exists := claims["Roles"]
	if !exists {
		return "", 0, []string{}, errors.New("roles not found in claims")
	}
	rolesAsInterfaceArray, ok := roles.([]interface{})
	if !ok {
		return "", 0, []string{}, errors.New("Error when parsing array of roles in custom claims of jwt token")
	}
	var rolesAsStringArray []string
	for i := 0; i < len(rolesAsInterfaceArray); i++ {
		roleString, ok := rolesAsInterfaceArray[i].(string)
		if !ok {
			return "", 0, []string{}, errors.New(fmt.Sprintf("Error when casting role under index %s to string", fmt.Sprintf("%d", i)))
		}
		rolesAsStringArray = append(rolesAsStringArray, roleString)
	}
	return username, uint(idAsInt), rolesAsStringArray, nil
}
