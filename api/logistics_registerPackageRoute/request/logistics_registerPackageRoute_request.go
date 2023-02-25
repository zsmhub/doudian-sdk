package logistics_registerPackageRoute_request

import (
	"doudian.com/open/sdk_golang/api/logistics_registerPackageRoute/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsRegisterPackageRouteRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsRegisterPackageRouteParam 
}
func (c *LogisticsRegisterPackageRouteRequest) GetUrlPath() string{
	return "/logistics/registerPackageRoute"
}


func New() *LogisticsRegisterPackageRouteRequest{
	request := &LogisticsRegisterPackageRouteRequest{
		Param: &LogisticsRegisterPackageRouteParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsRegisterPackageRouteRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_registerPackageRoute_response.LogisticsRegisterPackageRouteResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_registerPackageRoute_response.LogisticsRegisterPackageRouteResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsRegisterPackageRouteRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsRegisterPackageRouteRequest) GetParams() *LogisticsRegisterPackageRouteParam{
	return c.Param
}


type Contact struct {
	// 姓名
	Name *string `json:"name,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
}
type Town struct {
	// 区
	Name *string `json:"name,omitempty"`
	// code
	Code string `json:"code,omitempty"`
}
type Street struct {
	// 街道
	Name *string `json:"name,omitempty"`
	// code
	Code string `json:"code,omitempty"`
}
type CargoListItem struct {
	// 名称
	Name string `json:"name,omitempty"`
	// 质量
	Quantity int32 `json:"quantity,omitempty"`
	// 体积
	Volume int32 `json:"volume,omitempty"`
	// 总重
	TotalWeight *int32 `json:"total_weight,omitempty"`
	// 净重
	TotalNetWeight *int32 `json:"total_net_weight,omitempty"`
	// 单位
	Unit string `json:"unit,omitempty"`
}
type LogisticsRegisterPackageRouteParam struct {
	// 物流商
	Express *string `json:"express,omitempty"`
	// 运单号
	TrackNo *string `json:"track_no,omitempty"`
	// 外部单据id
	OuterOrderId string `json:"outer_order_id,omitempty"`
	// 外部子单据id
	OuterSubOrderId string `json:"outer_sub_order_id,omitempty"`
	// 回调url
	CallbackUrl *string `json:"callback_url,omitempty"`
	// 收件人
	Receiver *Receiver `json:"receiver,omitempty"`
	// 寄件人
	Sender *Sender `json:"sender,omitempty"`
	// 货品列表
	CargoList *[]CargoListItem `json:"cargo_list,omitempty"`
	// 拓展
	Extend map[string]string `json:"extend,omitempty"`
}
type Province struct {
	// 省
	Name *string `json:"name,omitempty"`
	// code
	Code string `json:"code,omitempty"`
}
type City struct {
	// 市
	Name *string `json:"name,omitempty"`
	// code
	Code string `json:"code,omitempty"`
}
type Address struct {
	// 省
	Province *Province `json:"province,omitempty"`
	// 市
	City *City `json:"city,omitempty"`
	// 区
	Town *Town `json:"town,omitempty"`
	// 街道
	Street *Street `json:"street,omitempty"`
	// 详细地址
	Detail *string `json:"detail,omitempty"`
}
type Receiver struct {
	// 联系人
	Contact *Contact `json:"contact,omitempty"`
	// 地址
	Address *Address `json:"address,omitempty"`
}
type Sender struct {
	// 联系人
	Contact *Contact `json:"contact,omitempty"`
	// 地址
	Address *Address `json:"address,omitempty"`
}
