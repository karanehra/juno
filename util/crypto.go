package util

import (
	"crypto/sha1"
	"encoding/hex"
)

//GetSHAToHex hashes a given string input using the SHA-1 algorithm and returns
//a hexadecimal representation of the same
func GetSHAToHex(value string) string {
	hasher := sha1.New()
	hasher.Write([]byte(value))
	res := hex.EncodeToString(hasher.Sum(nil))
	return res
}
