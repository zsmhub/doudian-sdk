package material_batchUploadVideoAsync_request

import (
	"doudian.com/open/sdk_golang/api/material_batchUploadVideoAsync/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialBatchUploadVideoAsyncRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialBatchUploadVideoAsyncParam 
}
func (c *MaterialBatchUploadVideoAsyncRequest) GetUrlPath() string{
	return "/material/batchUploadVideoAsync"
}


func New() *MaterialBatchUploadVideoAsyncRequest{
	request := &MaterialBatchUploadVideoAsyncRequest{
		Param: &MaterialBatchUploadVideoAsyncParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialBatchUploadVideoAsyncRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_batchUploadVideoAsync_response.MaterialBatchUploadVideoAsyncResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_batchUploadVideoAsync_response.MaterialBatchUploadVideoAsyncResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialBatchUploadVideoAsyncRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialBatchUploadVideoAsyncRequest) GetParams() *MaterialBatchUploadVideoAsyncParam{
	return c.Param
}


type MaterialsItem struct {
	// 该参数仅有2个作用：（1）一次请求中素材的唯一标示；（2）接口防并发，规则是：不同请求所有request_id排序之后拼接起来，若相同视为同一次请求
	RequestId *string `json:"request_id,omitempty"`
	// 文件夹id，“0”为素材中心根目录。若想创建文件夹，请参考：https://ehome.bytedance.net/djt/apiManage/doc/preview/946?doc=true
	FolderId *string `json:"folder_id,omitempty"`
	// 素材名称，长度限制为50个字符，最好带上后缀
	Name *string `json:"name,omitempty"`
	// 视频url。如果是二进制上传，请使用file_uri字段。url和file_uri二选一，不能同时为空
	Url string `json:"url,omitempty"`
	// 二进制文件对应的uri，获取方式请参考：https://op.jinritemai.com/docs/guide-docs/171/1719
	FileUri string `json:"file_uri,omitempty"`
	// 素材类型，请传固定值：video
	MaterialType *string `json:"material_type,omitempty"`
}
type MaterialBatchUploadVideoAsyncParam struct {
	// 素材信息
	Materials *[]MaterialsItem `json:"materials,omitempty"`
}
