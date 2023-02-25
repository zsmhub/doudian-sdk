package product_batchCreatePrettifyPic_request

import (
	"doudian.com/open/sdk_golang/api/product_batchCreatePrettifyPic/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type ProductBatchCreatePrettifyPicRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *ProductBatchCreatePrettifyPicParam 
}
func (c *ProductBatchCreatePrettifyPicRequest) GetUrlPath() string{
	return "/product/batchCreatePrettifyPic"
}


func New() *ProductBatchCreatePrettifyPicRequest{
	request := &ProductBatchCreatePrettifyPicRequest{
		Param: &ProductBatchCreatePrettifyPicParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *ProductBatchCreatePrettifyPicRequest) Execute(accessToken *doudian_sdk.AccessToken) (*product_batchCreatePrettifyPic_response.ProductBatchCreatePrettifyPicResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &product_batchCreatePrettifyPic_response.ProductBatchCreatePrettifyPicResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *ProductBatchCreatePrettifyPicRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *ProductBatchCreatePrettifyPicRequest) GetParams() *ProductBatchCreatePrettifyPicParam{
	return c.Param
}


type PrettifyPicDataText struct {
	// 文案内容 不可以超过5000字
	Text *string `json:"text,omitempty"`
	// 非必填。背景色
	BackgroundColor string `json:"background_color,omitempty"`
	// 非必填。字颜色
	FontColor string `json:"font_color,omitempty"`
	// 非必填。字号大小，在13~20之间
	FontSize int64 `json:"font_size,omitempty"`
	// 非必填。可选left、right、center；默认left
	TextAlign string `json:"text_align,omitempty"`
}
type PrettifyPicDataListItem struct {
	// 要转换的组件类型 目前可选:text  后续会推出其他类型的支持组件
	Name *string `json:"name,omitempty"`
	// 当name=text时有效
	PrettifyPicDataText *PrettifyPicDataText `json:"prettify_pic_data_text,omitempty"`
}
type ProductBatchCreatePrettifyPicParam struct {
	// 数组 用于批量构建图片。每批需要小于5
	PrettifyPicDataList *[]PrettifyPicDataListItem `json:"prettify_pic_data_list,omitempty"`
}
