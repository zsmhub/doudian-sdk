package open_binaryupload_request

import (
	"doudian.com/open/sdk_golang/api/open_binaryupload/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OpenBinaryuploadRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OpenBinaryuploadParam 
}
func (c *OpenBinaryuploadRequest) GetUrlPath() string{
	return "/open/binaryupload"
}


func New() *OpenBinaryuploadRequest{
	request := &OpenBinaryuploadRequest{
		Param: &OpenBinaryuploadParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OpenBinaryuploadRequest) Execute(accessToken *doudian_sdk.AccessToken) (*open_binaryupload_response.OpenBinaryuploadResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &open_binaryupload_response.OpenBinaryuploadResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OpenBinaryuploadRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OpenBinaryuploadRequest) GetParams() *OpenBinaryuploadParam{
	return c.Param
}


type OpenBinaryuploadParam struct {
	// 二进制内容，utf-8编码
	FileContent *string `json:"file_content,omitempty"`
}
