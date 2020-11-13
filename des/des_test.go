package aes

import (
	"testing"

	"github.com/iGoogle-ink/gotil/xlog"
)

var (
	secretKey = "GYBh3Rmey7nNzR/NpV0vAw=="
	iv        = "JR3unO2glQuMhUx3"
)

func TestDesCBCEncryptDecrypt(t *testing.T) {
	originData := "www.gopay.ink"
	xlog.Debug("originData:", originData)
	encryptData, err := DesCBCEncryptData([]byte(originData), []byte(secretKey))
	if err != nil {
		xlog.Error("DesCBCEncryptData:", err)
		return
	}
	xlog.Debug("encryptData:", string(encryptData))
	origin, err := DesCBCDecryptData(encryptData, []byte(secretKey))
	if err != nil {
		xlog.Error("DesCBCDecryptData:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}

func TestDesCBCEncryptDecryptIv(t *testing.T) {
	originData := "www.gopay.ink"
	xlog.Debug("originData:", originData)
	encryptData, err := DesCBCEncryptIvData([]byte(originData), []byte(secretKey), []byte(iv))
	if err != nil {
		xlog.Error("DesCBCEncryptIvData:", err)
		return
	}
	xlog.Debug("encryptData:", string(encryptData))
	origin, err := DesCBCDecryptIvData(encryptData, []byte(secretKey), []byte(iv))
	if err != nil {
		xlog.Error("DesCBCDecryptIvData:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}
