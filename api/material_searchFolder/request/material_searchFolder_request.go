package material_searchFolder_request

import (
	"doudian.com/open/sdk_golang/api/material_searchFolder/response"
	"doudian.com/open/sdk_golang/core"
	"encoding/json"
)

type MaterialSearchFolderRequest struct {
	doudian_sdk.BaseDoudianOpApiRequest
	Param *MaterialSearchFolderParam 
}
func (c *MaterialSearchFolderRequest) GetUrlPath() string{
	return "/material/searchFolder"
}


func New() *MaterialSearchFolderRequest{
	request := &MaterialSearchFolderRequest{
		Param: &MaterialSearchFolderParam{},
	}
	request.SetConfig(doudian_sdk.GlobalConfig)
	request.SetClient(doudian_sdk.DefaultDoudianOpApiClient)
	return request

}


func (c *MaterialSearchFolderRequest) Execute(accessToken *doudian_sdk.AccessToken) (*material_searchFolder_response.MaterialSearchFolderResponse, error){
	responseJson, err := c.GetClient().Request(c, accessToken)
	if err != nil {
		return nil, err
	}
	response := &material_searchFolder_response.MaterialSearchFolderResponse{}
	_ = json.Unmarshal([]byte(responseJson), response)
	return response, nil

}


func (c *MaterialSearchFolderRequest) GetParamObject() interface{}{
	return c.Param
}


func (c *MaterialSearchFolderRequest) GetParams() *MaterialSearchFolderParam{
	return c.Param
}


type MaterialSearchFolderParam struct {
	// 排序方式 0-创建时间倒序 1-创建时间正序 2-修改时间倒序 3-修改时间正序 4-文件夹名倒序 5-文件夹名正序
	OrderBy *int32 `json:"order_by,omitempty"`
	// 分页的页数，从1开始
	PageNum *int32 `json:"page_num,omitempty"`
	// 每页返回的数量。最大为100，默认为50
	PageSize *int32 `json:"page_size,omitempty"`
	// 需要搜索的文件名片段
	Name string `json:"name,omitempty"`
	// 文件夹id
	FolderId string `json:"folder_id,omitempty"`
	// 创建时间最小值，包含这一秒
	CreateTimeStart string `json:"create_time_start,omitempty"`
	// 创建时间最大值，包含这一秒
	CreateTimeEnd string `json:"create_time_end,omitempty"`
	// 父文件夹id
	ParentFolderId string `json:"parent_folder_id,omitempty"`
	// 文件夹状态。1-有效 4-在回收站中
	OperateStatus []int32 `json:"operate_status,omitempty"`
}
