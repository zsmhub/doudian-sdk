package product_listV2_request

import (
	"doudian.com/open/sdk_golang/api/product_listV2/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductListV2Request struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductListV2Param 
}
func (c *ProductListV2Request) GetUrlPath() string{
	return "/product/listV2"
}


func New() *ProductListV2Request{
	request := &ProductListV2Request{
		Param: &ProductListV2Param{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductListV2Request) Execute(accessToken *doudian_sdk.AccessToken) (*product_listV2_response.ProductListV2Response, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_listV2_response.ProductListV2Response{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductListV2Request) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductListV2Request) GetParams() *ProductListV2Param{
	return c.Param
}


type ProductListV2Param struct {
	// 商品在店铺中状态: 0-在线；1-下线；2-删除；详见商品状态机：https://op.jinritemai.com/docs/question-docs/92/2070
	Status int64 `json:"status,omitempty"`
	// 商品审核状态: 1-未提交；2-待审核；3-审核通过；4-审核未通过；5-封禁；7-审核通过待上架；详见商品状态机：https://op.jinritemai.com/docs/question-docs/92/2070
	CheckStatus int64 `json:"check_status,omitempty"`
	// 商品类型；0-普通；1-新客商品；3-虚拟；6-玉石闪购；7-云闪购 ；127-其他类型；
	ProductType int64 `json:"product_type,omitempty"`
	// 商品创建开始时间，unix时间戳，单位：秒；
	StartTime int64 `json:"start_time,omitempty"`
	// 商品创建结束时间，unix时间戳，单位：秒；
	EndTime int64 `json:"end_time,omitempty"`
	// 商品更新开始时间，unix时间戳，单位：秒；注意：查询范围是update_start_time和update_end_time之间的数据，不包含入参时间。
	UpdateStartTime int64 `json:"update_start_time,omitempty"`
	// 商品更新结束时间，unix时间戳，单位：秒；注意：查询范围是update_start_time和update_end_time之间的数据，不包含入参时间。
	UpdateEndTime int64 `json:"update_end_time,omitempty"`
	// 页码，从1开始，最大为100；page* size最大不能超过1万条
	Page *int64 `json:"page,omitempty"`
	// 页数，填写范围1~100；page* size最大不能超过1万条
	Size *int64 `json:"size,omitempty"`
	// 小时达商家使用的门店id
	StoreId int64 `json:"store_id,omitempty"`
	// 商品标题，支持模糊匹配
	Name string `json:"name,omitempty"`
	// 商品id，最大支持传入100个；
	ProductId []string `json:"product_id,omitempty"`
}
