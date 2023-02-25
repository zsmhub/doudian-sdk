package security_batchReportOrderSecurityEvent_response

import (
	"doudian.com/open/sdk_golang/core"
)

type SecurityBatchReportOrderSecurityEventResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *SecurityBatchReportOrderSecurityEventData `json:"data"`
}
type CustomError struct {
	// 错误码
	ErrNo int64 `json:"ErrNo"`
	// 错误原因
	ErrMsg string `json:"ErrMsg"`
}
type SecurityBatchReportOrderSecurityEventData struct {
	// 错误信息
	CustomError *CustomError `json:"CustomError"`
}
