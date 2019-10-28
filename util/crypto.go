package util

import "crypto/sha1"

//CreateHashSHA hashes a given string input using the SHA-1 algorithm
func CreateHashSHA(value string) string {
	hasher := sha1.New()
	hasher.Write([]byte(value))
	res := string(hasher.Sum(nil))
	return res
}
