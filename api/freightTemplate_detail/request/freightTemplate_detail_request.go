package freightTemplate_detail_request

import (
	"doudian.com/open/sdk_golang/api/freightTemplate_detail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type FreightTemplateDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *FreightTemplateDetailParam 
}
func (c *FreightTemplateDetailRequest) GetUrlPath() string{
	return "/freightTemplate/detail"
}


func New() *FreightTemplateDetailRequest{
	request := &FreightTemplateDetailRequest{
		Param: &FreightTemplateDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *FreightTemplateDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*freightTemplate_detail_response.FreightTemplateDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &freightTemplate_detail_response.FreightTemplateDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *FreightTemplateDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *FreightTemplateDetailRequest) GetParams() *FreightTemplateDetailParam{
	return c.Param
}


type FreightTemplateDetailParam struct {
	// 模板id
	FreightId *int64 `json:"freight_id,omitempty"`
}
