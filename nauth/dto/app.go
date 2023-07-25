package dto

import "github.com/netkitcloud/sdk-go/common"

type AddAppDto struct {
	Name   string   `json:"name"`
	Domain string   `json:"domain"`
	Scope  []string `json:"scope"`
}

type AppConfig struct {
	DefaultAuthorization bool `json:"default_authorization"`
	DisableLogin         bool `json:"disable_login"`
	DisableRegister      bool `json:"disable_register"`
}

type UpdateAppDto struct {
	Id     string    `json:"id"`
	Name   string    `json:"name"`
	Domain string    `json:"domain"`
	Scope  []string  `json:"scope"`
	Config AppConfig `json:"config"`
}

type App struct {
	AppId  string    `json:"appid"`
	Name   string    `json:"name"`
	Domain string    `json:"domain"`
	Scope  []string  `json:"scope"`
	Config AppConfig `json:"config"`
}

type ListAppDto struct {
	Data []App
	common.BaseResponse
}

type GetAppResponseDto struct {
	common.BaseResponse
	Data App
}
