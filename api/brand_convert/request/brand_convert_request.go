package brand_convert_request

import (
	"doudian.com/open/sdk_golang/api/brand_convert/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BrandConvertRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BrandConvertParam 
}
func (c *BrandConvertRequest) GetUrlPath() string{
	return "/brand/convert"
}


func New() *BrandConvertRequest{
	request := &BrandConvertRequest{
		Param: &BrandConvertParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *BrandConvertRequest) Execute(accessToken *doudian_sdk.AccessToken) (*brand_convert_response.BrandConvertResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &brand_convert_response.BrandConvertResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *BrandConvertRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *BrandConvertRequest) GetParams() *BrandConvertParam{
	return c.Param
}


type BrandConvertParam struct {
	// 品牌关系id，即/shop/brandList返回的id
	RelatedId *int64 `json:"related_id,omitempty"`
}
