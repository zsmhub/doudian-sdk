package brand_list_request

import (
	"doudian.com/open/sdk_golang/api/brand_list/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BrandListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BrandListParam 
}
func (c *BrandListRequest) GetUrlPath() string{
	return "/brand/list"
}


func New() *BrandListRequest{
	request := &BrandListRequest{
		Param: &BrandListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *BrandListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*brand_list_response.BrandListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &brand_list_response.BrandListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *BrandListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *BrandListRequest) GetParams() *BrandListParam{
	return c.Param
}


type BrandListParam struct {
	// （已停止使用）类目列表
	Categories []int64 `json:"categories,omitempty"`
	// （已停止使用）起始位
	Offset int64 `json:"offset,omitempty"`
	// （已停止使用）单次最大条数
	Size int64 `json:"size,omitempty"`
	// （已停止使用）排序顺序，默认为倒排 0:降序, 1:升序
	Sort int32 `json:"sort,omitempty"`
	// （已停止使用）品牌状态 1:审核中, 2:审核通过, 3:审核拒绝
	Status int32 `json:"status,omitempty"`
	// （已停止使用）是否返回完全的品牌信息
	FullBrandInfo bool `json:"full_brand_info,omitempty"`
	// （推荐使用，必填）类目id
	CategoryId int64 `json:"category_id,omitempty"`
	// 品牌前缀（中文或者英文），适用于不需要品牌资质的场景，根据前缀搜索品牌
	Query string `json:"query,omitempty"`
}
