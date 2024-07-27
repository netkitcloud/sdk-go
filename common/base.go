package common

type BaseResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}
type BaseDataResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Data    any    `json:"data,omitempty"`
}

type BaseListResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
	Total   int    `json:"total,omitempty"`
	Current int    `json:"current,omitempty"`
	Perpage int    `json:"per_page,omitempty"`
}

type PaginationParams struct {
	Current int `json:"current" default:"1" validate:"omitempty"`
	PerPage int `json:"per_page" default:"20" validate:"omitempty"`
}
