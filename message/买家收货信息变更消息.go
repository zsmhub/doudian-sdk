package doudian_message

import "encoding/json"

// 买家收货信息变更消息
// 文档：https://op.jinritemai.com/docs/message-docs/30/114

func init() {
	MessageDataParseHandle.RegisterMessage(TradeAddressChanged{})
}

type TradeAddressChanged struct {
	OrderStatus int64 `json:"order_status"`
	OrderType   int64 `json:"order_type"`
	PId         int64 `json:"p_id"`
	ReceiverMsg struct {
		Addr        string       `json:"addr"`
		AddrStruct  ReceiverAddr `json:"addr_struct"` // 自定义字段，方便获取地址内容
		Name        string       `json:"name"`
		EncryptName string       `json:"encrypt_name"`
		Tel         string       `json:"tel"`
		EncryptTel  string       `json:"encrypt_tel"`
	} `json:"receiver_msg"`
	SIds       []int64 `json:"s_ids"`
	ShopId     int64   `json:"shop_id"`
	UpdateTime int64   `json:"update_time"`
}

func (m TradeAddressChanged) Tag() string {
	return "105"
}

func (m TradeAddressChanged) Parse(data []byte) (MessageData, error) {
	var tmp TradeAddressChanged
	err := json.Unmarshal(data, &tmp)
	if err != nil {
		return tmp, err
	}
	tmp.ReceiverMsg.AddrStruct, err = new(ReceiverAddr).Parse(tmp.ReceiverMsg.Addr)
	return tmp, err
}

type ReceiverAddr struct {
	Province struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"province"`
	City struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"city"`
	Town struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"town"`
	Street struct {
		Name string `json:"name"`
		Id   string `json:"id"`
	} `json:"street"`
	Detail        string `json:"detail"`
	EncryptDetail string `json:"encrypt_detail"`
}

func (m ReceiverAddr) Parse(addr string) (ReceiverAddr, error) {
	var addStruct ReceiverAddr
	err := json.Unmarshal([]byte(addr), &addStruct)
	return addStruct, err
}
