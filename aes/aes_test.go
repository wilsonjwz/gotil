package aes

import (
	"testing"

	"github.com/iGoogle-ink/gotil/xlog"
)

var (
	secretKey = "GYBh3Rmey7nNzR/NpV0vAw=="
)

func TestAesCBCEncryptToString(t *testing.T) {
	originData := "www.gopay.ink"
	xlog.Debug("originData:", originData)
	encryptData, err := AesCBCEncryptToString([]byte(originData), secretKey)
	if err != nil {
		xlog.Error("AesCBCEncryptToString:", err)
		return
	}
	xlog.Debug("encryptData:", encryptData)
	origin, err := AesDecryptToBytes(encryptData, secretKey)
	if err != nil {
		xlog.Error("AesDecryptToBytes:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}
