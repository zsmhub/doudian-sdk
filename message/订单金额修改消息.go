package doudian_message

import "encoding/json"

// 订单金额修改消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/265

func init() {
	MessageDataParseHandle.RegisterMessage(TradeAmountChanged{})
}

type TradeAmountChanged struct {
	Biz         int64   `json:"biz"`
	ModifyTime  int64   `json:"modify_time"`
	OrderStatus int64   `json:"order_status"`
	OrderType   int64   `json:"order_type"`
	PId         int64   `json:"p_id"`
	PostAmount  int64   `json:"post_amount"`
	SIds        []int64 `json:"s_ids"`
	ShopId      int64   `json:"shop_id"`
	TotalAmount int64   `json:"total_amount"`
}

func (m TradeAmountChanged) Tag() string {
	return "109"
}

func (m TradeAmountChanged) Parse(data []byte) (MessageData, error) {
	var tmp TradeAmountChanged
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
