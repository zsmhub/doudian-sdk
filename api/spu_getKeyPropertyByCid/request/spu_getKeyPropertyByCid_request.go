package spu_getKeyPropertyByCid_request

import (
	"doudian.com/open/sdk_golang/api/spu_getKeyPropertyByCid/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuGetKeyPropertyByCidRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuGetKeyPropertyByCidParam 
}
func (c *SpuGetKeyPropertyByCidRequest) GetUrlPath() string{
	return "/spu/getKeyPropertyByCid"
}


func New() *SpuGetKeyPropertyByCidRequest{
	request := &SpuGetKeyPropertyByCidRequest{
		Param: &SpuGetKeyPropertyByCidParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuGetKeyPropertyByCidRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_getKeyPropertyByCid_response.SpuGetKeyPropertyByCidResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_getKeyPropertyByCid_response.SpuGetKeyPropertyByCidResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuGetKeyPropertyByCidRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuGetKeyPropertyByCidRequest) GetParams() *SpuGetKeyPropertyByCidParam{
	return c.Param
}


type SpuGetKeyPropertyByCidParam struct {
	// 类目id
	CategoryId *int64 `json:"category_id,omitempty"`
	// 页码
	Page *int64 `json:"page,omitempty"`
	// 每页大小
	Size *int64 `json:"size,omitempty"`
}
