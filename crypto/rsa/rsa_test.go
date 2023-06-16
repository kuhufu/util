package rsa

import (
	"testing"
)

func TestGenerateKey(t *testing.T) {
	GenerateKey()
}

func TestRSA_Decrypt(t *testing.T) {
	r := NewRSA(publicKey, privateKey)

	tests := []struct {
		plainText string
	}{
		{"hello world"},
		{"hello"},
		{"apple"},
	}

	for _, test := range tests {
		test := test

		t.Run(test.plainText, func(t *testing.T) {
			encrypt := r.Encrypt([]byte(test.plainText))
			decrypt := r.Decrypt(encrypt)

			if string(decrypt) != test.plainText {
				t.Error("decrypt error")
			}
		})
	}
}
