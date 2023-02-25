package product_cancelAudit_request

import (
	"doudian.com/open/sdk_golang/api/product_cancelAudit/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductCancelAuditRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductCancelAuditParam 
}
func (c *ProductCancelAuditRequest) GetUrlPath() string{
	return "/product/cancelAudit"
}


func New() *ProductCancelAuditRequest{
	request := &ProductCancelAuditRequest{
		Param: &ProductCancelAuditParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductCancelAuditRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_cancelAudit_response.ProductCancelAuditResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_cancelAudit_response.ProductCancelAuditResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductCancelAuditRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductCancelAuditRequest) GetParams() *ProductCancelAuditParam{
	return c.Param
}


type ProductCancelAuditParam struct {
	// 商品ID
	ProductId *int64 `json:"product_id,omitempty"`
}
