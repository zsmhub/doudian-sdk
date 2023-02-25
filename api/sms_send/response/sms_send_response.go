package sms_send_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsSendResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsSendData `json:"data"`
}
type SmsSendData struct {
	// 错误码，0表示成功
	Code int64 `json:"code"`
	// 说明
	Message string `json:"message"`
	// 消息的唯一标识，可以用于查询短信到达等
	MessageId string `json:"message_id"`
}
