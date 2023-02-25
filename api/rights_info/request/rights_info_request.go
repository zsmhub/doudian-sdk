package rights_info_request

import (
	"doudian.com/open/sdk_golang/api/rights_info/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type RightsInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *RightsInfoParam 
}
func (c *RightsInfoRequest) GetUrlPath() string{
	return "/rights/info"
}


func New() *RightsInfoRequest{
	request := &RightsInfoRequest{
		Param: &RightsInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *RightsInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*rights_info_response.RightsInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &rights_info_response.RightsInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *RightsInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *RightsInfoRequest) GetParams() *RightsInfoParam{
	return c.Param
}


type RightsInfoParam struct {
	// 授权主体类型，不传参默认查询店铺，2-ebill用户；3-供应商
	BizType int32 `json:"biz_type,omitempty"`
}
