package product_getRecommendName_request

import (
	"doudian.com/open/sdk_golang/api/product_getRecommendName/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductGetRecommendNameRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductGetRecommendNameParam 
}
func (c *ProductGetRecommendNameRequest) GetUrlPath() string{
	return "/product/getRecommendName"
}


func New() *ProductGetRecommendNameRequest{
	request := &ProductGetRecommendNameRequest{
		Param: &ProductGetRecommendNameParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductGetRecommendNameRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_getRecommendName_response.ProductGetRecommendNameResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_getRecommendName_response.ProductGetRecommendNameResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductGetRecommendNameRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductGetRecommendNameRequest) GetParams() *ProductGetRecommendNameParam{
	return c.Param
}


type SelectPropertyItem struct {
	// 属性id
	Value *int64 `json:"value,omitempty"`
	// 属性名称
	Name *string `json:"name,omitempty"`
}
type ProductGetRecommendNameParam struct {
	// 暂时仅开放一种场景：1. product_name_prefix，表示基于 命中的商品类目属性 推荐 商品标题前缀
	Scene *[]string `json:"scene,omitempty"`
	// 商品叶子类目id，请先使用【/product/getProductUpdateRule】接口查询类目id查看recommend_name_rule. satisfy_prefix是否=true
	CategoryLeafId *int64 `json:"category_leaf_id,omitempty"`
	// 一级类目id，scene为product_name_prefix时必传，没有时传0
	FirstCid *int64 `json:"first_cid,omitempty"`
	// 一级类目名，scene为product_name_prefix时必传
	FirstCidName *string `json:"first_cid_name,omitempty"`
	// 二级类目id，scene为product_name_prefix时必传，没有时传0
	SecondCid *int64 `json:"second_cid,omitempty"`
	// 二级类目名，scene为product_name_prefix时必传
	SecondCidName *string `json:"second_cid_name,omitempty"`
	// 三级类目id，scene为product_name_prefix时必传，没有时传0
	ThirdCid *int64 `json:"third_cid,omitempty"`
	// 三级类目名，scene为product_name_prefix时必传
	ThirdCidName *string `json:"third_cid_name,omitempty"`
	// 商品类目属性，参考"select_property":{"1687":[{"value":0,"name":"填入品牌名"}],"3320":[{"value":18972,"name":"99新"}]}
	SelectProperty *map[int64][]SelectPropertyItem `json:"select_property,omitempty"`
}
