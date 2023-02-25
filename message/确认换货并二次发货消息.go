package doudian_message

import (
	"encoding/json"
	"time"
)

// 确认换货并二次发货消息
// 文档：https://op.jinritemai.com/docs/message-docs/31/240

func init() {
	MessageDataParseHandle.RegisterMessage(RefundExchangeComfirmed{})
}

type RefundExchangeComfirmed struct {
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

func (m RefundExchangeComfirmed) Tag() string {
	return "211"
}

func (m RefundExchangeComfirmed) Parse(data []byte) (MessageData, error) {
	var tmp RefundExchangeComfirmed
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
