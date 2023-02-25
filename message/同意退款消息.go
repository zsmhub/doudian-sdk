package doudian_message

import (
	"encoding/json"
	"time"
)

// 同意退款消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/116

func init() {
	MessageDataParseHandle.RegisterMessage(RefundAgreed{})
}

type RefundAgreed struct {
	AftersaleId      int64     `json:"aftersale_id"`
	AftersaleStatus  int64     `json:"aftersale_status"`
	AftersaleType    int64     `json:"aftersale_type"`
	AgreeTime        int64     `json:"agree_time"`
	PId              int64     `json:"p_id"`
	ReasonCode       int64     `json:"reason_code"`
	RefundAmount     int64     `json:"refund_amount"`
	RefundPostAmount int64     `json:"refund_post_amount"`
	RefundVoucherNum int64     `json:"refund_voucher_num"`
	SId              int64     `json:"s_id"`
	ShopId           int64     `json:"shop_id"`
	UpdateTime       time.Time `json:"update_time"`
}

func (m RefundAgreed) Tag() string {
	return "201"
}

func (m RefundAgreed) Parse(data []byte) (MessageData, error) {
	var tmp RefundAgreed
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
