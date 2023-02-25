package doudian_message

import (
	"encoding/json"
)

// 供应商素材审核结果
// 文档：https://op.jinritemai.com/docs/message-docs/1165/2675

func init() {
	MessageDataParseHandle.RegisterMessage(MaterialAuditResultForBSCP{})
}

type MaterialAuditResultForBSCP struct {
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
	Size       int    `json:"size"`
	SupplierId string `json:"supplier_id"`
	UpdateTime string `json:"update_time"`
	VideoInfo  struct {
		Duration interface{} `json:"duration"`
		Format   string      `json:"format"`
		Height   int         `json:"height"`
		Vid      string      `json:"vid"`
		Width    int         `json:"width"`
	} `json:"video_info"`
}

func (m MaterialAuditResultForBSCP) Tag() string {
	return "138"
}

func (m MaterialAuditResultForBSCP) Parse(data []byte) (MessageData, error) {
	var tmp MaterialAuditResultForBSCP
	err := json.Unmarshal(data, &tmp)
	return tmp, err
}
