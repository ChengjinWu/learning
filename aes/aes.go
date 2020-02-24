package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

var iv = []byte("w123456789w123456789")

var cipherKeyLen = 16

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

//key的长度需要大于等于cipherKeyLen、blockSize
func AesEncrypt(origData []byte, key []byte) (string, error) {
	key = key[:cipherKeyLen]
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return base64.StdEncoding.EncodeToString(crypted), nil
}

func AesDecrypt(crypted string, key []byte) (string, error) {
	key = key[:cipherKeyLen]
	decodeData, err := base64.StdEncoding.DecodeString(crypted)
	if err != nil {
		return "", err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(decodeData))
	blockMode.CryptBlocks(origData, decodeData)
	origData = PKCS5UnPadding(origData)
	return string(origData), nil
}

func main() {
	fmt.Println(AesDecrypt("uuPkXLHSrLtK9B2FCYqV1Q==", []byte("w123456789w123456789")))
	fmt.Println(AesDecrypt("IUqJkwCkTZVsck4Bf4Yypg==", []byte("w123456789w123456789")))
	fmt.Println(AesDecrypt("aHxzZ/gFAabhPMYDpvXmag==", []byte("w123456789w123456789")))
	fmt.Println(AesDecrypt("ceUwNmHvb9u9w5vdNAFEwA==", []byte("w123456789w123456789")))
}
