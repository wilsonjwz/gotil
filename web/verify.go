package web

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/iGoogle-ink/goutil"
	"github.com/iGoogle-ink/goutil/ecode"
)

const (
	sign  = "sign"
	ts    = "ts"
	appid = "appid"
)

// VerifySign 验证Sign，签名规则，base64(md5(appid+path+ts))
func VerifySign(c *gin.Context) {
	ts := c.GetHeader(ts)
	tsTime := time.Unix(goutil.String2Int64(ts), 0)
	if time.Now().Sub(tsTime).Seconds() > 60 {
		JSON(c, nil, eno.InvalidSignErr)
		c.Abort()
		return
	}
	sign := c.GetHeader(sign)
	appid := c.GetHeader(appid)
	path := c.Request.RequestURI
	split := strings.Split(path, "?")
	if !CheckSign(sign, appid, split[0], ts) {
		JSON(c, nil, eno.InvalidSignErr)
		c.Abort()
		return
	}
}

func CheckSign(sign, appid, path, ts string) (ok bool) {
	hash := md5.New()
	hash.Write([]byte(appid))
	hash.Write([]byte(path))
	hash.Write([]byte(ts))
	calSign := hex.EncodeToString(hash.Sum(nil))
	return sign == calSign
}