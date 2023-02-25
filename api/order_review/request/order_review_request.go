package order_review_request

import (
	"doudian.com/open/sdk_golang/api/order_review/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type OrderReviewRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *OrderReviewParam 
}
func (c *OrderReviewRequest) GetUrlPath() string{
	return "/order/review"
}


func New() *OrderReviewRequest{
	request := &OrderReviewRequest{
		Param: &OrderReviewParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *OrderReviewRequest) Execute(accessToken *doudian_sdk.AccessToken) (*order_review_response.OrderReviewResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &order_review_response.OrderReviewResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *OrderReviewRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *OrderReviewRequest) GetParams() *OrderReviewParam{
	return c.Param
}


type OrderReviewParam struct {
	// 表示订单审核的类型 3001 通信卡审核
	TaskType *int64 `json:"task_type,omitempty"`
	// 0 审核通过 200001 下单身份信息180天内在该卡品运营商处重复下单，未通过审核 200002 下单身份信息已在该卡品运营商处办理了5张电话卡，未通过审核 200003 下单身份信息年龄不在16-60岁（部分卡品16-30岁），未通过审核 200004 下单地址为该卡品运营商禁售地区，未通过审核 200005 因其他原因，未能通过运营商审核，具体可联系商家
	RejectCode *int64 `json:"reject_code,omitempty"`
	// 审核的单id取决于什么审核 通信卡审核 店铺单id
	ObjectId *string `json:"object_id,omitempty"`
}
