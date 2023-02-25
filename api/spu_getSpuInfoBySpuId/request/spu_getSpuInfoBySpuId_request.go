package spu_getSpuInfoBySpuId_request

import (
	"doudian.com/open/sdk_golang/api/spu_getSpuInfoBySpuId/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuGetSpuInfoBySpuIdRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuGetSpuInfoBySpuIdParam 
}
func (c *SpuGetSpuInfoBySpuIdRequest) GetUrlPath() string{
	return "/spu/getSpuInfoBySpuId"
}


func New() *SpuGetSpuInfoBySpuIdRequest{
	request := &SpuGetSpuInfoBySpuIdRequest{
		Param: &SpuGetSpuInfoBySpuIdParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuGetSpuInfoBySpuIdRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_getSpuInfoBySpuId_response.SpuGetSpuInfoBySpuIdResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_getSpuInfoBySpuId_response.SpuGetSpuInfoBySpuIdResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuGetSpuInfoBySpuIdRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuGetSpuInfoBySpuIdRequest) GetParams() *SpuGetSpuInfoBySpuIdParam{
	return c.Param
}


type SpuGetSpuInfoBySpuIdParam struct {
	// spuId
	SpuId *int64 `json:"spu_id,omitempty"`
}
