package nauth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/netkitcloud/sdk-go/nauth/param"
)

// 增加组织
func (c *AuthenticationClient) CreateDepartment(params *param.CreateDepartment) (resp dto.DepartmentDto, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiDepartment, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 删除组织
func (c *AuthenticationClient) DeleteDepartment(department_id string) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialDepartment, department_id)
	body, err := c.SendHttpRequest(uri, http.MethodDelete, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 更新组织
func (c *AuthenticationClient) UpdateDepartment(department_id string, params *param.UpdateDepartment) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialDepartment, department_id)
	body, err := c.SendHttpRequest(uri, http.MethodPut, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取指定组织
func (c *AuthenticationClient) GetDepartment(department_id string) (resp dto.DepartmentDto, err error) {
	if department_id == "" {
		err = errors.New("department_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialDepartment, department_id)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取组织列表
func (c *AuthenticationClient) ListDepartment(pagination param.QueryDepartment) (resp dto.ListDepartmentDto, err error) {
	if err = c.validate.Struct(pagination); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiDepartment, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 在组织下绑定成员
func (c *AuthenticationClient) DepBindMember(department_id string, params *param.DepartmentUser) (resp common.BaseResponse, err error) {
	if department_id == "" {
		err = errors.New("department_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialDepartmentMemberBind, department_id)
	body, err := c.SendHttpRequest(uri, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 在组织下解绑成员
func (c *AuthenticationClient) DepUnbindMember(department_id string, params *param.DepartmentUser) (resp common.BaseResponse, err error) {
	if department_id == "" {
		err = errors.New("department_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialDepartmentMemberUnbind, department_id)
	body, err := c.SendHttpRequest(uri, http.MethodDelete, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取指定组织下所有用户信息
func (c *AuthenticationClient) GetDepartmentUsers(department_id string, pagination common.PaginationParams) (resp dto.ListDepartmentUserDto, err error) {
	if department_id == "" {
		err = errors.New("department_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialDepartmentMember, department_id)
	body, err := c.SendHttpRequest(uri, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 设置指定部门下用户是否为部门管理员
func (c *AuthenticationClient) SetDepartmentUsersIsmanage(department_id string,
	params *param.DepartmentUserManagerForm) (resp common.BaseResponse, err error) {
	if department_id == "" {
		err = errors.New("department_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialDepartmentMemberManager, department_id)
	body, err := c.SendHttpRequest(uri, http.MethodPut, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
