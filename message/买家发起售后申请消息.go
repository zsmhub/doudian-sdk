package doudian_message

import (
	"encoding/json"
	"time"
)

// 买家发起售后申请消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/115

func init() {
	MessageDataParseHandle.RegisterMessage(RefundCreated{})
}

type RefundCreated struct {
	AftersaleId      int64     `json:"aftersale_id"`
	AftersaleStatus  int64     `json:"aftersale_status"`
	AftersaleType    int64     `json:"aftersale_type"`
	ApplyTime        int64     `json:"apply_time"`
	PId              int64     `json:"p_id"`
	ReasonCode       int64     `json:"reason_code"`
	RefundAmount     int64     `json:"refund_amount"`
	RefundPostAmount int64     `json:"refund_post_amount"`
	RefundVoucherNum int64     `json:"refund_voucher_num"`
	SId              int64     `json:"s_id"`
	ShopId           int64     `json:"shop_id"`
	UpdateTime       time.Time `json:"update_time"`
}

func (m RefundCreated) Tag() string {
	return "200"
}

func (m RefundCreated) Parse(data []byte) (MessageData, error) {
	var tmp RefundCreated
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
