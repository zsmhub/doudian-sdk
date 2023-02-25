package material_recoverMaterial_request

import (
	"doudian.com/open/sdk_golang/api/material_recoverMaterial/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialRecoverMaterialRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialRecoverMaterialParam 
}
func (c *MaterialRecoverMaterialRequest) GetUrlPath() string{
	return "/material/recoverMaterial"
}


func New() *MaterialRecoverMaterialRequest{
	request := &MaterialRecoverMaterialRequest{
		Param: &MaterialRecoverMaterialParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialRecoverMaterialRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_recoverMaterial_response.MaterialRecoverMaterialResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_recoverMaterial_response.MaterialRecoverMaterialResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialRecoverMaterialRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialRecoverMaterialRequest) GetParams() *MaterialRecoverMaterialParam{
	return c.Param
}


type MaterialRecoverMaterialParam struct {
	// 素材id列表
	MaterialIds *[]string `json:"material_ids,omitempty"`
}
