package material_getFolderInfo_request

import (
	"doudian.com/open/sdk_golang/api/material_getFolderInfo/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialGetFolderInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialGetFolderInfoParam 
}
func (c *MaterialGetFolderInfoRequest) GetUrlPath() string{
	return "/material/getFolderInfo"
}


func New() *MaterialGetFolderInfoRequest{
	request := &MaterialGetFolderInfoRequest{
		Param: &MaterialGetFolderInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialGetFolderInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_getFolderInfo_response.MaterialGetFolderInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_getFolderInfo_response.MaterialGetFolderInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialGetFolderInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialGetFolderInfoRequest) GetParams() *MaterialGetFolderInfoParam{
	return c.Param
}


type MaterialGetFolderInfoParam struct {
	// 文件夹id
	FolderId *string `json:"folder_id,omitempty"`
	// 分页的页数，从1开始
	PageNum *int32 `json:"page_num,omitempty"`
	// 每页返回的数量。最大为100，默认为50
	PageSize *int32 `json:"page_size,omitempty"`
}
