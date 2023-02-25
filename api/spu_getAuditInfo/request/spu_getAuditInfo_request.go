package spu_getAuditInfo_request

import (
	"doudian.com/open/sdk_golang/api/spu_getAuditInfo/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuGetAuditInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuGetAuditInfoParam 
}
func (c *SpuGetAuditInfoRequest) GetUrlPath() string{
	return "/spu/getAuditInfo"
}


func New() *SpuGetAuditInfoRequest{
	request := &SpuGetAuditInfoRequest{
		Param: &SpuGetAuditInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuGetAuditInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_getAuditInfo_response.SpuGetAuditInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_getAuditInfo_response.SpuGetAuditInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuGetAuditInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuGetAuditInfoRequest) GetParams() *SpuGetAuditInfoParam{
	return c.Param
}


type SpuGetAuditInfoParam struct {
	// SPU编号；使用【/spu/createSpu】创建spuid创建成功后，可以使用spu_id查询审核状态；建议场景成功后间隔1小时后再查询审核结果；
	SpuId *int64 `json:"spu_id,omitempty"`
}
