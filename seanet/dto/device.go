package dto

import (
	"time"

	"github.com/netkitcloud/sdk-go/common"
)

type Device struct {
	Id                     int                      `json:"id"`
	Sn                     string                   `json:"sn"`
	Name                   string                   `json:"name"`
	Status                 int                      `json:"status"`
	Login                  int                      `json:"login"`
	Online                 int                      `json:"online"`
	FwVersion              string                   `json:"fw_version"`
	HwVersion              string                   `json:"hw_version"`
	ProductKey             string                   `json:"product_key"`
	UpdatedAt              *time.Time               `json:"updatedat"`
	CreatedAt              *time.Time               `json:"createdat"`
	IotPlatformParam       map[string]interface{}   `json:"-"`
	Config                 map[string]interface{}   `json:"config"`
	ExpectConfig           map[string]interface{}   `json:"expect_config"`
	ConfigVersion          int                      `json:"config_version"`
	ExpectConfigVersion    int                      `json:"expect_config_version"`
	ConfigVersionDifferent int                      `json:"config_version_different"`
	Activate               bool                     `json:"activate"`
	ActivatedAt            *time.Time               `json:"activatedat"`
	Lastcommunicationat    *time.Time               `json:"lastcommunicationat"`
	Address                string                   `json:"address"`
	Latlng                 string                   `json:"latlng"`
	HwinfoGroupId          string                   `json:"-"`
	ProductName            string                   `json:"product_name"`
	CfgTemplate            map[string]interface{}   `json:"cfg_template"`
	CfgFeature             []map[string]interface{} `json:"feature"`
	Iccid                  string                   `json:"iccid"`
	IsPlatformIccid        bool                     `json:"is_platform_iccid"`
}

type DeviceDto struct {
	Data Device
	common.BaseResponse
}

type ListDeviceDto struct {
	Data []Device
	common.BaseResponse
}

type DeviceLog struct {
	Sn         string      `json:"sn"`
	RecordTime string      `json:"record_time"`
	Msg        interface{} `json:"msg"`
}

type DeviceLogDto struct {
	common.BaseListResponse
	Data []DeviceLog
}

type SubDevice struct {
	Id            string `json:"id"`
	CreatedAt     string `json:"createdat"`
	GatewayRuleId string `json:"gateway_rule_id"`
	Devicekey     string `json:"devicekey"`
}

type SubDevices struct {
	Data []SubDevice
	common.BaseListResponse
}

// 向指定主题发送消息内容
type CmdDeviceDto struct {
	Topic   string `json:"topic"`
	Content string `json:"content" validate:"required"` // 多条指令，按顺序发送给设备
}

type Property struct {
	Ts  string `json:"ts"`
	Val any    `json:"val"`
}
type GetPropertyDto struct {
	Data []Property
	common.BaseListResponse
}
