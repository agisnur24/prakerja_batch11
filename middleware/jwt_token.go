package middleware

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type jwtCustomClaims struct {
	UserId uint
	Name   string
	jwt.RegisteredClaims
}

func GenerateTokenJWT(userId uint, name string) string {
	var claims = jwtCustomClaims{
		userId,
		name,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	resultJWT, _ := token.SignedString([]byte(os.Getenv("SECRET_JWT")))

	return resultJWT
}
