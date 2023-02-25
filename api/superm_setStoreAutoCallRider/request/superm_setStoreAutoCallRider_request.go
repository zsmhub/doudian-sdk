package superm_setStoreAutoCallRider_request

import (
	"doudian.com/open/sdk_golang/api/superm_setStoreAutoCallRider/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SupermSetStoreAutoCallRiderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SupermSetStoreAutoCallRiderParam 
}
func (c *SupermSetStoreAutoCallRiderRequest) GetUrlPath() string{
	return "/superm/setStoreAutoCallRider"
}


func New() *SupermSetStoreAutoCallRiderRequest{
	request := &SupermSetStoreAutoCallRiderRequest{
		Param: &SupermSetStoreAutoCallRiderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SupermSetStoreAutoCallRiderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*superm_setStoreAutoCallRider_response.SupermSetStoreAutoCallRiderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &superm_setStoreAutoCallRider_response.SupermSetStoreAutoCallRiderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SupermSetStoreAutoCallRiderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SupermSetStoreAutoCallRiderRequest) GetParams() *SupermSetStoreAutoCallRiderParam{
	return c.Param
}


type SupermSetStoreAutoCallRiderParam struct {
	// 店铺的门店ID
	StoreId *int64 `json:"store_id,omitempty"`
	// 自动呼叫运力设置状态；open 开启 close  关闭
	OpType *string `json:"op_type,omitempty"`
	// 自动呼叫运力策略； 0：接单后立即呼叫  1：接单后延迟呼叫
	ServiceType int64 `json:"service_type,omitempty"`
	// 接单后延时呼叫时间，service_type为1时生效，单位分钟，范围[1,15] 延迟时间仅支持1-15分钟
	DelayTime int64 `json:"delay_time,omitempty"`
}
