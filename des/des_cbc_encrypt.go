package aes

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
)

// 加密后转成Base64字符串
func DesCBCEncryptToString(originData []byte, secretKey string) (string, error) {
	bytes, err := encrypt(originData, secretKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// 加密后的Bytes数组
func DesCBCEncryptToBytes(originData []byte, secretKey string) ([]byte, error) {
	return encrypt(originData, secretKey)
}

func encrypt(originData []byte, secretKey string) ([]byte, error) {
	key := []byte(secretKey)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	originData = PKCS7Padding(originData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	secretData := make([]byte, len(originData))
	blockMode.CryptBlocks(secretData, originData)
	return secretData, nil
}
