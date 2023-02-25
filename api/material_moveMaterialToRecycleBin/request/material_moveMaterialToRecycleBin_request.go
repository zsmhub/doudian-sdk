package material_moveMaterialToRecycleBin_request

import (
	"doudian.com/open/sdk_golang/api/material_moveMaterialToRecycleBin/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialMoveMaterialToRecycleBinRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialMoveMaterialToRecycleBinParam 
}
func (c *MaterialMoveMaterialToRecycleBinRequest) GetUrlPath() string{
	return "/material/moveMaterialToRecycleBin"
}


func New() *MaterialMoveMaterialToRecycleBinRequest{
	request := &MaterialMoveMaterialToRecycleBinRequest{
		Param: &MaterialMoveMaterialToRecycleBinParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialMoveMaterialToRecycleBinRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_moveMaterialToRecycleBin_response.MaterialMoveMaterialToRecycleBinResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_moveMaterialToRecycleBin_response.MaterialMoveMaterialToRecycleBinResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialMoveMaterialToRecycleBinRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialMoveMaterialToRecycleBinRequest) GetParams() *MaterialMoveMaterialToRecycleBinParam{
	return c.Param
}


type MaterialMoveMaterialToRecycleBinParam struct {
	// 素材id列表，（1）数量不得超过100；（2）同一级目录下的素材
	MaterialIds *[]string `json:"material_ids,omitempty"`
}
