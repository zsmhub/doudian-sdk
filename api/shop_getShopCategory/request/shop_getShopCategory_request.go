package shop_getShopCategory_request

import (
	"doudian.com/open/sdk_golang/api/shop_getShopCategory/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ShopGetShopCategoryRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ShopGetShopCategoryParam 
}
func (c *ShopGetShopCategoryRequest) GetUrlPath() string{
	return "/shop/getShopCategory"
}


func New() *ShopGetShopCategoryRequest{
	request := &ShopGetShopCategoryRequest{
		Param: &ShopGetShopCategoryParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ShopGetShopCategoryRequest) Execute(accessToken *doudian_sdk.AccessToken) (*shop_getShopCategory_response.ShopGetShopCategoryResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &shop_getShopCategory_response.ShopGetShopCategoryResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ShopGetShopCategoryRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ShopGetShopCategoryRequest) GetParams() *ShopGetShopCategoryParam{
	return c.Param
}


type ShopGetShopCategoryParam struct {
	// 父类目id，根据父id可以获取子类目。首次请求传值为0 可以获取所有一级类目。循环调用接口获取最小层级类目id，根据响应参数判断is_leaf=true或false。is_leaf=true表示是叶子节点，获取最小层级类目id，is_leaf=false表示不是子节点，请求参数cid=上一次响应参数id，直到获取最小层级类目id。
	Cid *int64 `json:"cid,omitempty"`
}
