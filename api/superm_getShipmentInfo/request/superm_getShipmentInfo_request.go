package superm_getShipmentInfo_request

import (
	"doudian.com/open/sdk_golang/api/superm_getShipmentInfo/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermGetShipmentInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermGetShipmentInfoParam 
}
func (c *SupermGetShipmentInfoRequest) GetUrlPath() string{
	return "/superm/getShipmentInfo"
}


func New() *SupermGetShipmentInfoRequest{
	request := &SupermGetShipmentInfoRequest{
		Param: &SupermGetShipmentInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermGetShipmentInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_getShipmentInfo_response.SupermGetShipmentInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_getShipmentInfo_response.SupermGetShipmentInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermGetShipmentInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermGetShipmentInfoRequest) GetParams() *SupermGetShipmentInfoParam{
	return c.Param
}


type SupermGetShipmentInfoParam struct {
	// 店铺父订单号
	ShopOrderId *int64 `json:"shop_order_id,omitempty"`
	// 上门取运力对应售后单号
	AftersaleId int64 `json:"aftersale_id,omitempty"`
	// 配送类型，1 表示发货单，需要传入店铺订单号 2 表示上门取件，需要传入店铺订单号和售后单编号
	ShipmentType *int32 `json:"shipment_type,omitempty"`
}
