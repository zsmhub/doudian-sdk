package material_createFolder_request

import (
	"doudian.com/open/sdk_golang/api/material_createFolder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialCreateFolderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialCreateFolderParam 
}
func (c *MaterialCreateFolderRequest) GetUrlPath() string{
	return "/material/createFolder"
}


func New() *MaterialCreateFolderRequest{
	request := &MaterialCreateFolderRequest{
		Param: &MaterialCreateFolderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialCreateFolderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_createFolder_response.MaterialCreateFolderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_createFolder_response.MaterialCreateFolderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialCreateFolderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialCreateFolderRequest) GetParams() *MaterialCreateFolderParam{
	return c.Param
}


type MaterialCreateFolderParam struct {
	// 文件夹名称。最多20字符
	Name *string `json:"name,omitempty"`
	// 父文件夹id，根目录为0。不传该参数默认在根目录下创建文件夹
	ParentFolderId string `json:"parent_folder_id,omitempty"`
	// 文件夹类型。0-用户自建
	Type *int32 `json:"type,omitempty"`
}
