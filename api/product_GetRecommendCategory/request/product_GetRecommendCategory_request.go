package product_GetRecommendCategory_request

import (
	"doudian.com/open/sdk_golang/api/product_GetRecommendCategory/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductGetRecommendCategoryRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductGetRecommendCategoryParam 
}
func (c *ProductGetRecommendCategoryRequest) GetUrlPath() string{
	return "/product/GetRecommendCategory"
}


func New() *ProductGetRecommendCategoryRequest{
	request := &ProductGetRecommendCategoryRequest{
		Param: &ProductGetRecommendCategoryParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductGetRecommendCategoryRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_GetRecommendCategory_response.ProductGetRecommendCategoryResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_GetRecommendCategory_response.ProductGetRecommendCategoryResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductGetRecommendCategoryRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductGetRecommendCategoryRequest) GetParams() *ProductGetRecommendCategoryParam{
	return c.Param
}


type PicItem struct {
	// 图片链接，必须是素材中心的url
	Url *string `json:"url,omitempty"`
}
type ProductFormatNewItem struct {
	// 属性id
	Value int64 `json:"value,omitempty"`
	// 属性名称
	Name string `json:"name,omitempty"`
}
type ProductGetRecommendCategoryParam struct {
	// category_infer: 基于标题、图片等推断商品类目；product_info: 表示基于商品内容进行类目错放判断，需要传入商品类目、属性等；smart_publish: 表示图片预测类目，需要传入商品主图；
	Scene *string `json:"scene,omitempty"`
	// 商品主图图片url，scene为smart_publish时必传
	Pic []PicItem `json:"pic,omitempty"`
	// 商品类目id，scene为product_info时必传
	CategoryLeafId int64 `json:"category_leaf_id,omitempty"`
	// 商品标题，scene为category_infer时必填; sense为product_info时选填
	Name string `json:"name,omitempty"`
	// 商品类目属性
	ProductFormatNew map[int64]ProductFormatNewItem `json:"product_format_new,omitempty"`
	// 品牌id
	StandardBrandId int64 `json:"standard_brand_id,omitempty"`
}
