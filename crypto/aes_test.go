package crypto

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestAesDecrypt(t *testing.T) {
	var aeskey = []byte("321423u9y8d2fwfl")
	plainText := []byte("vdncloud123456")
	cipherText, err := AesEncrypt(plainText, aeskey)
	if err != nil {
		t.Error(err)
		return
	}

	pass64 := base64.StdEncoding.EncodeToString(cipherText)
	fmt.Printf("加密后:%v\n", pass64)

	bytesPass, err := base64.StdEncoding.DecodeString(pass64)
	if err != nil {
		t.Error(err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		t.Log(err)
		return
	}
	fmt.Printf("解密后:%s\n", tpass)
}
