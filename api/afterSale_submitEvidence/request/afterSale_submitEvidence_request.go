package afterSale_submitEvidence_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_submitEvidence/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleSubmitEvidenceRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleSubmitEvidenceParam 
}
func (c *AfterSaleSubmitEvidenceRequest) GetUrlPath() string{
	return "/afterSale/submitEvidence"
}


func New() *AfterSaleSubmitEvidenceRequest{
	request := &AfterSaleSubmitEvidenceRequest{
		Param: &AfterSaleSubmitEvidenceParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleSubmitEvidenceRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_submitEvidence_response.AfterSaleSubmitEvidenceResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_submitEvidence_response.AfterSaleSubmitEvidenceResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleSubmitEvidenceRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleSubmitEvidenceRequest) GetParams() *AfterSaleSubmitEvidenceParam{
	return c.Param
}


type AfterSaleSubmitEvidenceParam struct {
	// 售后单ID ，通过[/trade/refundListSearch](https://op.jinritemai.com/docs/api-docs/17/254) 或者 [/afterSale/refundProcessDetail](https://op.jinritemai.com/docs/api-docs/17/96) 获取  
	AftersaleId *int64 `json:"aftersale_id,omitempty"`
	// 备注
	Comment *string `json:"comment,omitempty"`
	// [https://xxxx.jpg](https://xxxx.jpg/) | 凭证，最多四张
	Evidence *[]string `json:"evidence,omitempty"`
}
