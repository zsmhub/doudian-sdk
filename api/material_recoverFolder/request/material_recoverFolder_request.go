package material_recoverFolder_request

import (
	"doudian.com/open/sdk_golang/api/material_recoverFolder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialRecoverFolderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialRecoverFolderParam 
}
func (c *MaterialRecoverFolderRequest) GetUrlPath() string{
	return "/material/recoverFolder"
}


func New() *MaterialRecoverFolderRequest{
	request := &MaterialRecoverFolderRequest{
		Param: &MaterialRecoverFolderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialRecoverFolderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_recoverFolder_response.MaterialRecoverFolderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_recoverFolder_response.MaterialRecoverFolderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialRecoverFolderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialRecoverFolderRequest) GetParams() *MaterialRecoverFolderParam{
	return c.Param
}


type MaterialRecoverFolderParam struct {
	// 需要恢复的文件夹的id列表，不能操作系统文件夹（0：根目录 -1：回收站）。批量操作每次最多20个。
	FolderIds []string `json:"folder_ids,omitempty"`
}
