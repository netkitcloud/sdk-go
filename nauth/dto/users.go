package dto

import "github.com/netkitcloud/sdk-go/common"

type UserListDto struct {
	Data []User
	common.BaseResponse
}
