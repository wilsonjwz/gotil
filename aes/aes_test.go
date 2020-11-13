package aes

import (
	"testing"

	"github.com/iGoogle-ink/gotil/xlog"
)

var (
	secretKey = "GYBh3Rmey7nNzR/NpV0vAw=="
	iv        = "1234567812345678"
)

func TestAesCBCEncryptDecrypt(t *testing.T) {
	originData := "www.gopay.ink"
	xlog.Debug("originData:", originData)
	encryptData, err := AesCBCEncryptData([]byte(originData), []byte(secretKey))
	if err != nil {
		xlog.Error("AesCBCEncryptToString:", err)
		return
	}
	xlog.Debug("encryptData:", string(encryptData))
	origin, err := AesCBCDecryptData(encryptData, []byte(secretKey))
	if err != nil {
		xlog.Error("AesDecryptToBytes:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}

func TestAesCBCEncryptDecryptIv(t *testing.T) {
	originData := "www.gopay.ink"
	xlog.Debug("originData:", originData)
	encryptData, err := AesCBCEncryptIvData([]byte(originData), []byte(secretKey), []byte(iv))
	if err != nil {
		xlog.Error("AesCBCEncryptIvData:", err)
		return
	}
	xlog.Debug("encryptData:", string(encryptData))
	origin, err := AesCBCDecryptIvData(encryptData, []byte(secretKey), []byte(iv))
	if err != nil {
		xlog.Error("AesCBCDecryptIvData:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}
