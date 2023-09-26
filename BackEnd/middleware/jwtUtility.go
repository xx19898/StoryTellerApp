package middleware

import (
	"errors"
	"fmt"
	"log"
	"strconv"
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
	jwt.RegisteredClaims
}

func GenerateJWTToken(Username string, Id int, Roles []string, Secret string, TokenType TokenType) (string, error) {
	timeValidInHours, err := decideTimeValidInHours(TokenType)
	fmt.Println("Time token is proper: " + strconv.Itoa(timeValidInHours))
	fmt.Println("Time now : " + jwt.NewNumericDate(time.Now()).String())
	fmt.Println("Time token expires: " + jwt.NewNumericDate(time.Now().Add(time.Hour*time.Duration(timeValidInHours))).String())
	if err != nil {
		log.Fatal(err)
	}
	customClaims := CustomClaims{
		Roles,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour + time.Duration(timeValidInHours))),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    Username,
			Subject:   "User",
			ID:        strconv.Itoa(Id),
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
	parsedToken, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretAsBytes := []byte(secret)
		return secretAsBytes, nil
	})
	claims := parsedToken.Claims
	fmt.Println(claims.GetExpirationTime())

	return err == nil
}

func ExtractRolesAndUsername(token string, secret string) (string, []string, error) {
	parsedToken, _ := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		secretAsBytes := []byte(secret)
		return secretAsBytes, nil
	})
	claims := parsedToken.Claims.(jwt.MapClaims)
	username, _ := claims.GetIssuer()
	fmt.Println(claims)
	roles, exists := claims["Roles"]
	if !exists {
		return "", []string{}, errors.New("roles not found in claims")
	}
	rolesAsInterfaceArray, ok := roles.([]interface{})
	if !ok {
		return "", []string{}, errors.New("Error when parsing array of roles in custom claims of jwt token")
	}
	var rolesAsStringArray []string
	for i := 0; i < len(rolesAsInterfaceArray); i++ {
		roleString, ok := rolesAsInterfaceArray[i].(string)
		if !ok {
			return "", []string{}, errors.New(fmt.Sprintf("Error when casting role under index %s to string", fmt.Sprintf("%d", i)))

		}
		rolesAsStringArray = append(rolesAsStringArray, roleString)
	}
	fmt.Println(rolesAsStringArray)
	return username, rolesAsStringArray, nil
}
