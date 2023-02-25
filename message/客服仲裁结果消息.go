package doudian_message

import (
	"encoding/json"
)

// 客服仲裁结果消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/259

func init() {
	MessageDataParseHandle.RegisterMessage(RefundArbitrateAudited{})
}

type RefundArbitrateAudited struct {
	AftersaleId         int64  `json:"aftersale_id"`
	AftersaleStatus     int64  `json:"aftersale_status"`
	AftersaleType       int64  `json:"aftersale_type"`
	ArbitrateConclusion int64  `json:"arbitrate_conclusion"`
	ArbitrateId         string `json:"arbitrate_id"`
	ArbitrateStatus     int64  `json:"arbitrate_status"`
	PId                 int64  `json:"p_id"`
	SId                 int64  `json:"s_id"`
	ShopId              int64  `json:"shop_id"`
}

func (m RefundArbitrateAudited) Tag() string {
	return "216"
}

func (m RefundArbitrateAudited) Parse(data []byte) (MessageData, error) {
	var tmp RefundArbitrateAudited
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
