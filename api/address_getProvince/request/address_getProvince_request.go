package address_getProvince_request

import (
	"doudian.com/open/sdk_golang/api/address_getProvince/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AddressGetProvinceRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AddressGetProvinceParam 
}
func (c *AddressGetProvinceRequest) GetUrlPath() string{
	return "/address/getProvince"
}


func New() *AddressGetProvinceRequest{
	request := &AddressGetProvinceRequest{
		Param: &AddressGetProvinceParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AddressGetProvinceRequest) Execute(accessToken *doudian_sdk.AccessToken) (*address_getProvince_response.AddressGetProvinceResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &address_getProvince_response.AddressGetProvinceResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AddressGetProvinceRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AddressGetProvinceRequest) GetParams() *AddressGetProvinceParam{
	return c.Param
}


type AddressGetProvinceParam struct {
}
