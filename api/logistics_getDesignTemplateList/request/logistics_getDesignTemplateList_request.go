package logistics_getDesignTemplateList_request

import (
	"doudian.com/open/sdk_golang/api/logistics_getDesignTemplateList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsGetDesignTemplateListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsGetDesignTemplateListParam 
}
func (c *LogisticsGetDesignTemplateListRequest) GetUrlPath() string{
	return "/logistics/getDesignTemplateList"
}


func New() *LogisticsGetDesignTemplateListRequest{
	request := &LogisticsGetDesignTemplateListRequest{
		Param: &LogisticsGetDesignTemplateListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsGetDesignTemplateListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_getDesignTemplateList_response.LogisticsGetDesignTemplateListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_getDesignTemplateList_response.LogisticsGetDesignTemplateListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsGetDesignTemplateListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsGetDesignTemplateListRequest) GetParams() *LogisticsGetDesignTemplateListParam{
	return c.Param
}


type LogisticsGetDesignTemplateListParam struct {
}
