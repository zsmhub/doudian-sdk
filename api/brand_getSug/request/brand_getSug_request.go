package brand_getSug_request

import (
	"doudian.com/open/sdk_golang/api/brand_getSug/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type BrandGetSugRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *BrandGetSugParam 
}
func (c *BrandGetSugRequest) GetUrlPath() string{
	return "/brand/getSug"
}


func New() *BrandGetSugRequest{
	request := &BrandGetSugRequest{
		Param: &BrandGetSugParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *BrandGetSugRequest) Execute(accessToken *doudian_sdk.AccessToken) (*brand_getSug_response.BrandGetSugResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &brand_getSug_response.BrandGetSugResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *BrandGetSugRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *BrandGetSugRequest) GetParams() *BrandGetSugParam{
	return c.Param
}


type FilterInfo struct {
	// 品牌ids
	BrandIds []int64 `json:"brand_ids,omitempty"`
	// 品牌类别
	BrandCategory []int64 `json:"brand_category,omitempty"`
	// 品牌状态: 1.在线 2.离线
	Status int64 `json:"status,omitempty"`
	// 品牌商标关联Id
	RelatedIds []int64 `json:"related_ids,omitempty"`
	// 商标IDs
	TradeMarkIds []string `json:"trade_mark_ids,omitempty"`
	// 废弃字段，请勿填写: 1. 审核中 2. 审核通过 3. 审核拒绝 4. 送审失败
	AuditStatus []int32 `json:"audit_status,omitempty"`
}
type ExtraConfig struct {
	// 是否忽略去重 使用原始品牌信息，默认取false
	UseOriginBrandInfo bool `json:"use_origin_brand_info,omitempty"`
	// 是否忽略新旧映射 使用老品牌信息，默认取false
	UseBrandInfo bool `json:"use_brand_info,omitempty"`
	// 使用品牌名去重，需要和抖店一致请取true
	UseBrandNameDeduplicate bool `json:"use_brand_name_deduplicate,omitempty"`
}
type BrandGetSugParam struct {
	// 前缀匹配的品牌名
	Query *string `json:"query,omitempty"`
	// 用户ID，可用默认值0
	UserId *int64 `json:"user_id,omitempty"`
	// 过滤用参数，不填则是全量召回
	FilterInfo *FilterInfo `json:"filter_info,omitempty"`
	// 是否读取老数据 默认为false
	ReadOld *bool `json:"read_old,omitempty"`
	// 业务线类型: 0. 国内品牌 1. 跨境品牌 3. 广告
	BizTypes *[]int32 `json:"biz_types,omitempty"`
	// 是否去重，一般选择true
	EnableDeduplicate *bool `json:"enable_deduplicate,omitempty"`
	// 额外配置，无特殊需求请按描述填写
	ExtraConfig *ExtraConfig `json:"extra_config,omitempty"`
}
