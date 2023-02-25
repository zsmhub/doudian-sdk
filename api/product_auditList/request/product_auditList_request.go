package product_auditList_request

import (
	"doudian.com/open/sdk_golang/api/product_auditList/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductAuditListRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductAuditListParam 
}
func (c *ProductAuditListRequest) GetUrlPath() string{
	return "/product/auditList"
}


func New() *ProductAuditListRequest{
	request := &ProductAuditListRequest{
		Param: &ProductAuditListParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductAuditListRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_auditList_response.ProductAuditListResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_auditList_response.ProductAuditListResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductAuditListRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductAuditListRequest) GetParams() *ProductAuditListParam{
	return c.Param
}


type ProductAuditListParam struct {
	// 指定审核状态返回商品列表：0-审核中 1-审核通过 2-审核拒绝
	PublishStatus *int64 `json:"publish_status,omitempty"`
	// 第几页（第一页为0，最大为99）
	Page *int64 `json:"page,omitempty"`
	// 每页返回条数，最多支持100条
	Size *int64 `json:"size,omitempty"`
	// 商品id
	ProductId int64 `json:"product_id,omitempty"`
}
