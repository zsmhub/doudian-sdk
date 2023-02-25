package sms_sign_search_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsSignSearchResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsSignSearchData `json:"data"`
}
type SignSearchListItem struct {
	// 短信发送渠道，主要做资源隔离
	SmsAccount string `json:"sms_account"`
	// 签名
	Sign string `json:"sign"`
}
type SmsSignSearchData struct {
	// 签名列表
	SignSearchList []SignSearchListItem `json:"sign_search_list"`
	// 总数
	Total int64 `json:"total"`
}
