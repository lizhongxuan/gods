/*
@Time : 2020-07-11 10:28
@Author : zhongxuanli
@File : 2_singleton
@Software: GoLand
*/


package encrpt

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"encoding/hex"
	"fmt"
)
/**
	Go 加密包里面有des,
	参考了这个博文:http://blog.studygolang.com/2013/01/go%E5%8A%A0%E5%AF%86%E8%A7%A3%E5%AF%86%E4%B9%8Bdes
 */
func main1() {
	key := []byte("example key 1234")
	key = key[:8]

	encryptStr, err := DesEncryptStrToStr("hello 世界 !", key)
	if err != nil {
		fmt.Println("des encrypt fail", err)
	} else {
		fmt.Println("des encrypt result : ", encryptStr)

		decryptStr, err := DesDecryptStrFromStr(encryptStr, key)
		if err != nil {
			fmt.Println("des decrypt fail", err)
		} else {
			fmt.Println("des decrypt result : ", decryptStr)
		}
	}
}

func DesEncryptStrToStr(origData string, key []byte) (string, error) {
	encryptBytes, err := DesEncrypt([]byte(origData), key)
	if err != nil {
		return "", err
	} else {
		return hex.EncodeToString(encryptBytes), nil
	}
}

func DesDecryptStrFromStr(crypted string, key []byte) (string, error) {
	encryptBytes, err := hex.DecodeString(crypted)
	if err != nil {
		return "", err
	} else {
		decryptBytes, err := DesDecrypt(encryptBytes, key)
		if err != nil {
			return "", err
		} else {
			return string(decryptBytes), nil
		}
	}
}

func DesEncrypt(origData, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = PKCS5Padding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	crypted := make([]byte, len(origData))
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext) % blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func DesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := des.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	origData := make([]byte, len(crypted))
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length - 1])
	return origData[:(length - unpadding)]
}
