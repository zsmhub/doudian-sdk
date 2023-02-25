package doudian_message

import (
	"encoding/json"
)

// 仲裁客服介入消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/627

func init() {
	MessageDataParseHandle.RegisterMessage(RefundArbitrateServiceIntervene{})
}

type RefundArbitrateServiceIntervene struct {
	AftersaleId     int64  `json:"aftersale_id"`
	AftersaleStatus int64  `json:"aftersale_status"`
	AftersaleType   int64  `json:"aftersale_type"`
	ArbitrateId     string `json:"arbitrate_id"`
	ArbitrateStatus int64  `json:"arbitrate_status"`
	PId             int64  `json:"p_id"`
	SId             int64  `json:"s_id"`
	ShopId          int64  `json:"shop_id"`
}

func (m RefundArbitrateServiceIntervene) Tag() string {
	return "217"
}

func (m RefundArbitrateServiceIntervene) Parse(data []byte) (MessageData, error) {
	var tmp RefundArbitrateServiceIntervene
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
