package doudian_message

import "encoding/json"

// 订单取消消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/123

func init() {
	MessageDataParseHandle.RegisterMessage(TradeCanceled{})
}

type TradeCanceled struct {
	Biz          int64   `json:"biz"`
	CancelReason string  `json:"cancel_reason"`
	CancelTime   int64   `json:"cancel_time"`
	OrderStatus  int64   `json:"order_status"`
	OrderType    int64   `json:"order_type"`
	PId          int64   `json:"p_id"`
	SIds         []int64 `json:"s_ids"`
	ShopId       int64   `json:"shop_id"`
}

func (m TradeCanceled) Tag() string {
	return "106"
}

func (m TradeCanceled) Parse(data []byte) (MessageData, error) {
	var tmp TradeCanceled
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
