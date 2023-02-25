package doudian_message

import "encoding/json"

// 订单发货时效变更
// 文档：https://op.jinritemai.com/docs/message-docs/30/1444

func init() {
	MessageDataParseHandle.RegisterMessage(TradeAppointmentV2{})
}

type TradeAppointmentV2 struct {
	ExpDeliveryTime int64   `json:"exp_delivery_time"`
	OrderType       int64   `json:"order_type"`
	PId             string  `json:"p_id"`
	SIds            []int64 `json:"s_ids"`
	ShopId          int64   `json:"shop_id"`
}

func (m TradeAppointmentV2) Tag() string {
	return "126"
}

func (m TradeAppointmentV2) Parse(data []byte) (MessageData, error) {
	var tmp TradeAppointmentV2
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
