package helpers

import (
	"errors"

	"github.com/golang-jwt/jwt"
)

type CustomClaims struct {
	roles []string `json:roles`
	id    int64    `json:id`
	jwt.StandardClaims
}

func CreateToken(username string, id int64, roles []string, expirationDate int64, issuedAt int64) (string, error) {
	secretKey, err := GetEnv("JWT_SECRET")

	if !err {
		return "", errors.New("could not load jwtSecret from env variables")
	}

	claims := CustomClaims{
		[]string{"User"},
		id,
		jwt.StandardClaims{
			ExpiresAt: expirationDate,
			IssuedAt:  issuedAt,
			Issuer:    username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	ss, error := token.SignedString([]byte(secretKey))
	if error != nil {
		return "", error
	}
	return ss, nil
}

func ParseJWTToken(token string) (*CustomClaims, error) {
	secretKey, ok := GetEnv("JWT_SECRET")
	if !ok {
		return &CustomClaims{}, errors.New("could not attain jwt secret from env variables")
	}

	parsedToken, error := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if error != nil {
		return &CustomClaims{}, error
	}

	if claims, ok := parsedToken.Claims.(*CustomClaims); ok && parsedToken.Valid {
		return claims, nil
	} else {
		return &CustomClaims{}, errors.New("could not extract claims from the jwt token and make sure it is correct")
	}
}
