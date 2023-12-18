package dto

import (
	"time"

	"github.com/netkitcloud/sdk-go/common"
)

type CreateDeviceDto struct {
	ProductKey       string                 `json:"product_key" validate:"required"`
	Sn               []string               `json:"sn" validate:"required"`
	IotPlatformParam map[string]interface{} `json:"iot_platform_param"`
	Name             string                 `json:"name"`
	HwinfoGroupId    string                 `json:"hwinfo_group_id"`
	TaskId           string                 `json:"task_id"`
	HwVersion        string                 `json:"hw_version"`
}

type UpdateDeviceDto struct {
	Sn                     string      `json:"sn" validate:"required"`
	Name                   string      `json:"name"`
	Status                 int         `json:"status"`
	Login                  int         `json:"login"`
	Online                 int         `json:"online"`
	HwVersion              string      `json:"hw_version"`
	ProductKey             string      `json:"product_key"`
	IotPlatformParam       interface{} `json:"iot_platform_param"`
	ConfigVersionDifferent *uint32     `json:"config_version_different"`
}

type QueryDeviceDto struct {
	common.PaginationParams
	ProductKey string `json:"product_key"`
	Name       string `json:"name" form:"name"`
	Sn         string `json:"sn" form:"sn"`
	Online     *int   `json:"online" form:"online"`
	IsStore    *bool  `json:"is_store" form:"is_store"`
	Search     string `json:"search" form:"search"`
}

type SwitchDeviceDto struct {
	Devicekey string   `json:"devicekey"`
	Contents  []string `json:"contents"` // 多条指令，按顺序发送给设备
}

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

type ListDeviceDto struct {
	Data []Device
	common.BaseResponse
}
