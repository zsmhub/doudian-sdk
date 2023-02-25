package doudian_message

import (
	"crypto/md5"
	"encoding/hex"
)

func VerifySign(body []byte, eventSign string) bool {
	signParam := GlobalConfig.AppKey + string(body) + GlobalConfig.AppSecret
	sign := Md5(signParam)
	return sign == eventSign
}

func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
