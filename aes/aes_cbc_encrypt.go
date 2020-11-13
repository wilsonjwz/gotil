package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
)

// 加密后转成Base64字符串
func AesCBCEncryptToString(originData []byte, secretKey string) (string, error) {
	bytes, err := encrypt(originData, secretKey)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(bytes), nil
}

// 加密后的Bytes数组
func AesCBCEncryptToBytes(originData []byte, secretKey string) ([]byte, error) {
	return encrypt(originData, secretKey)
}

func encrypt(originData []byte, secretKey string) ([]byte, error) {
	key := []byte(secretKey)
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	originData = PKCS5Padding(originData, blockSize)

	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	secretData := make([]byte, len(originData))
	blockMode.CryptBlocks(secretData, originData)
	return secretData, nil
}
