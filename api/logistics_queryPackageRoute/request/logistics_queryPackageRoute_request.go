package logistics_queryPackageRoute_request

import (
	"doudian.com/open/sdk_golang/api/logistics_queryPackageRoute/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsQueryPackageRouteRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsQueryPackageRouteParam 
}
func (c *LogisticsQueryPackageRouteRequest) GetUrlPath() string{
	return "/logistics/queryPackageRoute"
}


func New() *LogisticsQueryPackageRouteRequest{
	request := &LogisticsQueryPackageRouteRequest{
		Param: &LogisticsQueryPackageRouteParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsQueryPackageRouteRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_queryPackageRoute_response.LogisticsQueryPackageRouteResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_queryPackageRoute_response.LogisticsQueryPackageRouteResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsQueryPackageRouteRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsQueryPackageRouteRequest) GetParams() *LogisticsQueryPackageRouteParam{
	return c.Param
}


type Receiver struct {
	// 收件人姓名
	Name *string `json:"name,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 虚拟手机号
	VirtualMobile string `json:"virtual_mobile,omitempty"`
}
type Sender struct {
	// 寄件人姓名
	Name *string `json:"name,omitempty"`
	// 手机号
	Phone string `json:"phone,omitempty"`
	// 手机号
	Mobile string `json:"mobile,omitempty"`
	// 邮箱
	Email string `json:"email,omitempty"`
	// 虚拟手机号
	VirtualMobile string `json:"virtual_mobile,omitempty"`
}
type LogisticsQueryPackageRouteParam struct {
	// 运单号
	TrackNo *string `json:"track_no,omitempty"`
	// 物流公司
	Express *string `json:"express,omitempty"`
	// 收件人
	Receiver *Receiver `json:"receiver,omitempty"`
	// 寄件人
	Sender *Sender `json:"sender,omitempty"`
}
