package warehouse_info_request

import (
	"doudian.com/open/sdk_golang/api/warehouse_info/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type WarehouseInfoRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *WarehouseInfoParam 
}
func (c *WarehouseInfoRequest) GetUrlPath() string{
	return "/warehouse/info"
}


func New() *WarehouseInfoRequest{
	request := &WarehouseInfoRequest{
		Param: &WarehouseInfoParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *WarehouseInfoRequest) Execute(accessToken *doudian_sdk.AccessToken) (*warehouse_info_response.WarehouseInfoResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &warehouse_info_response.WarehouseInfoResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *WarehouseInfoRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *WarehouseInfoRequest) GetParams() *WarehouseInfoParam{
	return c.Param
}


type WarehouseInfoParam struct {
	// 仓库编码
	OutWarehouseId *string `json:"out_warehouse_id,omitempty"`
}
