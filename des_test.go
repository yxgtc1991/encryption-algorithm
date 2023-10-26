package encryption_algorithm

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"fmt"
	"testing"
)

// DES加密
func DesEncrypt(text, key []byte) ([]byte, error) {
	// 创建密码（根据密码加密）
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 区块大小
	blockSize := block.BlockSize()
	// 偏移量
	offset := key[:blockSize]
	// 填充
	textPadded := ZeroPadding(text, blockSize)
	// 加密算法
	blockMode := cipher.NewCBCEncrypter(block, offset)
	encrypted := make([]byte, len(textPadded))
	// 加密
	blockMode.CryptBlocks(encrypted, textPadded)
	return encrypted, nil
}

// DES解密
func DesDecrypt(text, key []byte) ([]byte, error) {
	// 创建密码（根据密码解密）
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	// 定义大小
	blockSize := block.BlockSize()
	// 定义偏移量
	offset := key[:blockSize]
	// 创建加密算法
	blockMode := cipher.NewCBCDecrypter(block, offset)
	// 创建空间
	decrypt := make([]byte, len(text))
	// 解密
	blockMode.CryptBlocks(decrypt, text)
	// 去除填充
	result := UnPadding(decrypt)
	return result, nil
}

func TestDes(t *testing.T) {
	// 定义待加密文本
	str := "hello world"
	// 定义密码（8字节）
	key := []byte("12345678") // 中文占3字节，3*5+1 = 16
	fmt.Println("原文本：", str, "，密码：", string(key))
	// 加密
	encrypted, err := DesEncrypt([]byte(str), key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密结果：", base64.StdEncoding.EncodeToString(encrypted))
	// 解密
	decrypted, err := DesDecrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密结果：", string(decrypted))
}
