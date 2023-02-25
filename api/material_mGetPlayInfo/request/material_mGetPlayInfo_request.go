package material_mGetPlayInfo_request

import (
	"doudian.com/open/sdk_golang/api/material_mGetPlayInfo/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialMGetPlayInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialMGetPlayInfoParam 
}
func (c *MaterialMGetPlayInfoRequest) GetUrlPath() string{
	return "/material/mGetPlayInfo"
}


func New() *MaterialMGetPlayInfoRequest{
	request := &MaterialMGetPlayInfoRequest{
		Param: &MaterialMGetPlayInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialMGetPlayInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_mGetPlayInfo_response.MaterialMGetPlayInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_mGetPlayInfo_response.MaterialMGetPlayInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialMGetPlayInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialMGetPlayInfoRequest) GetParams() *MaterialMGetPlayInfoParam{
	return c.Param
}


type MaterialMGetPlayInfoParam struct {
	// vid列表，列表长度(1, 20)
	VidList *[]string `json:"vid_list,omitempty"`
}
