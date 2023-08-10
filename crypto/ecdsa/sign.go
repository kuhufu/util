package ecdsa

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
)

func Sign(privateKey *ecdsa.PrivateKey, text string) string {

	hashed := sha256.Sum256([]byte(text))

	sign, err := ecdsa.SignASN1(rand.Reader, privateKey, hashed[:])
	if err != nil {
		return ""
	}

	//r, s, err := ecdsa.Sign(rand.Reader, privateKey, hashed[:])
	//if err != nil {
	//	return ""
	//}
	//
	//// asn1 output DER format
	//sign, _ := asn1.Marshal(struct {
	//	R *big.Int
	//	S *big.Int
	//}{
	//	R: r,
	//	S: s,
	//})

	return base64.StdEncoding.EncodeToString(sign)
}

func Verify(publicKey *ecdsa.PublicKey, text string, sign string) bool {

	hashed := sha256.Sum256([]byte(text))

	signDec, err := base64.StdEncoding.DecodeString(sign)
	if err != nil {
		return false
	}

	return ecdsa.VerifyASN1(publicKey, hashed[:], signDec)

	//var rs struct{ R, S *big.Int }
	//_, err = asn1.Unmarshal(signDec, &rs)
	//if err != nil {
	//	return false
	//}
	//
	//return ecdsa.Verify(publicKey, hashed[:], rs.R, rs.S)
}
