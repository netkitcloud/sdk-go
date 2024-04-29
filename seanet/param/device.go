package param

import "github.com/netkitcloud/sdk-go/common"

type UpdateDevice struct {
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

type CreateDevice struct {
	ProductKey       string                 `json:"product_key" validate:"required"`
	Sn               []string               `json:"sn" validate:"required"`
	IotPlatformParam map[string]interface{} `json:"iot_platform_param"`
	Name             string                 `json:"name"`
	HwinfoGroupId    string                 `json:"hwinfo_group_id"`
	TaskId           string                 `json:"task_id"`
	HwVersion        string                 `json:"hw_version"`
}

type QueryDevice struct {
	common.PaginationParams
	ProductKey string `json:"product_key"`
	Name       string `json:"name" form:"name"`
	Sn         string `json:"sn" form:"sn"`
	Online     *int   `json:"online" form:"online"`
	IsStore    *bool  `json:"is_store" form:"is_store"`
	Search     string `json:"search" form:"search"`
}

// 控制网关设备Form
type ControlGateway struct {
	Sn      string `json:"sn"`
	Cmd     string `json:"cmd"`
	Content string `json:"content"`
}

type SwitchDevice struct {
	Devicekey string   `json:"devicekey"`
	Contents  []string `json:"contents"` // 多条指令，按顺序发送给设备
}

type CreateGatewayRule struct {
	GatewayDeviceKey string `json:"devicekey"`
}

type DeleteGatewayRule struct {
	RuleIDs string `json:"ids"`
}

type CreateGatewayRuleDevices struct {
	GatewayRuleDeviceKeys []string `json:"devicekey"`
}

type DeleteGatewayRuleDevices struct {
	GatewayRuleDeviceKeys []string `json:"ids"`
}
