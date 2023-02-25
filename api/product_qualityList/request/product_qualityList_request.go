package product_qualityList_request

import (
	"doudian.com/open/sdk_golang/api/product_qualityList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductQualityListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductQualityListParam 
}
func (c *ProductQualityListRequest) GetUrlPath() string{
	return "/product/qualityList"
}


func New() *ProductQualityListRequest{
	request := &ProductQualityListRequest{
		Param: &ProductQualityListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductQualityListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_qualityList_response.ProductQualityListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_qualityList_response.ProductQualityListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductQualityListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductQualityListRequest) GetParams() *ProductQualityListParam{
	return c.Param
}


type ProductQualityListParam struct {
	// 商品ID
	ProductId int64 `json:"product_id,omitempty"`
	// 商品名字
	ProductName string `json:"product_name,omitempty"`
	// 排序方式，降序“desc”，升序“asc”， 不传默认降序
	OrderBy string `json:"order_by,omitempty"`
	// 分页参数，页数从1开始
	Page *int64 `json:"page,omitempty"`
	// 分页参数，每页大小，填写范围1~100；
	PageSize *int64 `json:"page_size,omitempty"`
	// 【已废弃】开发者可以传入：9999999999
	TaskId *int64 `json:"task_id,omitempty"`
	// 商品诊断状态，0-待诊断；1-待优化；2-已修改审核中；3-已优化；以前的status 字段因为配置错误已经删除，请使用此新字段
	DiagnoseStatus []int64 `json:"diagnose_status,omitempty"`
}
