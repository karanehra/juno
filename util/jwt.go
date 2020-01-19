package util

import (
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

//CreateJWTForIssuer signs and returs a JWT for the issuer
func CreateJWTForIssuer(issuer string) (string, error) {
	jwtTimeout, _ := strconv.Atoi(os.Getenv("JWT_VALIDITY_SECONDS"))
	claims := jwt.StandardClaims{
		ExpiresAt: time.Now().Unix() + int64(jwtTimeout),
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
