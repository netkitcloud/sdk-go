package param

import "github.com/netkitcloud/sdk-go/common"

type CreateDepartment struct {
	Name           string `json:"name"`
	Status         bool   `json:"status"`
	OrganizationID uint   `json:"organization_id"`
	Description    string `json:"description,omitempty"`
}

type UpdateDepartment struct {
	Name        string `json:"name"`
	Status      bool   `json:"status"`
	Description string `json:"description,omitempty"`
}

type QueryDepartment struct {
	common.PaginationParams
	Name           string `json:"name,omitempty" form:"name"`
	Status         bool   `json:"status,omitempty" form:"status"`
	OrganizationID uint   `json:"organization_id,omitempty" form:"organization_id"`
}
