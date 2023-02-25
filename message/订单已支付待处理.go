package doudian_message

import "encoding/json"

// 订单已支付待处理
// 文档：https://op.jinritemai.com/docs/message-docs/30/277

func init() {
	MessageDataParseHandle.RegisterMessage(TradePending{})
}

type TradePending struct {
	Biz         int64   `json:"biz"`
	OrderStatus int64   `json:"order_status"`
	OrderType   int64   `json:"order_type"`
	PId         int64   `json:"p_id"`
	SIds        []int64 `json:"s_ids"`
	ShopId      int64   `json:"shop_id"`
}

func (m TradePending) Tag() string {
	return "110"
}

func (m TradePending) Parse(data []byte) (MessageData, error) {
	var tmp TradePending
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
