package address_create_request

import (
	"doudian.com/open/sdk_golang/api/address_create/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AddressCreateRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AddressCreateParam 
}
func (c *AddressCreateRequest) GetUrlPath() string{
	return "/address/create"
}


func New() *AddressCreateRequest{
	request := &AddressCreateRequest{
		Param: &AddressCreateParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AddressCreateRequest) Execute(accessToken *doudian_sdk.AccessToken) (*address_create_response.AddressCreateResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &address_create_response.AddressCreateResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AddressCreateRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AddressCreateRequest) GetParams() *AddressCreateParam{
	return c.Param
}


type Address struct {
	// 联系人姓名
	UserName *string `json:"user_name,omitempty"`
	// 手机号码
	Mobile string `json:"mobile,omitempty"`
	// 省份ID
	ProvinceId *int64 `json:"province_id,omitempty"`
	// 城市ID
	CityId *int64 `json:"city_id,omitempty"`
	// 区ID
	TownId *int64 `json:"town_id,omitempty"`
	// 详细地址
	Detail *string `json:"detail,omitempty"`
	// 街道ID
	StreetId *int64 `json:"street_id,omitempty"`
	// 联系方式类型(0-手机,1-普通座机,2-企业座机)
	LinkType *int32 `json:"link_type,omitempty"`
	// 普通座机格式：区号-座机号-分机号(分机号选填)、区号3~4位、座机号7~8位、分机号不超过5位。企业座机：400/800开头不超过10位、95开头在5~8we
	FixedPhone string `json:"fixed_phone,omitempty"`
	// 售后备注，限制200字符
	Remark string `json:"remark,omitempty"`
}
type AddressCreateParam struct {
	// 地址信息
	Address *Address `json:"address,omitempty"`
	// 门店ID（新建地址绑定在该门店下，非门店场景无需填写）
	StoreId int64 `json:"store_id,omitempty"`
}
