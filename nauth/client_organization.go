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
func (c *AuthenticationClient) CreateOrganization(params *param.CreateOrganization) (resp dto.OrganizationDto, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiOrganization, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 删除组织
func (c *AuthenticationClient) DeleteOrganization(organization_id string) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialOrganization, organization_id)
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
func (c *AuthenticationClient) UpdateOrganization(organization_id string, params *param.UpdateOrganization) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialOrganization, organization_id)
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
func (c *AuthenticationClient) GetOrganization(organization_id string) (resp dto.OrganizationDto, err error) {
	if organization_id == "" {
		err = errors.New("organization_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialOrganization, organization_id)
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
func (c *AuthenticationClient) ListOrganization(pagination param.QueryOrganization) (resp dto.ListOrganizationDto, err error) {
	if err = c.validate.Struct(pagination); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiOrganization, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 在组织下添加成员
func (c *AuthenticationClient) OrgAddMember(organization_id string, params *dto.AddUserDto) (resp dto.AddOrganizationMemberDto, err error) {
	if organization_id == "" {
		err = errors.New("organization_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialOrganizationMember, organization_id)
	body, err := c.SendHttpRequest(uri, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 在组织下绑定成员
func (c *AuthenticationClient) OrgBindMember(organization_id string, params *param.OrganizationUser) (resp common.BaseResponse, err error) {
	if organization_id == "" {
		err = errors.New("organization_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialOrganizationMemberBind, organization_id)
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
func (c *AuthenticationClient) OrgUnbindMember(organization_id string, params *param.OrganizationUser) (resp common.BaseResponse, err error) {
	if organization_id == "" {
		err = errors.New("organization_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialOrganizationMemberUnbind, organization_id)
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
func (c *AuthenticationClient) GetOrganizationUsers(organization_id string, pagination common.PaginationParams) (resp dto.ListOrganizationUserDto, err error) {
	if organization_id == "" {
		err = errors.New("organization_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialOrganizationMember, organization_id)
	body, err := c.SendHttpRequest(uri, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 设置指定组织下用户是否为组织管理员
func (c *AuthenticationClient) SetOrganizationUsersIsmanage(organization_id string,
	params *param.OrganizationUserManagerForm) (resp common.BaseResponse, err error) {
	if organization_id == "" {
		err = errors.New("organization_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialOrganizationMemberManager, organization_id)
	body, err := c.SendHttpRequest(uri, http.MethodPut, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
