package doudian_message

import "encoding/json"

// 买家收货信息变更申请消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/285

func init() {
	MessageDataParseHandle.RegisterMessage(TradeAmountChangedApplied{})
}

type TradeAmountChangedApplied struct {
	ApplyTime       int64 `json:"apply_time"`
	PId             int64 `json:"p_id"`
	PostReceiverMsg *struct {
		Addr        string       `json:"addr"`
		AddrStruct  ReceiverAddr `json:"addr_struct"` // 自定义字段，方便获取地址内容
		Name        string       `json:"name"`
		EncryptName string       `json:"encrypt_name"`
		Tel         string       `json:"tel"`
		EncryptTel  string       `json:"encrypt_tel"`
	} `json:"post_receiver_msg"`
	ReceiverMsg *struct {
		Addr        string       `json:"addr"`
		AddrStruct  ReceiverAddr `json:"addr_struct"` // 自定义字段，方便获取地址内容
		Name        string       `json:"name"`
		EncryptName string       `json:"encrypt_name"`
		Tel         string       `json:"tel"`
		EncryptTel  string       `json:"encrypt_tel"`
	} `json:"receiver_msg"`
	ShopId int64 `json:"shop_id"`
}

func (m TradeAmountChangedApplied) Tag() string {
	return "111"
}

func (m TradeAmountChangedApplied) Parse(data []byte) (MessageData, error) {
	var tmp TradeAmountChangedApplied
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return tmp, err
	}
	tmp.ReceiverMsg.AddrStruct, err = new(ReceiverAddr).Parse(tmp.ReceiverMsg.Addr)
	if err != nil {
		return tmp, err
	}
	tmp.PostReceiverMsg.AddrStruct, err = new(ReceiverAddr).Parse(tmp.PostReceiverMsg.Addr)
	return tmp, err
}
