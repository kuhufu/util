package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

//填充，用在加密
func PKCS7Padding(cipherText []byte, blockSize int) []byte {
	padding := blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherText, padText...)
}

//去填充，用在解密
func PKCS7UnPadding(plainText []byte) []byte {
	length := len(plainText)
	unPadding := int(plainText[length-1])
	return plainText[:(length - unPadding)]
}

//key 长度必须为 16, 24 或 32 字节, 对应 AES-128, AES-192, AES-256
func AesEncrypt(plainText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockSize := block.BlockSize()
	plainText = PKCS7Padding(plainText, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key)
	cipherText := make([]byte, len(plainText))
	blockMode.CryptBlocks(cipherText, plainText)
	return cipherText, nil
}

//key 长度必须为 16, 24 或 32 字节, 对应 AES-128, AES-192, AES-256
func AesDecrypt(cipherText, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	blockMode := cipher.NewCBCDecrypter(block, key)
	plainText := make([]byte, len(cipherText))
	blockMode.CryptBlocks(plainText, cipherText)
	plainText = PKCS7UnPadding(plainText)
	return plainText, nil
}
