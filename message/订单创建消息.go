package doudian_message

import "encoding/json"

// 订单创建消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/109

func init() {
	MessageDataParseHandle.RegisterMessage(TradeCreate{})
}

type TradeCreate struct {
	Biz         int64   `json:"biz"`
	CreateTime  int64   `json:"create_time"`
	OrderStatus int64   `json:"order_status"`
	OrderType   int64   `json:"order_type"`
	PId         int64   `json:"p_id"`
	SIds        []int64 `json:"s_ids"`
	ShopId      int64   `json:"shop_id"`
}

func (m TradeCreate) Tag() string {
	return "100"
}

func (m TradeCreate) Parse(data []byte) (MessageData, error) {
	var tmp TradeCreate
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
