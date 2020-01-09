package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha1"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

//GetSHAToHex hashes a given string input using the SHA-1 algorithm and returns
//a hexadecimal representation of the same
func GetSHAToHex(value string) string {
	hasher := sha1.New()
	hasher.Write([]byte(value))
	res := hex.EncodeToString(hasher.Sum(nil))
	return res
}

//GenerateRsaKeyPair generates a public and private keypair
func GenerateRsaKeyPair() {
	randomizer := rand.Reader
	key, _ := rsa.GenerateKey(randomizer, 1024)
	str := x509.MarshalPKCS1PublicKey(&key.PublicKey)
	strP := x509.MarshalPKCS1PrivateKey(key)
	fmt.Println(base64.StdEncoding.EncodeToString(str))
	fmt.Println(base64.StdEncoding.EncodeToString(strP))
}
