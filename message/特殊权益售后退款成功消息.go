package doudian_message

import (
	"encoding/json"
)

// 特殊权益售后退款成功消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/1371

func init() {
	MessageDataParseHandle.RegisterMessage(RefundSpecialRefundSuccess{})
}

type RefundSpecialRefundSuccess struct {
	AftersaleId     int64 `json:"aftersale_id"`
	AftersaleStatus int64 `json:"aftersale_status"`
	AftersaleType   int64 `json:"aftersale_type"`
	PId             int64 `json:"p_id"`
	RefundAmount    int64 `json:"refund_amount"`
	SId             int64 `json:"s_id"`
	ShopId          int64 `json:"shop_id"`
	SuccessTime     int64 `json:"success_time"`
}

func (m RefundSpecialRefundSuccess) Tag() string {
	return "225"
}

func (m RefundSpecialRefundSuccess) Parse(data []byte) (MessageData, error) {
	var tmp RefundSpecialRefundSuccess
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
