package doudian_message

import (
	"encoding/json"
	"time"
)

// 同意补寄申请
// 文档：https://op.jinritemai.com/docs/message-docs/31/1984

func init() {
	MessageDataParseHandle.RegisterMessage(RefundAuditAgreeResend{})
}

type RefundAuditAgreeResend struct {
	AfterSaleStatus int64     `json:"after_sale_status"`
	AftersaleId     int64     `json:"aftersale_id"`
	AftersaleType   int64     `json:"aftersale_type"`
	PId             int64     `json:"p_id"`
	ReasonCode      int64     `json:"reason_code"`
	SId             int64     `json:"s_id"`
	ShopId          int64     `json:"shop_id"`
	UpdateTime      time.Time `json:"update_time"`
}

func (m RefundAuditAgreeResend) Tag() string {
	return "227"
}

func (m RefundAuditAgreeResend) Parse(data []byte) (MessageData, error) {
	var tmp RefundAuditAgreeResend
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
