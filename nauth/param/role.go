package param

import "github.com/netkitcloud/sdk-go/common"

type UpdateRole struct {
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status,omitempty"`
}

type CreateRole struct {
	Code        string `json:"code" validate:"required"`
	Description string `json:"description,omitempty"`
	Status      bool   `json:"status,omitempty"`
}

type QueryRole struct {
	common.PaginationParams
	Code   string `json:"code,omitempty" form:"code"`
	Status bool   `json:"status,omitempty" form:"status"`
}
