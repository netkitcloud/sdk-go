package nauth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/netkitcloud/sdk-go/nauth/param"
)

// 增加角色
func (c *AuthenticationClient) CreateRole(params *param.CreateRole) (resp dto.RoleDto, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiRole, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 更新角色
func (c *AuthenticationClient) DeleteRole(role_code string) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialRole, role_code)
	body, err := c.SendHttpRequest(uri, http.MethodDelete, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 更新角色
func (c *AuthenticationClient) UpdateRole(role_code string, params *param.UpdateRole) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialRole, role_code)
	body, err := c.SendHttpRequest(uri, http.MethodPut, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取指定角色
func (c *AuthenticationClient) GetRole(role_code string) (resp dto.RoleDto, err error) {
	if role_code == "" {
		err = errors.New("role_code is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialRole, role_code)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取角色列表
func (c *AuthenticationClient) ListRole(pagination param.QueryRole) (resp dto.ListRoleDto, err error) {
	if err = c.validate.Struct(pagination); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiRole, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
