package param

import "github.com/netkitcloud/sdk-go/common"

type CreateOrganization struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Category    string `json:"category"`
	Description string `json:"description,omitempty"`
}

type UpdateOrganization struct {
	Name        string `json:"name"`
	Email       string `json:"email"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Category    string `json:"category"`
	Description string `json:"description,omitempty"`
}

type QueryOrganization struct {
	common.PaginationParams
	Name     string `json:"name,omitempty" form:"name"`
	Email    string `json:"email,omitempty" form:"email"`
	Country  string `json:"country,omitempty" form:"country"`
	City     string `json:"city,omitempty" form:"city"`
	Category string `json:"category,omitempty" form:"category"`
}
