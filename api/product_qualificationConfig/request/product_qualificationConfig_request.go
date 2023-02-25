package product_qualificationConfig_request

import (
	"doudian.com/open/sdk_golang/api/product_qualificationConfig/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductQualificationConfigRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductQualificationConfigParam 
}
func (c *ProductQualificationConfigRequest) GetUrlPath() string{
	return "/product/qualificationConfig"
}


func New() *ProductQualificationConfigRequest{
	request := &ProductQualificationConfigRequest{
		Param: &ProductQualificationConfigParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductQualificationConfigRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_qualificationConfig_response.ProductQualificationConfigResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_qualificationConfig_response.ProductQualificationConfigResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductQualificationConfigRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductQualificationConfigRequest) GetParams() *ProductQualificationConfigParam{
	return c.Param
}


type ProductQualificationConfigParam struct {
	// 类目ID
	CategoryId *int64 `json:"category_id,omitempty"`
}
