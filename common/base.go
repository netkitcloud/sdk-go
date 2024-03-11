package common

type BaseResponse struct {
	Code    int    `json:"code"`
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
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
	Current int `json:"current"`
	PerPage int `json:"per_page"`
}
