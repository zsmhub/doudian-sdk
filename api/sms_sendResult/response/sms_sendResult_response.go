package sms_sendResult_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SmsSendResultResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SmsSendResultData `json:"data"`
}
type SmsSendResultData struct {
	// 数量
	Total int64 `json:"total"`
	// 短信结果列表
	SmsSendResultList []SmsSendResultListItem `json:"sms_send_result_list"`
}
type SmsSendResultListItem struct {
	// 发送时间-时间戳，单位秒
	SendTime int64 `json:"send_time"`
	// 短信内容
	SmsContent string `json:"sms_content"`
	// 未回执：1 发送失败：2 发送成功：3
	Status int64 `json:"status"`
	// 计费条数，如果短信过长，会分多次计费
	Count int64 `json:"count"`
	// 错误码
	Code int64 `json:"code"`
	// 错误说明
	Message string `json:"message"`
	// 消息ID
	MessageId string `json:"message_id"`
	// 透传字段，回执的时候原样返回给调用方。
	Tag string `json:"tag"`
}
