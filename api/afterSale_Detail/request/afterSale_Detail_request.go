package afterSale_Detail_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_Detail/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleDetailRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleDetailParam 
}
func (c *AfterSaleDetailRequest) GetUrlPath() string{
	return "/afterSale/Detail"
}


func New() *AfterSaleDetailRequest{
	request := &AfterSaleDetailRequest{
		Param: &AfterSaleDetailParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleDetailRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_Detail_response.AfterSaleDetailResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_Detail_response.AfterSaleDetailResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleDetailRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleDetailRequest) GetParams() *AfterSaleDetailParam{
	return c.Param
}


type AfterSaleDetailParam struct {
	// 售后单ID
	AfterSaleId *string `json:"after_sale_id,omitempty"`
	// 是否需要协商记录
	NeedOperationRecord bool `json:"need_operation_record,omitempty"`
	IsSearchable bool `json:"is_searchable,omitempty,omitempty"`
}
