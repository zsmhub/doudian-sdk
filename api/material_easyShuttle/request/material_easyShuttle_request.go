package material_easyShuttle_request

import (
	"doudian.com/open/sdk_golang/api/material_easyShuttle/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialEasyShuttleRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialEasyShuttleParam 
}
func (c *MaterialEasyShuttleRequest) GetUrlPath() string{
	return "/material/easyShuttle"
}


func New() *MaterialEasyShuttleRequest{
	request := &MaterialEasyShuttleRequest{
		Param: &MaterialEasyShuttleParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialEasyShuttleRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_easyShuttle_response.MaterialEasyShuttleResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_easyShuttle_response.MaterialEasyShuttleResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialEasyShuttleRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialEasyShuttleRequest) GetParams() *MaterialEasyShuttleParam{
	return c.Param
}


type MaterialEasyShuttleParam struct {
	// 文件夹ID列表，对该文件夹下的文件夹或素材进行操作；0:开放平台；1:图片根目录；2:视频根目录；-1:回收站
	FolderIdList *[]string `json:"folder_id_list,omitempty"`
	// 操作类型：移到回收站：to_recycle，彻底删除：delete_forever
	OperateType *string `json:"operate_type,omitempty"`
	// 操作的素材类型
	MaterialTypeList *[]string `json:"material_type_list,omitempty"`
	// 是否只操作素材
	OnlyMaterial bool `json:"only_material,omitempty"`
	// 所操作的素材或文件夹的创建时间在该时间点之前
	CreateTimeEnd string `json:"create_time_end,omitempty"`
}
