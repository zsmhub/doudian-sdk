package spu_getSpu_request

import (
	"doudian.com/open/sdk_golang/api/spu_getSpu/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuGetSpuRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuGetSpuParam 
}
func (c *SpuGetSpuRequest) GetUrlPath() string{
	return "/spu/getSpu"
}


func New() *SpuGetSpuRequest{
	request := &SpuGetSpuRequest{
		Param: &SpuGetSpuParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuGetSpuRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_getSpu_response.SpuGetSpuResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_getSpu_response.SpuGetSpuResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuGetSpuRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuGetSpuRequest) GetParams() *SpuGetSpuParam{
	return c.Param
}


type PropertyValuesItem struct {
	// 属性值ID
	ValueId int64 `json:"value_id,omitempty"`
	// 属性值
	ValueName string `json:"value_name,omitempty"`
}
type KeyPropertiesItem struct {
	// 属性ID
	PropertyId int64 `json:"property_id,omitempty"`
	// 属性值
	PropertyValues []PropertyValuesItem `json:"property_values,omitempty"`
}
type SpuGetSpuParam struct {
	// 叶子类目ID
	CategoryLeafId int64 `json:"category_leaf_id,omitempty"`
	// 关键属性
	KeyProperties []KeyPropertiesItem `json:"key_properties,omitempty"`
	// SPU ID（传SPUID时，关键属性不用传，传关键属性时，SPU ID不用传，两个都传的情况下，会以SPU ID为准，查询SPU）
	SpuId int64 `json:"spu_id,omitempty"`
}
