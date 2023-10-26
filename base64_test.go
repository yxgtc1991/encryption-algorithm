package encryption_algorithm

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestBase64(t *testing.T) {
	str := "james"
	fmt.Println("待加密文本: ", str)
	// 加密
	encode := base64.StdEncoding.EncodeToString([]byte(str))
	fmt.Println(encode)
	fmt.Printf("base64值: %x\n", encode)
	// 解密
	decode, err := base64.StdEncoding.DecodeString(encode)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", decode)
}
