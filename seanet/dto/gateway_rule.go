package dto

import "github.com/netkitcloud/sdk-go/common"

type GatewayRule struct {
	Id        string `json:"id"`        // gateway_rule_id
	DeviceKey string `json:"devicekey"` // gateway devicekey
	Name      string `json:"name"`      // gateway_rule_name
	CreatedAt string `json:"createdat,omitempty"`
	UpdatedAt string `json:"updatedat,omitempty"`
}

type GatewayRulesDto struct {
	Data []GatewayRule
	common.BaseListResponse
}


type GatewayRuleDevice struct {
	Name                string `json:"name"`
	CreatedAt           string `json:"createdat"`
	Devicekey           string `json:"devicekey"`
	GatewayRuleDeviceID string `json:"id"`
}

type GatewayRuleDevicesDto struct {
	Data []GatewayRuleDevice `json:"data"`
	common.BaseListResponse
}
