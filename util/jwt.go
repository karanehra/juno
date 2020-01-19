package util

import (
	"os"

	"github.com/dgrijalva/jwt-go"
)

//CreateJWTForIssuer signs and returs a JWT for the issuer
func CreateJWTForIssuer(issuer string) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    issuer,
	}
	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	key := []byte(os.Getenv("JWT_SECRET"))
	return newJwt.SignedString(key)
}

//IsJWTValid checks a jwt against the secret
func IsJWTValid(token string) bool {
	parsedToken, _ := jwt.ParseWithClaims(token, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	return parsedToken.Valid
}
