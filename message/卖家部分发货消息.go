package doudian_message

import "encoding/json"

// 卖家部分发货消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/208

func init() {
	MessageDataParseHandle.RegisterMessage(TradePartlySellerShip{})
}

type TradePartlySellerShip struct {
	Biz          int64 `json:"biz"`
	LogisticsMsg *struct {
		ExpressCompanyId string `json:"express_company_id"`
		LogisticsCode    string `json:"logistics_code"`
		PackId           string `json:"pack_id"`
		ShippedOrderInfo []struct {
			ProductId      int64 `json:"product_id"`
			ShippedNum     int64 `json:"shipped_num"`
			ShippedOrderId int64 `json:"shipped_order_id"`
			SkuId          int64 `json:"sku_id"`
			SkuNum         int64 `json:"sku_num"`
		} `json:"shipped_order_info"`
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

func (m TradePartlySellerShip) Tag() string {
	return "108"
}

func (m TradePartlySellerShip) Parse(data []byte) (MessageData, error) {
	var tmp TradePartlySellerShip
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
