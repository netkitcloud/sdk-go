package param

import "github.com/netkitcloud/sdk-go/common"

type UpdateRole struct {
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status,omitempty"`
	ActionIDs   []uint `json:"action_ids"` // 更新授权资源列表（覆盖之前的权限信息）
}

type CreateRole struct {
	Code        string `json:"code" validate:"required"`
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status,omitempty"`
	ActionIDs   []uint `json:"action_ids"` // 授权资源列表
}

type QueryRole struct {
	common.PaginationParams
	Code   string `json:"code,omitempty" form:"code"`
	Status bool   `json:"status,omitempty" form:"status"`
}

// UserBindRoleForm 用户绑定/解绑角色表单
type UserBindRoleForm struct {
	UserUIDs []string `json:"user_uids" validate:"required"`
	RoleCode string   `json:"role_code" validate:"required"`
}

// UpdateUserRoleForm 更新用户角色表单（覆盖）
type UpdateUserRoleForm struct {
	UserUID   string   `json:"user_uid" validate:"required"`
	RoleCodes []string `json:"role_codes" validate:"required"`
}
