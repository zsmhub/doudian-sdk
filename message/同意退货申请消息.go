package doudian_message

import (
	"encoding/json"
	"time"
)

// 同意退货申请消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/117

func init() {
	MessageDataParseHandle.RegisterMessage(RefundReturnApplyAgreed{})
}

type RefundReturnApplyAgreed struct {
	AftersaleId      int64     `json:"aftersale_id"`
	AftersaleStatus  int64     `json:"aftersale_status"`
	AftersaleType    int64     `json:"aftersale_type"`
	PId              int64     `json:"p_id"`
	ReasonCode       int64     `json:"reason_code"`
	RefundAmount     int64     `json:"refund_amount"`
	RefundPostAmount int64     `json:"refund_post_amount"`
	RefundVoucherNum int64     `json:"refund_voucher_num"`
	RefuseTime       int64     `json:"refuse_time"`
	SId              int64     `json:"s_id"`
	ShopId           int64     `json:"shop_id"`
	UpdateTime       time.Time `json:"update_time"`
}

func (m RefundReturnApplyAgreed) Tag() string {
	return "202"
}

func (m RefundReturnApplyAgreed) Parse(data []byte) (MessageData, error) {
	var tmp RefundReturnApplyAgreed
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
