package address_areaList_request

import (
	"doudian.com/open/sdk_golang/api/address_areaList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AddressAreaListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AddressAreaListParam 
}
func (c *AddressAreaListRequest) GetUrlPath() string{
	return "/address/areaList"
}


func New() *AddressAreaListRequest{
	request := &AddressAreaListRequest{
		Param: &AddressAreaListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AddressAreaListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*address_areaList_response.AddressAreaListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &address_areaList_response.AddressAreaListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AddressAreaListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AddressAreaListRequest) GetParams() *AddressAreaListParam{
	return c.Param
}


type AddressAreaListParam struct {
	CityId *int64 `json:"city_id,omitempty"`
}
