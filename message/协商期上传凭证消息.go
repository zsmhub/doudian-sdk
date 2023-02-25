package doudian_message

import (
	"encoding/json"
)

// 协商期上传凭证消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/628

func init() {
	MessageDataParseHandle.RegisterMessage(RefundArbitrateDiscussUpload{})
}

type RefundArbitrateDiscussUpload struct {
	AftersaleId     int64  `json:"aftersale_id"`
	AftersaleStatus int64  `json:"aftersale_status"`
	AftersaleType   int64  `json:"aftersale_type"`
	ArbitrateId     string `json:"arbitrate_id"`
	ArbitrateStatus int64  `json:"arbitrate_status"`
	PId             int64  `json:"p_id"`
	SId             int64  `json:"s_id"`
	ShopId          int64  `json:"shop_id"`
}

func (m RefundArbitrateDiscussUpload) Tag() string {
	return "218"
}

func (m RefundArbitrateDiscussUpload) Parse(data []byte) (MessageData, error) {
	var tmp RefundArbitrateDiscussUpload
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
