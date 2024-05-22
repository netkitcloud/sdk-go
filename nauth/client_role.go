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
func (c *AuthenticationClient) DeleteRole(role_id string) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialRole, role_id)
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
func (c *AuthenticationClient) UpdateRole(role_id string, params *param.UpdateRole) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialRole, role_id)
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
func (c *AuthenticationClient) GetRole(role_id string) (resp dto.RoleDto, err error) {
	if role_id == "" {
		err = errors.New("role_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialRole, role_id)
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

// 用户绑定角色
func (c *AuthenticationClient) UserBindRole(params *param.UserBindRoleForm) (resp common.BaseResponse, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiUserRole, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 用户解绑角色
func (c *AuthenticationClient) UserUnbindRole(params *param.UserBindRoleForm) (resp common.BaseResponse, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiUserRole, http.MethodDelete, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
