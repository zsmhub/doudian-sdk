package doudian_message

import (
	"encoding/json"
)

// 售后超时时长变更消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/226

func init() {
	MessageDataParseHandle.RegisterMessage(RefundExpirationChange{})
}

type RefundExpirationChange struct {
	AftersaleId     int64 `json:"aftersale_id"`
	AftersaleStatus int64 `json:"aftersale_status"`
	AftersaleType   int64 `json:"aftersale_type"`
	PId             int64 `json:"p_id"`
	SId             int64 `json:"s_id"`
	ShopId          int64 `json:"shop_id"`
	StatusDeadline  int64 `json:"status_deadline"`
}

func (m RefundExpirationChange) Tag() string {
	return "209"
}

func (m RefundExpirationChange) Parse(data []byte) (MessageData, error) {
	var tmp RefundExpirationChange
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
