package middleware

import "github.com/golang-jwt/jwt/v5"

type TokenType struct{
	typeName string
}

var(
	RefreshToken = TokenType{"RefreshToken"}
	AccessToken = TokenType{"AccessToken"}
)

type CustomClaims struct{
	Roles []string
	jwt.RegisteredClaims
}

func GenerateJWTToken(Username string, Roles []string, Secret string) string {
	timeValidInHours := 
	customClaims := CustomClaims{
		Roles,
		jwt.RegisteredClaims{
			ExpiresAt: jwt,
		}
	}
}

func decideTimeValidInHours(tokenType TokenType) int{
	if tokenType == RefreshToken{
		return 24,nil
	}
	if tokenType == AccessToken {
		return 2,nil
	}
	return 0, errors.New("unknown type of token: " + TokenType.typeName)
}



func ValidateJWTToken(token string, secret string) bool {
	return false
}
