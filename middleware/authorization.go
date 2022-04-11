package middleware

import (
	"github.com/golang-jwt/jwt"
	"os"
	"strings"
)

func SplitToken(headerToken string) string {
	parsToken := strings.SplitAfter(headerToken, " ")
	tokenString := parsToken[1]
	return tokenString
}

func AuthenticateToken(tokenString string) error {
	//token check
	_, err := jwt.ParseWithClaims(tokenString, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})
	if err != nil {
		return err
	}
	return nil

}
