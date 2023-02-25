package logistics_listShopNetsite_request

import (
	"doudian.com/open/sdk_golang/api/logistics_listShopNetsite/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type LogisticsListShopNetsiteRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *LogisticsListShopNetsiteParam 
}
func (c *LogisticsListShopNetsiteRequest) GetUrlPath() string{
	return "/logistics/listShopNetsite"
}


func New() *LogisticsListShopNetsiteRequest{
	request := &LogisticsListShopNetsiteRequest{
		Param: &LogisticsListShopNetsiteParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *LogisticsListShopNetsiteRequest) Execute(accessToken *doudian_sdk.AccessToken) (*logistics_listShopNetsite_response.LogisticsListShopNetsiteResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &logistics_listShopNetsite_response.LogisticsListShopNetsiteResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *LogisticsListShopNetsiteRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *LogisticsListShopNetsiteRequest) GetParams() *LogisticsListShopNetsiteParam{
	return c.Param
}


type LogisticsListShopNetsiteParam struct {
	// 物流服务商编码（想获取全量物流，则传空字符串）
	LogisticsCode *string `json:"logistics_code,omitempty"`
}
