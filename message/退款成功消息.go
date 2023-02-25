package doudian_message

import (
	"encoding/json"
	"time"
)

// 退款成功消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/121

func init() {
	MessageDataParseHandle.RegisterMessage(RefundSuccess{})
}

type RefundSuccess struct {
	AftersaleId      int64     `json:"aftersale_id"`
	AftersaleStatus  int64     `json:"aftersale_status"`
	AftersaleType    int64     `json:"aftersale_type"`
	PId              int64     `json:"p_id"`
	ReasonCode       int64     `json:"reason_code"`
	RefundAmount     int64     `json:"refund_amount"`
	RefundPostAmount int64     `json:"refund_post_amount"`
	RefundVoucherNum int64     `json:"refund_voucher_num"`
	SId              int64     `json:"s_id"`
	ShopId           int64     `json:"shop_id"`
	SuccessTime      int64     `json:"success_time"`
	UpdateTime       time.Time `json:"update_time"`
}

func (m RefundSuccess) Tag() string {
	return "206"
}

func (m RefundSuccess) Parse(data []byte) (MessageData, error) {
	var tmp RefundSuccess
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
