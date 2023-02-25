package doudian_message

import "encoding/json"

// 发货物流信息变更消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/113

func init() {
	MessageDataParseHandle.RegisterMessage(TradeLogisticsChanged{})
}

type TradeLogisticsChanged struct {
	Biz          int64 `json:"biz"`
	LogisticsMsg *struct {
		ExpressCompanyId string `json:"express_company_id"`
		LogisticsCode    string `json:"logistics_code"`
	} `json:"logistics_msg"`
	OrderStatus int64 `json:"order_status"`
	OrderType   int64 `json:"order_type"`
	PId         int64 `json:"p_id"`
	ReceiverMsg *struct {
		Addr string `json:"addr"`
		Name string `json:"name"`
		Tel  string `json:"tel"`
	} `json:"receiver_msg"`
	SIds       []int64 `json:"s_ids"`
	ShopId     int64   `json:"shop_id"`
	UpdateTime int64   `json:"update_time"`
}

func (m TradeLogisticsChanged) Tag() string {
	return "104"
}

func (m TradeLogisticsChanged) Parse(data []byte) (MessageData, error) {
	var tmp TradeLogisticsChanged
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
