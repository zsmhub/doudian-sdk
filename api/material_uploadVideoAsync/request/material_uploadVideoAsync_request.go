package material_uploadVideoAsync_request

import (
	"doudian.com/open/sdk_golang/api/material_uploadVideoAsync/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialUploadVideoAsyncRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialUploadVideoAsyncParam 
}
func (c *MaterialUploadVideoAsyncRequest) GetUrlPath() string{
	return "/material/uploadVideoAsync"
}


func New() *MaterialUploadVideoAsyncRequest{
	request := &MaterialUploadVideoAsyncRequest{
		Param: &MaterialUploadVideoAsyncParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialUploadVideoAsyncRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_uploadVideoAsync_response.MaterialUploadVideoAsyncResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_uploadVideoAsync_response.MaterialUploadVideoAsyncResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialUploadVideoAsyncRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialUploadVideoAsyncRequest) GetParams() *MaterialUploadVideoAsyncParam{
	return c.Param
}


type MaterialUploadVideoAsyncParam struct {
	// 父文件夹id，0为根目录。若需要创建文件夹，请参考：https://ehome.bytedance.net/djt/apiManage/doc/preview/946?doc=true
	FolderId *string `json:"folder_id,omitempty"`
	// 视频url，url和file_uri二选一，不能同时为空，如果2者都传则取url
	Url string `json:"url,omitempty"`
	// 视频名称，不得超过50个字符，最好带上后缀
	Name *string `json:"name,omitempty"`
	// 二进制文件对应的uri，获取方式请参考：https://op.jinritemai.com/docs/guide-docs/171/1719
	FileUri string `json:"file_uri,omitempty"`
}
