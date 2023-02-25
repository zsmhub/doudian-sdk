package spu_getSpuTpl_request

import (
	"doudian.com/open/sdk_golang/api/spu_getSpuTpl/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuGetSpuTplRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuGetSpuTplParam 
}
func (c *SpuGetSpuTplRequest) GetUrlPath() string{
	return "/spu/getSpuTpl"
}


func New() *SpuGetSpuTplRequest{
	request := &SpuGetSpuTplRequest{
		Param: &SpuGetSpuTplParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuGetSpuTplRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_getSpuTpl_response.SpuGetSpuTplResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_getSpuTpl_response.SpuGetSpuTplResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuGetSpuTplRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuGetSpuTplRequest) GetParams() *SpuGetSpuTplParam{
	return c.Param
}


type SpuGetSpuTplParam struct {
	// 类目id
	CategoryId *int64 `json:"category_id,omitempty"`
}
