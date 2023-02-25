package spu_getSpuRule_request

import (
	"doudian.com/open/sdk_golang/api/spu_getSpuRule/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuGetSpuRuleRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuGetSpuRuleParam 
}
func (c *SpuGetSpuRuleRequest) GetUrlPath() string{
	return "/spu/getSpuRule"
}


func New() *SpuGetSpuRuleRequest{
	request := &SpuGetSpuRuleRequest{
		Param: &SpuGetSpuRuleParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuGetSpuRuleRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_getSpuRule_response.SpuGetSpuRuleResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_getSpuRule_response.SpuGetSpuRuleResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuGetSpuRuleRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuGetSpuRuleRequest) GetParams() *SpuGetSpuRuleParam{
	return c.Param
}


type SpuGetSpuRuleParam struct {
	// 类目ID
	CategoryId *int64 `json:"category_id,omitempty"`
}
