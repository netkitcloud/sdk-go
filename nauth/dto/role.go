package dto

import "github.com/netkitcloud/sdk-go/common"

type Role struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Description string `json:"description"`
	Status      bool   `json:"status"`
	Createdat   string `json:"createdat"`
	Updatedat   string `json:"updatedat"`

	ActionArr []map[string][]string `json:"action_arr"` // [map[data1:[read write] data2:[read]]]
}

type RoleDto struct {
	Data Role
	common.BaseResponse
}

type ListRoleDto struct {
	Data []Role
	common.BaseListResponse
}
