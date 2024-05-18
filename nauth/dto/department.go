package dto

import (
	"github.com/netkitcloud/sdk-go/common"
)

// 部门模型
type Department struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	Name           string `json:"name"`
	Description    string `json:"description"`
	Status         bool   `json:"status"`
	Createdat      string `json:"createdat"`
	Updatedat      string `json:"updatedat"`

	Organization Organization `json:"organization"`
	// Users []User `json:"users"`
}

type DepartmentDto struct {
	Data Department
	common.BaseResponse
}

type ListDepartmentDto struct {
	Data []Department
	common.BaseListResponse
}

// 组织用户模型
type DepartmentUser struct {
	ID           uint   `json:"id"`
	DepartmentID uint   `json:"department_id"`
	UserUID      string `json:"user_uid"`
	IsManager    bool   `json:"is_manager"`
	Createdat    string `json:"createdat"`
	Updatedat    string `json:"updatedat"`

	User       User       `json:"user"`
	Department Department `json:"department"`
}
