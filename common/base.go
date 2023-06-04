package common

type BaseResponse struct {
	Code    int    `json:"code,omitempty"`
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type PaginationParams struct {
	Current int `json:"current"`
	PerPage int `json:"per_page"`
}
