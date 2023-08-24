package dto

import (
	"github.com/netkitcloud/sdk-go/common"
)

type AccessKey struct {
	UID       string `json:"uid,omitempty"`
	AccessKey string `json:"access_key,omitempty"`
	Secret    string `json:"access_secret,omitempty"`
	Comment   string `json:"comment,omitempty"`
	UpdatedAt string `json:"updatedat,omitempty"`
	CreatedAt string `json:"createdat,omitempty"`
}

type AddAccessKeyDto struct {
	Comment string `json:"comment,omitempty"`
}

type UpdateAccessKeyDto struct {
	AccessKey string `json:"access_key,omitempty"`
	Comment   string `json:"comment,omitempty"`
}

type ResetAccessSecretDto struct {
	Data string
	common.BaseResponse
}

type AccessKeyResponseDto struct {
	common.BaseResponse
	Data AccessKey
}

type ListAccessKeyDto struct {
	common.BaseResponse
	Data []AccessKey
}
