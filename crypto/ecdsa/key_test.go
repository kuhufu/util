package ecdsa

import (
	"testing"
)

func Test_loadKey(t *testing.T) {
	text := "hahahah"

	privateKey, err := LoadPrivateKeyFromPemFile("private.pem")
	if err != nil {
		t.Error(err)
	}
	publicKey, err := LoadPublicKeyFromPemFile("public.pem")
	if err != nil {
		t.Error(err)
	}

	sign := Sign(privateKey, text)

	verify := Verify(publicKey, text, sign)

	t.Log(verify)
}

func Test_g(t *testing.T) {
	Generate()

}
