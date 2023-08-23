package dto

import (
	"time"

	"github.com/netkitcloud/sdk-go/common"
)

type CreateRuleForwardDto struct {
	Name        string `json:"name" validate:"required"`
	Enable      int    `json:"enable"`
	Productkey  string `json:"productkey"`
	Devicekey   string `json:"devicekey"`
	Sourceid    int    `json:"sourceid" validate:"required"`
	Targetid    int    `json:"targetid" validate:"required"`
	Target      string `json:"target" validate:"required"`
	Description string `json:"description"`
	ProjectKey  string
}

type UpdateRuleForwardDto struct {
	ID          string `json:"ruleid" validate:"required"`
	Name        string `json:"name"`
	Enable      *int   `json:"enable"`
	Description string `json:"description"`
	Target      string `json:"target"`
}

type RuleForward struct {
	Id          string     `json:"ruleid"` // 规则ID
	Name        string     `json:"name"`   // 规则名
	SourceId    int        `json:"sourceid"`
	TargetId    int        `json:"targetid"`
	SourceName  string     `json:"rule_source_name"`    // 规则源
	TargetName  string     `json:"rule_target_name"`    // 规则名
	Target      string     `json:"target"`              // 规则目标
	Enable      int        `json:"enable"`              // 规则启用状态
	Description string     `json:"description"`         // 规则描述
	UpdatedAt   *time.Time `json:"updatedat,omitempty"` // 规则更新时间
	CreatedAt   *time.Time `json:"createdat,omitempty"` // 规则创建时间
	ProductKey  string     `json:"productkey"`          // 规则关联的产品ID
	ProductName string     `json:"product_name"`        // 规则关联的产品名
	DeviceKey   string     `json:"devicekey"`           // 规则关联的设备ID
	DeviceName  string     `json:"device_name"`         // 规则关联的设备ID
	ProjectKey  string     `json:"projectkey"`
	Topic       string     `json:"topic"`
}

type ListRuleForwardDto struct {
	Data []RuleForward
	common.BaseResponse
}
