package dto

import "github.com/netkitcloud/sdk-go/common"

// 资源操作模型
type Action struct {
	ID           uint   `json:"id"`
	Code         string `json:"code"`
	Description  string `json:"description"`
	ResourceCode string `json:"resource_code"`
	Createdat    string `json:"createdat"`
	Updatedat    string `json:"updatedat"`

	Resource Resource `json:"resource"`
}

type ActionDto struct {
	Data Action
	common.BaseResponse
}

type ListActionDto struct {
	Data []Action
	common.BaseListResponse
}

type VerifyActionDto struct {
	Data struct {
		HasPermission bool `json:"hasPermission"`
	} `json:"data"`
	common.BaseResponse
}
