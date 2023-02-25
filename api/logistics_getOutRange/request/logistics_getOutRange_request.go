package logistics_getOutRange_request

import (
	"doudian.com/open/sdk_golang/api/logistics_getOutRange/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsGetOutRangeRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsGetOutRangeParam 
}
func (c *LogisticsGetOutRangeRequest) GetUrlPath() string{
	return "/logistics/getOutRange"
}


func New() *LogisticsGetOutRangeRequest{
	request := &LogisticsGetOutRangeRequest{
		Param: &LogisticsGetOutRangeParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsGetOutRangeRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_getOutRange_response.LogisticsGetOutRangeResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_getOutRange_response.LogisticsGetOutRangeResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsGetOutRangeRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsGetOutRangeRequest) GetParams() *LogisticsGetOutRangeParam{
	return c.Param
}


type SenderAddress struct {
	// CHN
	CountryCode *string `json:"country_code,omitempty"`
	// 省份名称
	ProvinceName *string `json:"province_name,omitempty"`
	// 城市名称
	CityName *string `json:"city_name,omitempty"`
	// 区名称
	DistrictName *string `json:"district_name,omitempty"`
	// 街道名称
	StreetName string `json:"street_name,omitempty"`
	// 详细地址
	DetailAddress *string `json:"detail_address,omitempty"`
}
type ReceiverAddress struct {
	// CHN
	CountryCode *string `json:"country_code,omitempty"`
	// 省份名称
	ProvinceName *string `json:"province_name,omitempty"`
	// 城市名称
	CityName *string `json:"city_name,omitempty"`
	// 区名称
	DistrictName *string `json:"district_name,omitempty"`
	// 街道名称
	StreetName string `json:"street_name,omitempty"`
	// 详细地址，支持密文
	DetailAddress *string `json:"detail_address,omitempty"`
}
type ServiceListItem struct {
	// code
	ServiceCode string `json:"service_code,omitempty"`
	// value
	ServiceValue string `json:"service_value,omitempty"`
}
type DeliveryReq struct {
	// 是否接受仅镇中心派送  目前仅支持德邦
	IsCenterDelivery bool `json:"is_center_delivery,omitempty"`
	// 是否接受合伙人自提，目前仅支持德邦
	IsPickupSelf bool `json:"is_pickup_self,omitempty"`
}
type LogisticsGetOutRangeParam struct {
	// 快递公司编码
	LogisticsCode *string `json:"logistics_code,omitempty"`
	// 发货地址
	SenderAddress *SenderAddress `json:"sender_address,omitempty"`
	// 收件地址
	ReceiverAddress *ReceiverAddress `json:"receiver_address,omitempty"`
	// 类型（0-揽派合一；1-揽收区域；2-派送区域） 0：取senderAddress, receiverAddress值 1：取senderAddress值 2：取receiverAddress值
	Type *int16 `json:"type,omitempty"`
	// 增值服务 目前只支持德邦
	ServiceList []ServiceListItem `json:"service_list,omitempty"`
	// 产品类型 目前只支持德邦
	ProductType string `json:"product_type,omitempty"`
	// 投递要求 目前只支持德邦
	DeliveryReq *DeliveryReq `json:"delivery_req,omitempty"`
}
