package aes

import (
	"crypto/cipher"
	"crypto/des"
	"encoding/base64"
	"encoding/json"
	"errors"
	"reflect"
)

// 解密数据的Bytes数组
func DesDecryptToBytes(base64Data, secretKey string) ([]byte, error) {
	return decrypt(base64Data, secretKey)
}

// 解密数据到结构体
func DesDecryptToStruct(base64Data, secretKey string, beanPtr interface{}) (err error) {
	//验证参数类型
	beanValue := reflect.ValueOf(beanPtr)
	if beanValue.Kind() != reflect.Ptr {
		return errors.New("传入参数类型必须是以指针形式")
	}
	//验证interface{}类型
	if beanValue.Elem().Kind() != reflect.Struct {
		return errors.New("传入interface{}必须是结构体")
	}
	originByte, err := decrypt(base64Data, secretKey)
	if err != nil {
		return err
	}
	//解析
	err = json.Unmarshal(originByte, beanPtr)
	if err != nil {
		return err
	}
	return nil
}

// 解密数据到Map集合
func DesDecryptToMap(base64Data, secretKey string) (mapData map[string]interface{}, err error) {
	originByte, err := decrypt(base64Data, secretKey)
	if err != nil {
		return nil, err
	}
	//解析
	mapData = make(map[string]interface{}, 0)
	err = json.Unmarshal(originByte, &mapData)
	if err != nil {
		return nil, err
	}
	return mapData, nil
}

func decrypt(data, secretKey string) (originByte []byte, err error) {
	secretData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	key := []byte(secretKey)
	block, err := des.NewTripleDESCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	originByte = make([]byte, len(secretData))
	blockMode.CryptBlocks(originByte, secretData)
	return PKCS7UnPadding(originByte), nil
}
