package afterSale_rejectReasonCodeList_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AfterSaleRejectReasonCodeListResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AfterSaleRejectReasonCodeListData `json:"data"`
}
type ItemsItem struct {
	// 售后审核拒绝原因枚举编码
	RejectReasonCode int64 `json:"reject_reason_code"`
	// 售后审核拒绝原因文案
	Reason string `json:"reason"`
	// 凭证描述文案
	EvidenceDescription string `json:"evidence_description"`
	// 是否需要上传凭证，Y必填，N非必填
	EvidenceNeed string `json:"evidence_need"`
	// 凭证示例图片链接
	Image string `json:"image"`
	// 订单类型，即订单信息中order_type   枚举：0-普通实物订单 1-全款预售订单  2-虚拟商品订单 3-快闪店订单 4-电子券  5-三方核销 6-服务市场 -1-通用,不考虑订单类型
	OrderType int64 `json:"order_type"`
	// 是否收到货，0未收到 1收到 -1通用,不考虑是否收到货
	Pkg int64 `json:"pkg"`
}
type AfterSaleRejectReasonCodeListData struct {
	// 售后商家拒绝原因详情列表
	Items []ItemsItem `json:"items"`
}
