package encryption_algorithm

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestSha256(t *testing.T) {
	str := "james"
	fmt.Println("待加密文本: ", str)
	// 实例化sha256
	sha256Obj := sha256.New()
	// 写入待加密内容
	sha256Obj.Write([]byte(str))
	// 计算哈希
	result := sha256Obj.Sum(nil)
	fmt.Println(result)
	fmt.Printf("sha256值: %x\n", result)
}
