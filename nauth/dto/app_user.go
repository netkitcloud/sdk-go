package dto

import "github.com/netkitcloud/sdk-go/common"

type AppUserCreateDto struct {
	AppId  string   `json:"appid"`
	UID    []string `json:"uid"`
	Permit int      `json:"permit"`
}

type AppUserUpdateDto AppUserCreateDto

type AppUserDeleteDto struct {
	AppId string   `json:"appid"`
	UID   []string `json:"uid"`
}

type AppUserQuery struct {
	AppId   string `form:"appid" json:"appid"`
	Search  string `form:"search" json:"search"`
	PerPage int    `form:"per_page" json:"per_page"`
	Current int    `form:"current" json:"current"`
}

type AppUser struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	AppId    string `json:"appid"`
	UID      string `json:"uid"`
	Permit   int    `json:"permit"`
}

type ListAppUserDto struct {
	Data []AppUser
	common.BaseResponse
}

type GetAppUserResponseDto struct {
	common.BaseResponse
	Data AppUser
}
