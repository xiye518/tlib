package AES

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

/*
AES加密方式，CBC模式，PKCS7填充
*/
type AESCrypt struct{}

func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padText...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unPadding := int(origData[length-1])
	return origData[:(length - unPadding)]
}

// 加密
func (a *AESCrypt) Encrypt(origData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS7Padding(origData, aes.BlockSize)
	iv := key[:aes.BlockSize] // iv不同，则加密后的数据也不同
	blockMode := cipher.NewCBCEncrypter(block, iv)
	encryptData := make([]byte, len(origData))
	blockMode.CryptBlocks(encryptData, origData)
	return encryptData, nil
}

// 解密
func (a *AESCrypt) Decrypt(encryptData, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	iv := key[:aes.BlockSize]
	blockMode := cipher.NewCBCDecrypter(block, iv)
	origData := make([]byte, len(encryptData))
	blockMode.CryptBlocks(origData, encryptData)
	origData = PKCS7UnPadding(origData)
	return origData, nil
}
