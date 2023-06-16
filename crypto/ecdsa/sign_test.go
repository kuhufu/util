package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"testing"
)

func TestSign(t *testing.T) {
	text := "hahahah"

	privateKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		t.Error(err)
	}

	sign := Sign(privateKey, text)

	publicKey := privateKey.PublicKey
	verify := Verify(&publicKey, text, sign)

	t.Log(verify)
}
