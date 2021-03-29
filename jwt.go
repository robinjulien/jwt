package jwt

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
)

var (
	InvalidInput   error = errors.New("Invalid input. Cannot convert into JWT.")
	InvalidHeader  error = errors.New("Invalid header.")
	InvalidPayload error = errors.New("Invalid payload.")
)

type Header struct {
	Alg string `json:"alg"`
	Typ string `json:"typ"`
}

type Payload map[string]interface{}

type RawJWT []byte

type JWT struct {
	Header    Header
	Payload   Payload
	Signature []byte
}

func VerifyRaw(raw []byte, key []byte) bool {
	parts := bytes.Split(raw, []byte("."))

	if len(parts) != 3 {
		return false
	}

	var header Header
	headerJSON, err := Base64URLDecode(parts[0])

	if err != nil {
		return false
	}

	err = json.Unmarshal(headerJSON, &header)

	if err != nil {
		return false
	}

	hashfn, err := GetHashFn(header.Alg)

	if err != nil {
		return false
	}

	var message []byte
	message = append(message, parts[0]...)
	message = append(message, '.')
	message = append(message, parts[1]...)

	fmt.Println(string(message), string(parts[2]))

	signature, err := Base64URLDecode(parts[2])

	if err != nil {
		return false
	}

	return ValidMAC(message, signature, key, hashfn)
}
