package jwt

import (
	"bytes"
	"testing"
)

func TestBase64URLEncode(t *testing.T) {
	in := [][]byte{
		[]byte("azertyuiop"),
		[]byte("az+/-_"),
	}
	out := [][]byte{
		[]byte("YXplcnR5dWlvcA"),
		[]byte("YXorLy1f"),
	}
	for i, v := range in {
		res := Base64URLEncode(v)
		if !bytes.Equal(out[i], res) {
			t.Errorf("Error Base64URLEncode index %d, got %s.", i, string(res))
		}
	}
}

func TestBase64URLDecode(t *testing.T) {
	in := [][]byte{
		[]byte("YXplcnR5dWlvcA"),
		[]byte("YXorLy1f"),
	}
	out := [][]byte{
		[]byte("azertyuiop"),
		[]byte("az+/-_"),
	}
	for i, v := range in {
		res, err := Base64URLDecode(v)
		if !bytes.Equal(out[i], res) || err != nil {
			t.Errorf("Error Base64URLDecode index %d.", i)
		}
	}
}
