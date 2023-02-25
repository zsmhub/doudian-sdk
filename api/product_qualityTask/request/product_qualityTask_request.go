package product_qualityTask_request

import (
	"doudian.com/open/sdk_golang/api/product_qualityTask/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductQualityTaskRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductQualityTaskParam 
}
func (c *ProductQualityTaskRequest) GetUrlPath() string{
	return "/product/qualityTask"
}


func New() *ProductQualityTaskRequest{
	request := &ProductQualityTaskRequest{
		Param: &ProductQualityTaskParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductQualityTaskRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_qualityTask_response.ProductQualityTaskResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_qualityTask_response.ProductQualityTaskResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductQualityTaskRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductQualityTaskRequest) GetParams() *ProductQualityTaskParam{
	return c.Param
}


type ProductQualityTaskParam struct {
	// 是否只返回简要信息，不写默认false
	BriefOnly bool `json:"brief_only,omitempty"`
}
