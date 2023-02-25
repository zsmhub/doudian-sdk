package doudian_message

import "encoding/json"

// 商品变更消息
// 文档：https://op.jinritemai.com/docs/message-docs/60/686

func init() {
	MessageDataParseHandle.RegisterMessage(ProductChange{})
}

type ProductChange struct {
	CheckRejectReason string `json:"check_reject_reason"`
	Event             int64  `json:"event"`
	OutProductId      int64  `json:"out_product_id"`
	ProductId         int64  `json:"product_id"`
	ShopId            int64  `json:"shop_id"`
	UpdateTime        int64  `json:"update_time"`
}

func (m ProductChange) Tag() string {
	return "400"
}

func (m ProductChange) Parse(data []byte) (MessageData, error) {
	var tmp ProductChange
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
