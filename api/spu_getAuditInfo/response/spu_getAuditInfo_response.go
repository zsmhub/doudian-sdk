package spu_getAuditInfo_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SpuGetAuditInfoResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SpuGetAuditInfoData `json:"data"`
}
type SpuGetAuditInfoData struct {
	// SPU编号
	SpuId int64 `json:"spu_id"`
	// SPU状态，1:已上线，2:已下线，3:审核中，4:审核不通过
	SpuStatus int64 `json:"spu_status"`
	// 审核驳回原因
	RejectReason string `json:"reject_reason"`
	// 审核记录创建时间
	CreateTime string `json:"create_time"`
	// 审核记录更新时间
	UpdateTime string `json:"update_time"`
}
