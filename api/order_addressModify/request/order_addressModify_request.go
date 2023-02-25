package order_addressModify_request

import (
	"doudian.com/open/sdk_golang/api/order_addressModify/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderAddressModifyRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderAddressModifyParam 
}
func (c *OrderAddressModifyRequest) GetUrlPath() string{
	return "/order/addressModify"
}


func New() *OrderAddressModifyRequest{
	request := &OrderAddressModifyRequest{
		Param: &OrderAddressModifyParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderAddressModifyRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_addressModify_response.OrderAddressModifyResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_addressModify_response.OrderAddressModifyResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderAddressModifyRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderAddressModifyRequest) GetParams() *OrderAddressModifyParam{
	return c.Param
}


type Province struct {
	// 省id
	Id string `json:"id,omitempty"`
	// 省名称
	Name string `json:"name,omitempty"`
}
type City struct {
	// 城市id
	Id string `json:"id,omitempty"`
	// 城市名称
	Name string `json:"name,omitempty"`
}
type Town struct {
	// 镇id
	Id string `json:"id,omitempty"`
	// 镇名称
	Name string `json:"name,omitempty"`
}
type Street struct {
	// 街道id
	Id string `json:"id,omitempty"`
	// 街道名称
	Name string `json:"name,omitempty"`
}
type PostAddr struct {
	// 省
	Province *Province `json:"province,omitempty"`
	// 城市
	City *City `json:"city,omitempty"`
	// 镇
	Town *Town `json:"town,omitempty"`
	// 街道
	Street *Street `json:"street,omitempty"`
	// 详细地址
	AddressDetail string `json:"address_detail,omitempty"`
	// 暂时未使用的字段
	LandMark string `json:"land_mark,omitempty"`
}
type OrderAddressModifyParam struct {
	// 订单号
	OrderId *string `json:"order_id,omitempty"`
	// 修改收货地址，post_addr中的所有字段都必传。
	PostAddr *PostAddr `json:"post_addr,omitempty"`
	// 收货人姓名
	PostReceiver string `json:"post_receiver,omitempty"`
	// 收货人电话号码
	PostTel string `json:"post_tel,omitempty"`
}
