package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
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

func TestSign3(t *testing.T) {
	respStr := `{
"version": "4.0",
"ad-network-id": "com.example",
"source-identifier": "5239",
"app-id": 525463029,
"transaction-id": "6aafb7a5-0170-41b5-bbe4-fe71dedf1e30",
"redownload": false,
"source-domain": "example.com",
"fidelity-type": 1,
"did-win": true,
"conversion-value": 63,
"postback-sequence-index": 0,
"attribution-signature": "MEUCIGRmSMrqedNu6uaHyhVcifs118R5z/AB6cvRaKrRRHWRAiEAv96ne3dKQ5kJpbsfk4eYiePmrZUU6sQmo+7zfP/1Bxo="
}`
	var resp = map[string]interface{}{}
	json.Unmarshal([]byte(respStr), &resp)

	fields := []string{
		"version",
		"ad-network-id",
		"source-identifier",
		"app-id",
		"transaction-id",
		"redownload",
		"source-app-id",
		"source-domain",
		"fidelity-type",
		"did-win",
		"postback-sequence-index",
	}

	var fieldVals []string
	for _, v := range fields {
		fv := resp[v]
		if fv == nil {
			continue
		}
		fieldVals = append(fieldVals, fmt.Sprintf("%v", fv))
	}

	text := strings.Join(fieldVals, "\u2063")

	s := "MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEWdp8GPcGqmhgzEFj9Z2nSpQVddayaPe4FMzqM9wib1+aHaaIzoHoLN9zW4K8y4SPykE3YVK3sVqW6Af0lfx3gg=="
	bytes, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		t.Error(err)
	}

	publicKey, err := LoadPublicKeyFromBytes(bytes)
	if err != nil {
		t.Error(err)
	}

	sign := fmt.Sprintf("%v", resp["attribution-signature"])

	t.Log(Verify(publicKey, text, sign))

}
