package material_get_cap_info_request

import (
	"doudian.com/open/sdk_golang/api/material_get_cap_info/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialGetCapInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialGetCapInfoParam 
}
func (c *MaterialGetCapInfoRequest) GetUrlPath() string{
	return "/material/get_cap_info"
}


func New() *MaterialGetCapInfoRequest{
	request := &MaterialGetCapInfoRequest{
		Param: &MaterialGetCapInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialGetCapInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_get_cap_info_response.MaterialGetCapInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_get_cap_info_response.MaterialGetCapInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialGetCapInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialGetCapInfoRequest) GetParams() *MaterialGetCapInfoParam{
	return c.Param
}


type MaterialGetCapInfoParam struct {
}
