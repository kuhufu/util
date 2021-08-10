package crypto

import (
	"testing"
)

//生成密钥公钥文件
func TestGenerateRSAKey(t *testing.T) {
	GenerateRSAKey(2048)
}
