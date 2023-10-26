package encryption_algorithm

import (
	"crypto/sha512"
	"fmt"
	"testing"
)

func TestSha512(t *testing.T) {
	str := "james"
	fmt.Println("待加密文本: ", str)
	// 实例化sha512
	sha512Obj := sha512.New()
	// 写入待加密内容
	sha512Obj.Write([]byte(str))
	// 计算哈希
	result := sha512Obj.Sum(nil)
	fmt.Println(result)
	fmt.Printf("sha512值: %x\n", result)
}
