package afterSale_addOrderRemark_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_addOrderRemark/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleAddOrderRemarkRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleAddOrderRemarkParam 
}
func (c *AfterSaleAddOrderRemarkRequest) GetUrlPath() string{
	return "/afterSale/addOrderRemark"
}


func New() *AfterSaleAddOrderRemarkRequest{
	request := &AfterSaleAddOrderRemarkRequest{
		Param: &AfterSaleAddOrderRemarkParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleAddOrderRemarkRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_addOrderRemark_response.AfterSaleAddOrderRemarkResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_addOrderRemark_response.AfterSaleAddOrderRemarkResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleAddOrderRemarkRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleAddOrderRemarkRequest) GetParams() *AfterSaleAddOrderRemarkParam{
	return c.Param
}


type AfterSaleAddOrderRemarkParam struct {
	// 订单ID，与售后单ID二选一传入
	OrderId string `json:"order_id,omitempty"`
	// 售后单ID，与订单ID二选一传入，二者均传入时售后单ID的优先级更高
	AfterSaleId string `json:"after_sale_id,omitempty"`
	// 商家添加的备注信息
	Remark *string `json:"remark,omitempty"`
}
