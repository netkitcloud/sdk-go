package param

import "github.com/netkitcloud/sdk-go/common"

type CreateAction struct {
	Code         string `json:"code" validate:"required"`
	Description  string `json:"description,omitempty"`
}

type UpdateAction struct {
	Description string `json:"description,omitempty"`
}

type QueryAction struct {
	common.PaginationParams
	Code         string `json:"code,omitempty" form:"code"`
	ResourceCode string `json:"resource_code,omitempty" form:"resource_code"`
}
