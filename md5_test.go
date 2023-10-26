package encryption_algorithm

import (
	"crypto/md5"
	"fmt"
	"testing"
)

func TestMd5(t *testing.T) {
	str := "james"
	fmt.Println("待加密文本: ", str)
	// 实例化md5
	md5Obj := md5.New()
	// 写入待加密内容
	md5Obj.Write([]byte(str))
	// 计算md5
	result := md5Obj.Sum(nil)
	fmt.Println(result)
	fmt.Printf("md5值: %x\n", result)
}
