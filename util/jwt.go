package util

import "github.com/dgrijalva/jwt-go"

type CustomJWTClaims struct {
}

func CreateJWTForIssuer(issuer string) {
	claims := jwt.StandardClaims{
		ExpiresAt: 10000,
		Issuer:    issuer,
	}
	token := jwt.NewWithClaims(jwt.SigningMethodRS512, claims)
}
