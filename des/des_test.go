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
	encryptData, err := DesCBCEncryptToString([]byte(originData), secretKey)
	if err != nil {
		xlog.Error("DesCBCEncryptToString:", err)
		return
	}
	xlog.Debug("encryptData:", encryptData)
	origin, err := DesDecryptToBytes(encryptData, secretKey)
	if err != nil {
		xlog.Error("DesDecryptToBytes:", err)
		return
	}
	xlog.Debug("origin:", string(origin))
}
