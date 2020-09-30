package str

import (
	"reflect"
	"unsafe"
)

func StringToBytes(s string) []byte {
	sh := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bh := reflect.SliceHeader{Data: sh.Data, Len: sh.Len, Cap: sh.Len}
	return *(*[]byte)(unsafe.Pointer(&bh))
}

func BytesToString(bytes []byte) string {
	return *(*string)(unsafe.Pointer(&bytes))
}
