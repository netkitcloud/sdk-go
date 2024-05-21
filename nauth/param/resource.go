package param

import "github.com/netkitcloud/sdk-go/common"

type CreateResource struct {
	Code        string         `json:"code" validate:"required"`
	Type        string         `json:"type" validate:"required"`
	Description string         `json:"description" binding:"omitempty"`
	Actions     []CreateAction `json:"actions" binding:"omitempty"`
}

type UpdateResource struct {
	Type        string `json:"type"`
	Description string `json:"description" bind:"omitempty"`
}

type QueryResource struct {
	common.PaginationParams
	Code string `json:"code" binding:"omitempty" form:"code"`
	Type string `json:"type" binding:"omitempty" form:"type"`
}

// 权限查询表单
type QueryVerifyAction struct {
	ResourceCode string `json:"resource_code" form:"resource_code" validate:"required"`
	ActionCode   string `json:"action_code" form:"action_code" validate:"required"`
	UID          string `json:"uid" form:"uid" validate:"required"`
	TenantID     string `json:"tenant_id" form:"tenant_id"`
}

type TargetType string

const (
	TargetTypeUser       TargetType = "USER"
	TargetTypeRole       TargetType = "ROLE"
	TargetTypeDepartment TargetType = "DEPARTMENT"
)

type ResourceItemDto struct {
	Code         string   `json:"code"`          // 资源code, eg: device/product
	ActionCodes  []string `json:"action_codes"`  // 资源操作, eg: create, update, delete
	ResourceType string   `json:"resource_type"` // 资源类型，如数据、API、按钮
}

type AuthorizeResourceItem struct {
	TargetType        TargetType        `json:"target_type"`        // 授权类型：用户、角色、部门
	TargetIdentifiers []string          `json:"target_identifiers"` // 授权的标识符：用户uid、角色code、部门code
	Resources         []ResourceItemDto `json:"resources"`          // 授权资源列表
}

// 资源操作授权表单
type AuthorizeResourcesDto struct {
	List []AuthorizeResourceItem `json:"list"`
}
