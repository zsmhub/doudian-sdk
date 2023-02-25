package doudian_message

import (
	"encoding/json"
)

// 买家取消仲裁消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/260

func init() {
	MessageDataParseHandle.RegisterMessage(RefundArbitrateCancelled{})
}

type RefundArbitrateCancelled struct {
	AftersaleId     int64  `json:"aftersale_id"`
	AftersaleStatus int64  `json:"aftersale_status"`
	AftersaleType   int64  `json:"aftersale_type"`
	ArbitrateId     string `json:"arbitrate_id"`
	ArbitrateStatus int64  `json:"arbitrate_status"`
	PId             int64  `json:"p_id"`
	SId             int64  `json:"s_id"`
	ShopId          int64  `json:"shop_id"`
}

func (m RefundArbitrateCancelled) Tag() string {
	return "215"
}

func (m RefundArbitrateCancelled) Parse(data []byte) (MessageData, error) {
	var tmp RefundArbitrateCancelled
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
