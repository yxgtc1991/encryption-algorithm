package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"flag"
	"fmt"
	"os"
	"testing"
)

func GenRsaKey(bits int) error {
	// 生成私钥
	privateKey, err := rsa.GenerateKey(rand.Reader, bits)
	if err != nil {
		fmt.Println("生成私钥失败")
		return err
	}
	fmt.Println("私钥：", privateKey)
	// 处理私钥
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	// 块处理
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	// 创建文件存储私钥
	privateKeyFile, err := os.Create("RSA/RSA_generate_key/rsa_private_key.pem")
	if err != nil {
		fmt.Println("创建私钥文件失败")
		return err
	}
	// 延迟关闭
	defer privateKeyFile.Close()
	// pem编码，写入文件
	err = pem.Encode(privateKeyFile, block)
	if err != nil {
		fmt.Println("写入私钥文件失败")
		return err
	}
	// 生成公钥
	publicKey := &privateKey.PublicKey
	fmt.Println("公钥：", publicKey)
	// 处理公钥
	derPkix, err := x509.MarshalPKIXPublicKey(publicKey)
	if err != nil {
		fmt.Println("处理公钥失败")
		return err
	}
	// 按照块处理
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	// 创建文件存储公钥
	publicKeyFile, err := os.Create("RSA/RSA_generate_key/rsa_public_key.pem")
	if err != nil {
		fmt.Println("创建公钥文件失败")
		return err
	}
	// 延迟关闭
	defer publicKeyFile.Close()
	// pem编码，写入文件
	err = pem.Encode(publicKeyFile, block)
	if err != nil {
		fmt.Println("写入公钥文件失败")
		return err
	}
	return nil
}

func TestGenRsaKey(t *testing.T) {
	// 生成密钥
	var bits int
	flag.IntVar(&bits, "b", 1024, "密码长度默认1024")
	// 生成私钥
	err := GenRsaKey(bits)
	if err != nil {
		panic(err)
	}
}
