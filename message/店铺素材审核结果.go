package doudian_message

import (
	"encoding/json"
)

// 店铺素材审核结果
// 文档：https://op.jinritemai.com/docs/message-docs/1165/2674

func init() {
	MessageDataParseHandle.RegisterMessage(MaterialAuditResultForShop{})
}

type MaterialAuditResultForShop struct {
	AuditStatus     int    `json:"audit_status"`
	AuditStatusDesc string `json:"audit_status_desc"`
	ByteUrl         string `json:"byte_url"`
	CreateTime      string `json:"create_time"`
	DeleteTime      string `json:"delete_time"`
	FolderId        string `json:"folder_id"`
	MaterialId      string `json:"material_id"`
	MaterialType    string `json:"material_type"`
	Name            string `json:"name"`
	OperateStatus   int    `json:"operate_status"`
	OriginUrl       string `json:"origin_url"`
	PhotoInfo       struct {
		Format string `json:"format"`
		Height int    `json:"height"`
		Width  int    `json:"width"`
	} `json:"photo_info"`
	ShopId     int         `json:"shop_id"`
	Size       int         `json:"size"`
	UpdateTime string      `json:"update_time"`
	VideoInfo  interface{} `json:"video_info"`
}

func (m MaterialAuditResultForShop) Tag() string {
	return "137"
}

func (m MaterialAuditResultForShop) Parse(data []byte) (MessageData, error) {
	var tmp MaterialAuditResultForShop
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
