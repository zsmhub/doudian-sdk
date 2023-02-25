package warehouse_setAddr_request

import (
	"doudian.com/open/sdk_golang/api/warehouse_setAddr/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type WarehouseSetAddrRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *WarehouseSetAddrParam 
}
func (c *WarehouseSetAddrRequest) GetUrlPath() string{
	return "/warehouse/setAddr"
}


func New() *WarehouseSetAddrRequest{
	request := &WarehouseSetAddrRequest{
		Param: &WarehouseSetAddrParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *WarehouseSetAddrRequest) Execute(accessToken *doudian_sdk.AccessToken) (*warehouse_setAddr_response.WarehouseSetAddrResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &warehouse_setAddr_response.WarehouseSetAddrResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *WarehouseSetAddrRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *WarehouseSetAddrRequest) GetParams() *WarehouseSetAddrParam{
	return c.Param
}


type Addr struct {
	// 一级地址
	AddrId1 *int64 `json:"addr_id1,omitempty"`
	// 二级地址
	AddrId2 int64 `json:"addr_id2,omitempty"`
	// 三级地址
	AddrId3 int64 `json:"addr_id3,omitempty"`
	// 四级地址
	AddrId4 int64 `json:"addr_id4,omitempty"`
}
type WarehouseSetAddrParam struct {
	// 外部仓库ID
	OutWarehouseId *string `json:"out_warehouse_id,omitempty"`
	// 仓库地址
	Addr *Addr `json:"addr,omitempty"`
}
