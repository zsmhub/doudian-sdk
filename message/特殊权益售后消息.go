package doudian_message

import (
	"encoding/json"
)

// 特殊权益售后消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/1372

func init() {
	MessageDataParseHandle.RegisterMessage(RefundSpecialRefund{})
}

type RefundSpecialRefund struct {
	AftersaleId     int64 `json:"aftersale_id"`
	AftersaleStatus int64 `json:"aftersale_status"`
	AftersaleType   int64 `json:"aftersale_type"`
	ApplyTime       int64 `json:"apply_time"`
	PId             int64 `json:"p_id"`
	RefundAmount    int64 `json:"refund_amount"`
	SId             int64 `json:"s_id"`
	ShopId          int64 `json:"shop_id"`
}

func (m RefundSpecialRefund) Tag() string {
	return "224"
}

func (m RefundSpecialRefund) Parse(data []byte) (MessageData, error) {
	var tmp RefundSpecialRefund
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
