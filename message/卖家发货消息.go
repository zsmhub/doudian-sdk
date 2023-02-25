package doudian_message

import "encoding/json"

// 卖家发货消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/111

func init() {
	MessageDataParseHandle.RegisterMessage(TradeSellerShip{})
}

type TradeSellerShip struct {
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

func (m TradeSellerShip) Tag() string {
	return "102"
}

func (m TradeSellerShip) Parse(data []byte) (MessageData, error) {
	var tmp TradeSellerShip
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
