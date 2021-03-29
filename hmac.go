package jwt

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/sha512"
	"errors"
	"hash"
)

var (
	InvalidAlg error = errors.New("Invalid Alg")
)

func ValidMAC(message, messageMAC, key []byte, hashfn func() hash.Hash) bool {
	mac := hmac.New(hashfn, key)
	mac.Write(message)
	expectedMAC := mac.Sum(nil)
	return hmac.Equal(messageMAC, expectedMAC)
}

func GetMac(message, key []byte, hashfn func() hash.Hash) []byte {
	mac := hmac.New(hashfn, key)
	mac.Write(message)
	return mac.Sum(nil)
}

func GetHashFn(s string) (func() hash.Hash, error) {
	switch s {
	case "HS256":
		return sha256.New, nil
	case "HS384":
		return sha512.New384, nil
	case "HS512":
		return sha512.New, nil
	default:
		return sha256.New, InvalidAlg
	}
}
