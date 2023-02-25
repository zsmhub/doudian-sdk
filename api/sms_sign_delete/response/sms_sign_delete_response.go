package sms_sign_delete_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsSignDeleteResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsSignDeleteData `json:"data"`
}
type SmsSignDeleteData struct {
	// 是否成功 0表示成功
	Code int64 `json:"code"`
	// 说明
	Message string `json:"message"`
}
