package dto

import "github.com/netkitcloud/sdk-go/common"

// 资源模型
type Resource struct {
	ID          uint   `json:"id"`
	Code        string `json:"code"`
	Type        string `json:"status"`
	Description string `json:"description"`
	Createdat   string `json:"createdat"`
	Updatedat   string `json:"updatedat"`

	Actions   []Action `json:"actions"`
	ActionArr []string `json:"action_arr"` // [map[data1:[read write] data2:[read]]]
}

type ResourceDto struct {
	Data Resource
	common.BaseResponse
}

type ListResourceDto struct {
	Data []Resource
	common.BaseListResponse
}
