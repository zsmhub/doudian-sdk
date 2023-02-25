package doudian_message

import "encoding/json"

// 订单支付/确认消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/110

func init() {
	MessageDataParseHandle.RegisterMessage(TradePaid{})
}

type TradePaid struct {
	Biz         int64   `json:"biz"`
	OrderStatus int64   `json:"order_status"`
	OrderType   int64   `json:"order_type"`
	PId         int64   `json:"p_id"`
	PayAmount   int64   `json:"pay_amount"`
	PayTime     int64   `json:"pay_time"`
	PayType     int64   `json:"pay_type"`
	SIds        []int64 `json:"s_ids"`
	ShopId      int64   `json:"shop_id"`
}

func (m TradePaid) Tag() string {
	return "101"
}

func (m TradePaid) Parse(data []byte) (MessageData, error) {
	var tmp TradePaid
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
