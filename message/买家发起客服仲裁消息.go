package doudian_message

import (
	"encoding/json"
)

// 买家发起客服仲裁消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/256

func init() {
	MessageDataParseHandle.RegisterMessage(RefundArbitrateApplied{})
}

type RefundArbitrateApplied struct {
	AftersaleId     int64  `json:"aftersale_id"`
	AftersaleStatus int64  `json:"aftersale_status"`
	AftersaleType   int64  `json:"aftersale_type"`
	ArbitrateId     string `json:"arbitrate_id"`
	ArbitrateStatus int64  `json:"arbitrate_status"`
	PId             int64  `json:"p_id"`
	SId             int64  `json:"s_id"`
	ShopId          int64  `json:"shop_id"`
}

func (m RefundArbitrateApplied) Tag() string {
	return "212"
}

func (m RefundArbitrateApplied) Parse(data []byte) (MessageData, error) {
	var tmp RefundArbitrateApplied
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
