package afterSale_applyLogisticsIntercept_request

import (
	"doudian.com/open/sdk_golang/api/afterSale_applyLogisticsIntercept/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type AfterSaleApplyLogisticsInterceptRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *AfterSaleApplyLogisticsInterceptParam 
}
func (c *AfterSaleApplyLogisticsInterceptRequest) GetUrlPath() string{
	return "/afterSale/applyLogisticsIntercept"
}


func New() *AfterSaleApplyLogisticsInterceptRequest{
	request := &AfterSaleApplyLogisticsInterceptRequest{
		Param: &AfterSaleApplyLogisticsInterceptParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *AfterSaleApplyLogisticsInterceptRequest) Execute(accessToken *doudian_sdk.AccessToken) (*afterSale_applyLogisticsIntercept_response.AfterSaleApplyLogisticsInterceptResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &afterSale_applyLogisticsIntercept_response.AfterSaleApplyLogisticsInterceptResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *AfterSaleApplyLogisticsInterceptRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *AfterSaleApplyLogisticsInterceptRequest) GetParams() *AfterSaleApplyLogisticsInterceptParam{
	return c.Param
}


type InterceptTargetsItem struct {
	// 物流公司编码
	CompanyCode *string `json:"company_code,omitempty"`
	// 物流单号
	TrackingNo *string `json:"tracking_no,omitempty"`
}
type AfterSaleApplyLogisticsInterceptParam struct {
	// 售后单ID
	AfterSaleId *int64 `json:"after_sale_id,omitempty"`
	// 操作来源（1:商家  3:客服）
	OpFrom *int32 `json:"op_from,omitempty"`
	// 要拦截的包裹
	InterceptTargets *[]InterceptTargetsItem `json:"intercept_targets,omitempty"`
}
