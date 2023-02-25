package doudian_message

import "encoding/json"

// 抖店商品内容变更消息
// 文档：https://op.jinritemai.com/docs/message-docs/60/2759

func init() {
	MessageDataParseHandle.RegisterMessage(ProductInfoChange{})
}

type ProductInfoChange struct {
	ChannelInfo *struct {
		ChannelId            int64 `json:"channel_id"`
		ChannelMainProductId int64 `json:"channel_main_product_id"`
		ChannelType          int64 `json:"channel_type"`
	} `json:"channel_info"`
	MsgType   int64 `json:"msg_type"`
	ProductId int64 `json:"product_id"`
	ShopId    int64 `json:"shop_id"`
	Status    *struct {
		CheckStatus string `json:"check_status"`
		DraftStatus string `json:"draft_status"`
		Status      string `json:"status"`
	} `json:"status"`
}

func (m ProductInfoChange) Tag() string {
	return "417"
}

func (m ProductInfoChange) Parse(data []byte) (MessageData, error) {
	var tmp ProductInfoChange
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
