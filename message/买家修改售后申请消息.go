package doudian_message

import (
	"encoding/json"
	"time"
)

// 买家修改售后申请消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/230

func init() {
	MessageDataParseHandle.RegisterMessage(RefundModified{})
}

type RefundModified struct {
	AftersaleId      int64     `json:"aftersale_id"`
	AftersaleStatus  int64     `json:"aftersale_status"`
	AftersaleType    int64     `json:"aftersale_type"`
	ModifyTime       int64     `json:"modify_time"`
	PId              int64     `json:"p_id"`
	ReasonCode       int64     `json:"reason_code"`
	RefundAmount     int64     `json:"refund_amount"`
	RefundPostAmount int64     `json:"refund_post_amount"`
	RefundVoucherNum int64     `json:"refund_voucher_num"`
	SId              int64     `json:"s_id"`
	ShopId           int64     `json:"shop_id"`
	UpdateTime       time.Time `json:"update_time"`
}

func (m RefundModified) Tag() string {
	return "208"
}

func (m RefundModified) Parse(data []byte) (MessageData, error) {
	var tmp RefundModified
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
