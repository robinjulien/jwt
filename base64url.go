package jwt

import "encoding/base64"

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
