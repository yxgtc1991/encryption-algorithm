package encryption_algorithm

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
)

func GenRsaKey(bits int) error {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		panic("私钥生成失败")
	}
	fmt.Println("私钥：", privateKey)
	// 处理私钥
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 块处理
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	// 文件存储私钥

}
