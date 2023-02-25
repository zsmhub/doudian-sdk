package afterSale_submitEvidence_response

import (
	"doudian.com/open/sdk_golang/core"
)

type AfterSaleSubmitEvidenceResponse struct {
	doudian_sdk.BaseDoudianOpApiResponse
	Data *AfterSaleSubmitEvidenceData `json:"data"`
}
type AfterSaleSubmitEvidenceData struct {
}
