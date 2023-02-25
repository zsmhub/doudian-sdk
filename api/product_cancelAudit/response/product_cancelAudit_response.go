package product_cancelAudit_response

import (
	"doudian.com/open/sdk_golang/core"
)

type ProductCancelAuditResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *ProductCancelAuditData `json:"data"`
}
type ProductCancelAuditData struct {
}
