package address_list_request

import (
	"doudian.com/open/sdk_golang/api/address_list/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AddressListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AddressListParam 
}
func (c *AddressListRequest) GetUrlPath() string{
	return "/address/list"
}


func New() *AddressListRequest{
	request := &AddressListRequest{
		Param: &AddressListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AddressListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*address_list_response.AddressListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &address_list_response.AddressListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AddressListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AddressListRequest) GetParams() *AddressListParam{
	return c.Param
}


type AddressListParam struct {
	// 【已废弃】店铺id
	ShopId int64 `json:"shop_id,omitempty"`
	// 翻页每页数量，默认10
	PageSize *int64 `json:"page_size,omitempty"`
	// 翻页页数，从1开始
	PageNo *int64 `json:"page_no,omitempty"`
	// 排序方式支持asc/desc
	OrderBy *string `json:"order_by,omitempty"`
	// 排序字段；create_time-创建时间排序，update_time-更新时间排序；
	OrderField *string `json:"order_field,omitempty"`
	// 门店ID，抖超小时达业务使用传入后只获取该门店下地址列表；
	StoreId int64 `json:"store_id,omitempty"`
}
