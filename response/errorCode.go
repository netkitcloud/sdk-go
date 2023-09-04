// #Title  errorcode.go
// #Description  配置返回的错误结构体和方法
// #Author	rong	2022/02/25	18:16
// #Update
package response

import (
	"fmt"
)

// Error 请求返回错误结构
type Error struct {
	code    int
	message string
}

var codes = map[int]string{}

// #title	NewError
// #description	创建接口标准返回错误
// #author	rong	2022/03/04
// #param	code int 错误码
//
//	message string 错误信息
//
// #return  *Error 接口标准错误指针
func NewError(code int, message string) *Error {
	if _, ok := codes[code]; ok {
		panic(fmt.Sprintf("错误码 %d 已经存在，请更换一个", code))
	}
	codes[code] = message
	return &Error{code: code, message: message}
}

// #title	Code
// #description	获取接口标准错误错误码
// #author	rong	2022/03/04
// #param	无
// #return  int 接口标准错误码
func (e *Error) Code() int {
	return e.code
}

// #title	Msg
// #description	获取接口标准错误错误信息
// #author	rong	2022/03/04
// #param	无
// #return  int 接口标准错误信息
func (e *Error) Msg() string {
	return e.message
}
