package spu_addShopSpu_request

import (
	"doudian.com/open/sdk_golang/api/spu_addShopSpu/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuAddShopSpuRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuAddShopSpuParam 
}
func (c *SpuAddShopSpuRequest) GetUrlPath() string{
	return "/spu/addShopSpu"
}


func New() *SpuAddShopSpuRequest{
	request := &SpuAddShopSpuRequest{
		Param: &SpuAddShopSpuParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuAddShopSpuRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_addShopSpu_response.SpuAddShopSpuResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_addShopSpu_response.SpuAddShopSpuResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuAddShopSpuRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuAddShopSpuRequest) GetParams() *SpuAddShopSpuParam{
	return c.Param
}


type PropertyInfosItem struct {
	// 属性id
	PropertyId *int64 `json:"property_id,omitempty"`
	// 属性值id,可输入类型的属性valueId为0
	ValueId *int64 `json:"value_id,omitempty"`
	// 父属性值id
	ParentValueId int64 `json:"parent_value_id,omitempty"`
	// 属性层级
	Level int64 `json:"level,omitempty"`
	// 属性值名称，非可输入类型时不用传
	ValueName string `json:"value_name,omitempty"`
}
type SpuAddShopSpuParam struct {
	// spuName
	SpuName *string `json:"spu_name,omitempty"`
	// 产品唯一标识，没有填写空字符串
	UpcCode string `json:"upc_code,omitempty"`
	// 类目id
	CategoryId *int64 `json:"category_id,omitempty"`
	// 品牌id
	BrandId *int64 `json:"brand_id,omitempty"`
	// 属性信息
	PropertyInfos *[]PropertyInfosItem `json:"property_infos,omitempty"`
}
