package response

import (
	"encoding/json"
)

type ApiResponse struct {
	ID      string      `json:"id,omitempty"`       // 当前请求的唯一ID，便于问题定位，忽略也可以
	Status  bool        `json:"status"`             // 请求是否成功
	Code    int         `json:"code"`               // 业务编码
	Message string      `json:"message,omitempty"`  // 错误描述
	Total   *int        `json:"total,omitempty"`    // 总数量
	Current int         `json:"current,omitempty"`  // 当前页码
	PerPage int         `json:"per_page,omitempty"` // 每页数量
	Size    *int        `json:"size,omitempty"`     // 当前页数量
	Data    interface{} `json:"data,omitempty"`     // 成功时返回的数据
}

// #title	ResponseSuccess
// #description	返回成功JSON信息
// #author	feng	2022/02/27	09:16
// #return *ApiResponse interface 标准返回类
func ResponseSuccess() *ApiResponse {
	res := &ApiResponse{
		Status: true,
		Code:   0,
	}
	return res
}

// #title	NewResponse
// #description	返回错误的JSON信息
// #author	feng	2022/02/27	09:16
// #param	err interface 错误响应类
// #return *ApiResponse interface 标准返回类
func NewResponse(err *Error) *ApiResponse {
	status := false
	if err.Code() == 0 {
		status = true
	}

	res := &ApiResponse{
		Status:  status,
		Code:    err.Code(),
		Message: err.Msg(),
	}
	return res
}

// #title	NewResponseData
// #description	返回错误信息并包含data的JSON信息,一般用于展示列表和详情
// #author	feng	2022/02/28
// #param	err interface 错误响应类
//
//	data interface{} 展示的json数据
//
// #return *ApiResponse interface 标准返回类
func NewResponseData(err *Error, data interface{}) *ApiResponse {
	status := false
	if err.Code() == 0 {
		status = true
	}
	return &ApiResponse{
		Status:  status,
		Code:    err.Code(),
		Data:    data,
		Message: err.Msg(),
	}
}

// #title	NewResponseMessage
// #description	返回错误的JSON信息，自带错误提示msg包含具体错误信息,
// #author	feng	2022/02/27	09:16
// #param	err interface 错误响应类
//
//	message string 具体错误信息
//
// #return *ApiResponse interface 标准返回类
func NewResponseMessage(err *Error, message string) *ApiResponse {
	status := false
	if err.Code() == 0 {
		status = true
	}
	return &ApiResponse{
		Status:  status,
		Code:    err.Code(),
		Message: err.Msg() + " : " + message,
	}
}

func (res *ApiResponse) WithData(data interface{}) *ApiResponse {
	res.Data = data
	return res
}

func (res *ApiResponse) WithID(id string) *ApiResponse {
	res.ID = id
	return res
}

func (res *ApiResponse) WithMessage(msg string) *ApiResponse {
	res.Message = msg
	return res
}

func (res *ApiResponse) WithTotal(total int) *ApiResponse {
	res.Total = &total
	return res
}

func (res *ApiResponse) WithSize(size int) *ApiResponse {
	res.Size = &size
	return res
}

func (res *ApiResponse) WithCurrent(current int) *ApiResponse {
	res.Current = current
	return res
}

func (res *ApiResponse) WithPerPage(perPage int) *ApiResponse {
	res.PerPage = perPage
	return res
}

// ToString 返回 JSON 格式的错误详情
func (res *ApiResponse) ToString() string {
	err := &ApiResponse{
		Code:    res.Code,
		Message: res.Message,
		Data:    res.Data,
		ID:      res.ID,
	}

	raw, _ := json.Marshal(err)
	return string(raw)
}
