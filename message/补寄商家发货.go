package doudian_message

import (
	"encoding/json"
	"time"
)

// 补寄商家发货
// 文档：https://op.jinritemai.com/docs/message-docs/31/1983

func init() {
	MessageDataParseHandle.RegisterMessage(RefundResendFillLogistics{})
}

type RefundResendFillLogistics struct {
	AftersaleId     int64     `json:"aftersale_id"`
	AftersaleStatus int64     `json:"aftersale_status"`
	AftersaleType   int64     `json:"aftersale_type"`
	PId             int64     `json:"p_id"`
	ReasonCode      int64     `json:"reason_code"`
	SId             int64     `json:"s_id"`
	ShopId          int64     `json:"shop_id"`
	UpdateTime      time.Time `json:"update_time"`
}

func (m RefundResendFillLogistics) Tag() string {
	return "226"
}

func (m RefundResendFillLogistics) Parse(data []byte) (MessageData, error) {
	var tmp RefundResendFillLogistics
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
