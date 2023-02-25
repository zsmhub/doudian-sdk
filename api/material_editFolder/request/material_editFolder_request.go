package material_editFolder_request

import (
	"doudian.com/open/sdk_golang/api/material_editFolder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialEditFolderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialEditFolderParam 
}
func (c *MaterialEditFolderRequest) GetUrlPath() string{
	return "/material/editFolder"
}


func New() *MaterialEditFolderRequest{
	request := &MaterialEditFolderRequest{
		Param: &MaterialEditFolderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialEditFolderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_editFolder_response.MaterialEditFolderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_editFolder_response.MaterialEditFolderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialEditFolderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialEditFolderRequest) GetParams() *MaterialEditFolderParam{
	return c.Param
}


type MaterialEditFolderParam struct {
	// 文件夹id，不能操作系统文件夹（0：根目录 -1：回收站）
	FolderId *string `json:"folder_id,omitempty"`
	// 新的文件夹名称
	Name string `json:"name,omitempty"`
	// 需要移动到的父文件夹id
	ToFolderId string `json:"to_folder_id,omitempty"`
}
