package utils

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
)

// 16,24,32 分别对应aes-128,192,256
var PwdKey = []byte("DIDLSasdfadf3453")

func PKCS7Padding(cipherText []byte,blockSize int) []byte {
	padding:=blockSize - len(cipherText)%blockSize
	padText := bytes.Repeat([]byte{byte(padding)},padding)
	return append(cipherText,padText...)
}

// 反向填充，删除已填充的字符串
func PKCS7UnPadding(orgData []byte) ([]byte,error) {
	//获取数据长度
	length:=len(orgData)
	if length == 0 {
		return nil, errors.New("加密数据错误")
	}else{
		//获取填充字符串长度
		unpadding:= int(orgData[length-1])
		//删除填充的字符，返回明文
		return orgData[:(length- unpadding)],nil
	}
}

func AesCrypt(orgData []byte,key []byte) ([]byte,error) {
	// 创建加密实例
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	//创建快大小
	blockSize := block.BlockSize()
	// 对加密数据填充，让数据长度满足需求
	orgData = PKCS7Padding(orgData,blockSize)

	//使用加密算法CBC加密模式
	blockMode :=cipher.NewCBCEncrypter(block,key[:blockSize])
	crypted:= make([]byte,len(orgData))
	// 执行加密
	blockMode.CryptBlocks(crypted,orgData)
	return crypted,nil
}

func AesDecrypt(crypted []byte,key []byte) ([]byte,error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize:=block.BlockSize()

	blockMode:=cipher.NewCBCDecrypter(block,key[:blockSize])
	orgData := make([]byte,len(crypted))
	// 加密和解密都是这个函数
	blockMode.CryptBlocks(orgData,crypted)
	//去除填充
	org, err := PKCS7UnPadding(orgData)
	if err != nil {
		return nil, nil
	}
	return org,err
}

func EncryptBase64(data []byte) (string,error) {
	crypted, err := AesCrypt(data, PwdKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(crypted),err
}

func DecryptBase64AndAes(str string) ([]byte,error) {
	decoded, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		return nil, err
	}
	return AesDecrypt(decoded,PwdKey)
}

