package spu_batchUploadImg_request

import (
	"doudian.com/open/sdk_golang/api/spu_batchUploadImg/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type SpuBatchUploadImgRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *SpuBatchUploadImgParam 
}
func (c *SpuBatchUploadImgRequest) GetUrlPath() string{
	return "/spu/batchUploadImg"
}


func New() *SpuBatchUploadImgRequest{
	request := &SpuBatchUploadImgRequest{
		Param: &SpuBatchUploadImgParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *SpuBatchUploadImgRequest) Execute(accessToken *doudian_sdk.AccessToken) (*spu_batchUploadImg_response.SpuBatchUploadImgResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &spu_batchUploadImg_response.SpuBatchUploadImgResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *SpuBatchUploadImgRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *SpuBatchUploadImgRequest) GetParams() *SpuBatchUploadImgParam{
	return c.Param
}


type ImageUrlListItem struct {
	// 图片名
	Name *string `json:"name,omitempty"`
	// 公网可访问URL
	Url *string `json:"url,omitempty"`
}
type SpuBatchUploadImgParam struct {
	// 图片列表
	ImageUrlList *[]ImageUrlListItem `json:"image_url_list,omitempty"`
}
