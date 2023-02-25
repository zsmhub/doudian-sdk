package material_moveFolderToRecycleBin_request

import (
	"doudian.com/open/sdk_golang/api/material_moveFolderToRecycleBin/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialMoveFolderToRecycleBinRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialMoveFolderToRecycleBinParam 
}
func (c *MaterialMoveFolderToRecycleBinRequest) GetUrlPath() string{
	return "/material/moveFolderToRecycleBin"
}


func New() *MaterialMoveFolderToRecycleBinRequest{
	request := &MaterialMoveFolderToRecycleBinRequest{
		Param: &MaterialMoveFolderToRecycleBinParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialMoveFolderToRecycleBinRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_moveFolderToRecycleBin_response.MaterialMoveFolderToRecycleBinResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_moveFolderToRecycleBin_response.MaterialMoveFolderToRecycleBinResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialMoveFolderToRecycleBinRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialMoveFolderToRecycleBinRequest) GetParams() *MaterialMoveFolderToRecycleBinParam{
	return c.Param
}


type MaterialMoveFolderToRecycleBinParam struct {
	// 移动到回收站的文件夹id列表
	FolderIds *[]string `json:"folder_ids,omitempty"`
}
