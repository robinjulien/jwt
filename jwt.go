package jwt

import (
	"encoding/base64"
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload map[string]interface{}

func Base64URLDecode(data []byte) ([]byte, error) {
	dbuf := make([]byte, base64.RawURLEncoding.DecodedLen(len(data)))
	n, err := base64.RawURLEncoding.Decode(dbuf, data)
	return dbuf[:n], err
}

func Base64URLEncode(data []byte) []byte {
	dbuf := make([]byte, base64.RawURLEncoding.EncodedLen(len(data)))
	base64.RawURLEncoding.Encode(dbuf, data)

	return dbuf
}
