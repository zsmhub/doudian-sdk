package freightTemplate_list_request

import (
	"doudian.com/open/sdk_golang/api/freightTemplate_list/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type FreightTemplateListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *FreightTemplateListParam 
}
func (c *FreightTemplateListRequest) GetUrlPath() string{
	return "/freightTemplate/list"
}


func New() *FreightTemplateListRequest{
	request := &FreightTemplateListRequest{
		Param: &FreightTemplateListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *FreightTemplateListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*freightTemplate_list_response.FreightTemplateListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &freightTemplate_list_response.FreightTemplateListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *FreightTemplateListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *FreightTemplateListRequest) GetParams() *FreightTemplateListParam{
	return c.Param
}


type FreightTemplateListParam struct {
	// 运费模板名称，支持模糊搜索
	Name string `json:"name,omitempty"`
	// 页数（默认为0，第一页从0开始）
	Page string `json:"page,omitempty"`
	// 每页模板数（默认为10）
	Size string `json:"size,omitempty"`
}
