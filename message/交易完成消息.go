package doudian_message

import "encoding/json"

// 交易完成消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/112

func init() {
	MessageDataParseHandle.RegisterMessage(TradeSuccess{})
}

type TradeSuccess struct {
	Biz          int64   `json:"biz"`
	CompleteTime int64   `json:"complete_time"`
	OrderStatus  int64   `json:"order_status"`
	OrderType    int64   `json:"order_type"`
	PId          int64   `json:"p_id"`
	SIds         []int64 `json:"s_ids"`
	ShopId       int64   `json:"shop_id"`
}

func (m TradeSuccess) Tag() string {
	return "103"
}

func (m TradeSuccess) Parse(data []byte) (MessageData, error) {
	var tmp TradeSuccess
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
