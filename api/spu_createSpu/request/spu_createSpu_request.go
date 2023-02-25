package spu_createSpu_request

import (
	"doudian.com/open/sdk_golang/api/spu_createSpu/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuCreateSpuRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuCreateSpuParam 
}
func (c *SpuCreateSpuRequest) GetUrlPath() string{
	return "/spu/createSpu"
}


func New() *SpuCreateSpuRequest{
	request := &SpuCreateSpuRequest{
		Param: &SpuCreateSpuParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuCreateSpuRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_createSpu_response.SpuCreateSpuResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_createSpu_response.SpuCreateSpuResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuCreateSpuRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuCreateSpuRequest) GetParams() *SpuCreateSpuParam{
	return c.Param
}


type SpuCreateSpuParam struct {
	// 叶子类目ID
	CategoryLeafId *int64 `json:"category_leaf_id,omitempty"`
	// SPU图片（/spu/batchUploadImg返回的URL）
	SpuImages []string `json:"spu_images,omitempty"`
	// 属性信息（/spu/getSpuRule接口可获取）
	PropertyInfos *[]PropertyInfosItem `json:"property_infos,omitempty"`
	// SPU实物图，用于证明SPU存在，比如包装图，版权页图（/spu/batchUploadImg返回的URL）
	SpuActualImages []string `json:"spu_actual_images,omitempty"`
}
type ValuesItem struct {
	// 属性值ID
	ValueId *int64 `json:"value_id,omitempty"`
	// 属性值
	ValueName *string `json:"value_name,omitempty"`
}
type PropertyInfosItem struct {
	// 属性ID
	PropertyId *int64 `json:"property_id,omitempty"`
	// 属性值
	Values *[]ValuesItem `json:"values,omitempty"`
}
