package doudian_message

import "encoding/json"

// 卖家添加备注消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/571

func init() {
	MessageDataParseHandle.RegisterMessage(TradeMemoModify{})
}

type TradeMemoModify struct {
	PId        int64  `json:"p_id"`
	Remark     string `json:"remark"`
	ShopId     int64  `json:"shop_id"`
	Star       int64  `json:"star"`
	UpdateTime int64  `json:"update_time"`
}

func (m TradeMemoModify) Tag() string {
	return "113"
}

func (m TradeMemoModify) Parse(data []byte) (MessageData, error) {
	var tmp TradeMemoModify
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
