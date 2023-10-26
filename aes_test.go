package encryption_algorithm

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"testing"
)

// 全零填充
func ZeroPadding(text []byte, blockSize int) []byte {
	// 计算最后一个区块的长度
	lastBlock := len(text) % blockSize
	// 计算填充长度
	paddingLen := blockSize - lastBlock
	// 全零填充
	padding := bytes.Repeat([]byte{0}, paddingLen)
	result := append(text, padding...)
	return result
}

// 去除填充
func UnPadding(encodeText []byte) []byte {
	// 去除尾部的0
	UnPad := bytes.TrimRightFunc(encodeText, func(r rune) bool {
		return r == rune(0)
	})
	return UnPad
}

// AES加密
func AesEncrypt(text, key []byte) ([]byte, error) {
	// 创建密码（根据密码加密）
	block, err := aes.NewCipher(key)
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

// AES解密
func AesDecrypt(text, key []byte) ([]byte, error) {
	// 创建密码（根据密码解密）
	block, err := aes.NewCipher(key)
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

func TestAes(t *testing.T) {
	// 定义待加密文本
	str := "hello world"
	// 定义密码（16/24/32字节）
	key := []byte("我是密码啊!") // 中文占3字节，3*5+1 = 16
	fmt.Println("原文本：", str, "，密码：", string(key))
	// 加密
	encrypted, err := AesEncrypt([]byte(str), key)
	if err != nil {
		panic(err)
	}
	fmt.Println("加密结果：", base64.StdEncoding.EncodeToString(encrypted))
	// 解密
	decrypted, err := AesDecrypt(encrypted, key)
	if err != nil {
		panic(err)
	}
	fmt.Println("解密结果：", string(decrypted))
}
