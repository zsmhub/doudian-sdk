package order_getServiceList_request

import (
	"doudian.com/open/sdk_golang/api/order_getServiceList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderGetServiceListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderGetServiceListParam 
}
func (c *OrderGetServiceListRequest) GetUrlPath() string{
	return "/order/getServiceList"
}


func New() *OrderGetServiceListRequest{
	request := &OrderGetServiceListRequest{
		Param: &OrderGetServiceListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderGetServiceListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_getServiceList_response.OrderGetServiceListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_getServiceList_response.OrderGetServiceListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderGetServiceListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderGetServiceListRequest) GetParams() *OrderGetServiceListParam{
	return c.Param
}


type OrderGetServiceListParam struct {
	// 开始时间时间戳
	StartTime *int64 `json:"start_time,omitempty"`
	// 结束时间时间戳，必须大于开始时间
	EndTime *int64 `json:"end_time,omitempty"`
	// 1、不传代表获取全部服务请求     2、操作状态：0 #待处理 10 #审核中 100 #已关闭 
	Status int32 `json:"status,omitempty"`
	// 订单id
	OrderId int64 `json:"order_id,omitempty"`
	// 服务单id
	ServiceId int64 `json:"service_id,omitempty"`
	// 1、默认按服务单创建时间搜索     2、值为“create_time”：按服务单创建时间；值为“update_time”：按服务单更新时间
	OrderBy string `json:"order_by,omitempty"`
	// 页数（默认值为1，第一页从1开始）
	Page int32 `json:"page,omitempty"`
	// 每页订单数（默认为10，最大100），超过100则会按照最大值100处理
	Size int32 `json:"size,omitempty"`
	// 排序方式：ASC按时间升序，  DESC按时间降序     默认DESC    
	Order string `json:"order,omitempty"`
}
