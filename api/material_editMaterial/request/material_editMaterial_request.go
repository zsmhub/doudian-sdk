package material_editMaterial_request

import (
	"doudian.com/open/sdk_golang/api/material_editMaterial/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialEditMaterialRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialEditMaterialParam 
}
func (c *MaterialEditMaterialRequest) GetUrlPath() string{
	return "/material/editMaterial"
}


func New() *MaterialEditMaterialRequest{
	request := &MaterialEditMaterialRequest{
		Param: &MaterialEditMaterialParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialEditMaterialRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_editMaterial_response.MaterialEditMaterialResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_editMaterial_response.MaterialEditMaterialResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialEditMaterialRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialEditMaterialRequest) GetParams() *MaterialEditMaterialParam{
	return c.Param
}


type MaterialEditMaterialParam struct {
	// 素材id
	MaterialId *string `json:"material_id,omitempty"`
	// 素材名称，不得超过50个字符
	MaterialName string `json:"material_name,omitempty"`
	// 目标文件夹id，"0"--素材中心
	ToFolderId string `json:"to_folder_id,omitempty"`
}
