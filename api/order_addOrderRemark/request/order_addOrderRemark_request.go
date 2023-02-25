package order_addOrderRemark_request

import (
	"doudian.com/open/sdk_golang/api/order_addOrderRemark/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderAddOrderRemarkRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderAddOrderRemarkParam 
}
func (c *OrderAddOrderRemarkRequest) GetUrlPath() string{
	return "/order/addOrderRemark"
}


func New() *OrderAddOrderRemarkRequest{
	request := &OrderAddOrderRemarkRequest{
		Param: &OrderAddOrderRemarkParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderAddOrderRemarkRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_addOrderRemark_response.OrderAddOrderRemarkResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_addOrderRemark_response.OrderAddOrderRemarkResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderAddOrderRemarkRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderAddOrderRemarkRequest) GetParams() *OrderAddOrderRemarkParam{
	return c.Param
}


type OrderAddOrderRemarkParam struct {
	// 父订单id
	OrderId *string `json:"order_id,omitempty"`
	// 备注内容，最大不得超过60个字符
	Remark *string `json:"remark,omitempty"`
	// 是否加旗标，不填则默认为否 true：需要加旗标 false：不加旗标
	IsAddStar string `json:"is_add_star,omitempty"`
	// 标星等级，范围0～5 0为灰色旗标，5为红色旗标，数字越大颜色越深 0灰 1紫 2青 3绿 4橙 5红
	Star string `json:"star,omitempty"`
}
