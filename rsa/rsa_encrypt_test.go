package rsa

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"io"
	"os"
	"testing"
)

var (
	publicKey  = []byte("")
	privateKey = []byte("")
)

func ReadKey() {
	// 文件路径
	publicKeyPath := "RSA/RSA_generate_key/rsa_public_key.pem"
	privateKeyPath := "RSA/RSA_generate_key/rsa_private_key.pem"
	// 打开文件
	public, err := os.Open(publicKeyPath)
	if err != nil {
		panic(err)
	}
	defer public.Close()
	private, err := os.Open(privateKeyPath)
	if err != nil {
		panic(err)
	}
	defer private.Close()
	// 读取内容
	publicKey, err = io.ReadAll(public)
	if err != nil {
		panic(err)
	}
	privateKey, err = io.ReadAll(private)
	if err != nil {
		panic(err)
	}
}

func RSAEncrypt(text []byte) ([]byte, error) {
	// pem编码
	block, _ := pem.Decode(publicKey)
	if block == nil {
		return nil, errors.New("私钥无效")
	}
	// 获取公钥
	publicInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 断言类型转换（interface -> *rsa.PublicKey）
	public := publicInterface.(*rsa.PublicKey)
	// 加密
	encode, err := rsa.EncryptPKCS1v15(rand.Reader, public, text)
	if err != nil {
		return nil, err
	}
	return encode, nil
}

func RSADecrypt(text []byte) ([]byte, error) {
	// pem解码
	block, _ := pem.Decode(privateKey)
	if block == nil {
		return nil, errors.New("私钥无效")
	}
	// x509反序列化，获取私钥
	private, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	// 解密
	decrypt, err := rsa.DecryptPKCS1v15(rand.Reader, private, text)
	if err != nil {
		return nil, err
	}
	return decrypt, nil
}

func TestRsaEncrypt(t *testing.T) {
	// 读取公钥秘钥
	ReadKey()
	// 待加密文本
	text := "james"
	// 加密
	encrypted, err := RSAEncrypt([]byte(text))
	if err != nil {
		panic(err)
	}
	fmt.Println("加密：", base64.StdEncoding.EncodeToString(encrypted))
	// 解密
	result, err := RSADecrypt(encrypted)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密：", string(result))
}
