package material_queryMaterialDetail_request

import (
	"doudian.com/open/sdk_golang/api/material_queryMaterialDetail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialQueryMaterialDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialQueryMaterialDetailParam 
}
func (c *MaterialQueryMaterialDetailRequest) GetUrlPath() string{
	return "/material/queryMaterialDetail"
}


func New() *MaterialQueryMaterialDetailRequest{
	request := &MaterialQueryMaterialDetailRequest{
		Param: &MaterialQueryMaterialDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialQueryMaterialDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_queryMaterialDetail_response.MaterialQueryMaterialDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_queryMaterialDetail_response.MaterialQueryMaterialDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialQueryMaterialDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialQueryMaterialDetailRequest) GetParams() *MaterialQueryMaterialDetailParam{
	return c.Param
}


type MaterialQueryMaterialDetailParam struct {
	// 素材id；可使用【/material/batchUploadImageSync】【/material/uploadImageSync】【/material/searchMaterial  】接口获取
	MaterialId *string `json:"material_id,omitempty"`
}
