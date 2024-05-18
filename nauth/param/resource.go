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
