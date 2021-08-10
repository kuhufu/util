package str

import (
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBytesToString(t *testing.T) {
	fmt.Println(BytesToString([]byte("sdfdf")))
}

func TestStringToBytes(t *testing.T) {
	fmt.Println(StringToBytes("aaa"))

	sum := sha256.Sum256([]byte("asdfsd"))

	toString := base64.StdEncoding.EncodeToString(sum[:])

	fmt.Println(len(sum))
	fmt.Println(len([]byte(toString)))
}
