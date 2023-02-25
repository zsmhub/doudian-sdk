package material_deleteMaterial_request

import (
	"doudian.com/open/sdk_golang/api/material_deleteMaterial/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialDeleteMaterialRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialDeleteMaterialParam 
}
func (c *MaterialDeleteMaterialRequest) GetUrlPath() string{
	return "/material/deleteMaterial"
}


func New() *MaterialDeleteMaterialRequest{
	request := &MaterialDeleteMaterialRequest{
		Param: &MaterialDeleteMaterialParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialDeleteMaterialRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_deleteMaterial_response.MaterialDeleteMaterialResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_deleteMaterial_response.MaterialDeleteMaterialResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialDeleteMaterialRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialDeleteMaterialRequest) GetParams() *MaterialDeleteMaterialParam{
	return c.Param
}


type MaterialDeleteMaterialParam struct {
	// 要删除的素材id列表，最大不得超过100个，只支持删除同一级目录下的素材
	MaterialIds *[]string `json:"material_ids,omitempty"`
}
