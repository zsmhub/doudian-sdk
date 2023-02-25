package logistics_customTemplateList_request

import (
	"doudian.com/open/sdk_golang/api/logistics_customTemplateList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsCustomTemplateListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsCustomTemplateListParam 
}
func (c *LogisticsCustomTemplateListRequest) GetUrlPath() string{
	return "/logistics/customTemplateList"
}


func New() *LogisticsCustomTemplateListRequest{
	request := &LogisticsCustomTemplateListRequest{
		Param: &LogisticsCustomTemplateListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsCustomTemplateListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_customTemplateList_response.LogisticsCustomTemplateListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_customTemplateList_response.LogisticsCustomTemplateListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsCustomTemplateListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsCustomTemplateListRequest) GetParams() *LogisticsCustomTemplateListParam{
	return c.Param
}


type LogisticsCustomTemplateListParam struct {
	// 物流商编码
	LogisticsCode string `json:"logistics_code,omitempty"`
}
