package order_downloadShopAccountItem_request

import (
	"doudian.com/open/sdk_golang/api/order_downloadShopAccountItem/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderDownloadShopAccountItemRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderDownloadShopAccountItemParam 
}
func (c *OrderDownloadShopAccountItemRequest) GetUrlPath() string{
	return "/order/downloadShopAccountItem"
}


func New() *OrderDownloadShopAccountItemRequest{
	request := &OrderDownloadShopAccountItemRequest{
		Param: &OrderDownloadShopAccountItemParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderDownloadShopAccountItemRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_downloadShopAccountItem_response.OrderDownloadShopAccountItemResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_downloadShopAccountItem_response.OrderDownloadShopAccountItemResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderDownloadShopAccountItemRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderDownloadShopAccountItemRequest) GetParams() *OrderDownloadShopAccountItemParam{
	return c.Param
}


type OrderDownloadShopAccountItemParam struct {
	// 开始时间
	StartTime string `json:"start_time,omitempty"`
	// 结束时间
	EndTime string `json:"end_time,omitempty"`
	// 订单号(如果订单号未传，则时间必须传)
	OrderId string `json:"order_id,omitempty"`
	// 动账账户 0: 所有 1: 微信 2:支付宝 3:合众支付 4:聚合支付
	AccountType int32 `json:"account_type,omitempty"`
	// 计费类型 0:全部 1:鲁班广告 2:精选联盟 3:值点商城 4:小店自卖 5:橙子建站 6:POI 7:抖+ 8:穿山甲 9:服务市场 10:服务市场外包客服 11:学浪
	BizType int32 `json:"biz_type,omitempty"`
	// 时间类型 0 动账时间 1 下单时间
	TimeType int32 `json:"time_type,omitempty"`
	// 产品id
	ProductId string `json:"product_id,omitempty"`
}
