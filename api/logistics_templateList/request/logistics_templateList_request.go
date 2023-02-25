package logistics_templateList_request

import (
	"doudian.com/open/sdk_golang/api/logistics_templateList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsTemplateListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsTemplateListParam 
}
func (c *LogisticsTemplateListRequest) GetUrlPath() string{
	return "/logistics/templateList"
}


func New() *LogisticsTemplateListRequest{
	request := &LogisticsTemplateListRequest{
		Param: &LogisticsTemplateListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsTemplateListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_templateList_response.LogisticsTemplateListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_templateList_response.LogisticsTemplateListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsTemplateListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsTemplateListRequest) GetParams() *LogisticsTemplateListParam{
	return c.Param
}


type LogisticsTemplateListParam struct {
}
