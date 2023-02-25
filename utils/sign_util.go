package utils

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// Sign 计算签名
func Sign(appKey, appSecret, method string, timestamp int64, paramJson string) string {
	// 按给定规则拼接参数
	paramPattern := "app_key" + appKey + "method" + method + "param_json" + paramJson + "timestamp" + strconv.FormatInt(timestamp, 10) + "v2"
	signPattern := appSecret + paramPattern + appSecret
	fmt.Println("sign_pattern:" + signPattern)
	return Hmac(signPattern, appSecret)
}

func SpiSign(appKey, appSecret, timestamp string, sortedParamStr string, signMethod string) string {

	paramPattern := "app_key" + appKey + "param_json" + sortedParamStr +"timestamp" + timestamp
	signPattern := appSecret + paramPattern + appSecret
	if signMethod == "hmac-sha256" {
		return Hmac(signPattern, appSecret)
	} else {
		return Md5(signPattern)
	}
}

// Hmac 计算hmac
func Hmac(s string, appSecret string) string {
	h := hmac.New(sha256.New, []byte(appSecret))
	_, _ = h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5(s string) string {
	h := md5.New()
	_, _ = io.WriteString(h, s)
	return hex.EncodeToString(h.Sum(nil))
}



// Marshal 序列化参数
func Marshal(o interface{}) string {
	// 序列化一次
	raw, _ := json.Marshal(o)

	// 反序列化为map
	m := make(map[string]interface{})
	reader := bytes.NewReader(raw)
	decode := json.NewDecoder(reader)
	decode.UseNumber()
	_ = decode.Decode(&m)

	// 重新做一次序列化，并禁用Html Escape
	buffer := bytes.NewBufferString("")
	encoder := json.NewEncoder(buffer)
	encoder.SetEscapeHTML(false)
	_ = encoder.Encode(m)

	marshal := strings.TrimSpace(buffer.String()) // Trim掉末尾的换行符
	return marshal
}


