package util

import (
	"fmt"
	"os"

	"github.com/dgrijalva/jwt-go"
)

//CreateJWTForIssuer signs and returs a JWT for the issuer
func CreateJWTForIssuer(issuer string) (string, error) {
	claims := jwt.StandardClaims{
		ExpiresAt: 10000,
		Issuer:    issuer,
	}
	newJwt := jwt.NewWithClaims(jwt.SigningMethodHS512, claims)
	key := []byte(os.Getenv("JWT_SECRET"))
	fmt.Println(key)
	return newJwt.SignedString(key)
}
