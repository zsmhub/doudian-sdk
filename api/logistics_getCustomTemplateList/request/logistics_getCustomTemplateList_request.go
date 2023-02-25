package logistics_getCustomTemplateList_request

import (
	"doudian.com/open/sdk_golang/api/logistics_getCustomTemplateList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsGetCustomTemplateListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsGetCustomTemplateListParam 
}
func (c *LogisticsGetCustomTemplateListRequest) GetUrlPath() string{
	return "/logistics/getCustomTemplateList"
}


func New() *LogisticsGetCustomTemplateListRequest{
	request := &LogisticsGetCustomTemplateListRequest{
		Param: &LogisticsGetCustomTemplateListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsGetCustomTemplateListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_getCustomTemplateList_response.LogisticsGetCustomTemplateListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_getCustomTemplateList_response.LogisticsGetCustomTemplateListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsGetCustomTemplateListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsGetCustomTemplateListRequest) GetParams() *LogisticsGetCustomTemplateListParam{
	return c.Param
}


type LogisticsGetCustomTemplateListParam struct {
	// 物流服务商编码（若为空代表查询全部）
	LogisticsCode string `json:"logistics_code,omitempty"`
}
