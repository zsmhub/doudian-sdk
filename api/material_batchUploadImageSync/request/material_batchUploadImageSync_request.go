package material_batchUploadImageSync_request

import (
	"doudian.com/open/sdk_golang/api/material_batchUploadImageSync/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialBatchUploadImageSyncRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialBatchUploadImageSyncParam 
}
func (c *MaterialBatchUploadImageSyncRequest) GetUrlPath() string{
	return "/material/batchUploadImageSync"
}


func New() *MaterialBatchUploadImageSyncRequest{
	request := &MaterialBatchUploadImageSyncRequest{
		Param: &MaterialBatchUploadImageSyncParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialBatchUploadImageSyncRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_batchUploadImageSync_response.MaterialBatchUploadImageSyncResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_batchUploadImageSync_response.MaterialBatchUploadImageSyncResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialBatchUploadImageSyncRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialBatchUploadImageSyncRequest) GetParams() *MaterialBatchUploadImageSyncParam{
	return c.Param
}


type MaterialsItem struct {
	// 该参数仅有2个作用：（1）一次请求中素材的唯一标示；（2）接口防并发，规则是：不同请求所有request_id排序之后拼接起来，若相同视为同一次请求
	RequestId *string `json:"request_id,omitempty"`
	// 文件夹id，0为素材中心根目录。若想创建文件夹，请参考：https://ehome.bytedance.net/djt/apiManage/doc/preview/946?doc=true
	FolderId *string `json:"folder_id,omitempty"`
	// 素材名称，长度限制为50个字符
	Name *string `json:"name,omitempty"`
	// 图片url。如果是二进制上传，请使用file_uri字段。url和file_uri二选一，不能同时为空
	Url string `json:"url,omitempty"`
	// 二进制文件对应的uri，获取方式请参考：https://op.jinritemai.com/docs/guide-docs/171/1719
	FileUri string `json:"file_uri,omitempty"`
	// 素材类型，请传固定值：photo
	MaterialType *string `json:"material_type,omitempty"`
}
type MaterialBatchUploadImageSyncParam struct {
	// 素材信息
	Materials *[]MaterialsItem `json:"materials,omitempty"`
	// 是否需要去重（true/false），默认为false。去重是指：存在已经审核通过切内容内容相同的图片，直接返回已存在的图片地址。
	NeedDistinct bool `json:"need_distinct,omitempty"`
}
