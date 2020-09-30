package str

import (
	"fmt"
	"testing"
)

func TestBytesToString(t *testing.T) {
	fmt.Println(BytesToString([]byte("sdfdf")))
}

func TestStringToBytes(t *testing.T) {
	fmt.Println(StringToBytes("aaa"))
}
