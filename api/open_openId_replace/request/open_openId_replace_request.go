package open_openId_replace_request

import (
	"doudian.com/open/sdk_golang/api/open_openId_replace/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OpenOpenIdReplaceRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OpenOpenIdReplaceParam 
}
func (c *OpenOpenIdReplaceRequest) GetUrlPath() string{
	return "/open/openId/replace"
}


func New() *OpenOpenIdReplaceRequest{
	request := &OpenOpenIdReplaceRequest{
		Param: &OpenOpenIdReplaceParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OpenOpenIdReplaceRequest) Execute(accessToken *doudian_sdk.AccessToken) (*open_openId_replace_response.OpenOpenIdReplaceResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &open_openId_replace_response.OpenOpenIdReplaceResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OpenOpenIdReplaceRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OpenOpenIdReplaceRequest) GetParams() *OpenOpenIdReplaceParam{
	return c.Param
}


type OpenOpenIdReplaceParam struct {
	// 传入一个老openId列表, 最多传入100个openId, openId格式如示例
	OpenIds *[]string `json:"open_ids,omitempty"`
}
