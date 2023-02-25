package material_uploadImageSync_request

import (
	"doudian.com/open/sdk_golang/api/material_uploadImageSync/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialUploadImageSyncRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialUploadImageSyncParam 
}
func (c *MaterialUploadImageSyncRequest) GetUrlPath() string{
	return "/material/uploadImageSync"
}


func New() *MaterialUploadImageSyncRequest{
	request := &MaterialUploadImageSyncRequest{
		Param: &MaterialUploadImageSyncParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialUploadImageSyncRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_uploadImageSync_response.MaterialUploadImageSyncResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_uploadImageSync_response.MaterialUploadImageSyncResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialUploadImageSyncRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialUploadImageSyncRequest) GetParams() *MaterialUploadImageSyncParam{
	return c.Param
}


type MaterialUploadImageSyncParam struct {
	// 文件夹id，0为根目录
	FolderId *string `json:"folder_id,omitempty"`
	// 图片url，必须是公网可访问。url和file_uri二选一，不能同时为空，如果2者都不为空取url
	Url string `json:"url,omitempty"`
	// 图片名称，开发者自定义，不得超过50个字符。
	MaterialName *string `json:"material_name,omitempty"`
	// 二进制文件对应的uri，获取方式请参考：https://op.jinritemai.com/docs/guide-docs/171/1719
	FileUri string `json:"file_uri,omitempty"`
	// 是否需要去重（true/false），默认为false。去重是指：存在已经审核通过且内容相同的图片，直接返回已存在的图片地址。
	NeedDistinct bool `json:"need_distinct,omitempty"`
}
