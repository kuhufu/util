package rsa

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
)

func GenerateKey() {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		fmt.Println(err)
	}

	encrypt, err := rsa.EncryptOAEP(
		sha256.New(),
		rand.Reader,
		&key.PublicKey,
		[]byte("test"),
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(base64.StdEncoding.EncodeToString(encrypt))

	decrypt, err := key.Decrypt(nil, encrypt, &rsa.OAEPOptions{Hash: crypto.SHA256})
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(decrypt))

}

type RSA struct {
	publicPem  string
	privatePem string

	privateKey *rsa.PrivateKey
	publicKey  *rsa.PublicKey
}

func NewRSA(publicPem, privatePem string) *RSA {
	block, _ := pem.Decode([]byte(publicKey))
	publicKeyI, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	publicKey := publicKeyI.(*rsa.PublicKey)

	block, _ = pem.Decode([]byte(privateKey))
	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		panic(err)
	}

	return &RSA{
		publicPem:  publicPem,
		privatePem: privatePem,
		privateKey: privateKey,
		publicKey:  publicKey,
	}
}

func (r *RSA) Encrypt(plainText []byte) []byte {
	bytes, err := rsa.EncryptPKCS1v15(rand.Reader, r.publicKey, plainText)
	if err != nil {
		panic(err)
	}

	return bytes
}

func (r *RSA) Decrypt(cipher []byte) []byte {
	bytes, err := rsa.DecryptPKCS1v15(rand.Reader, r.privateKey, cipher)
	if err != nil {
		panic(err)
	}

	return bytes
}
