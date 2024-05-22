package nauth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/netkitcloud/sdk-go/nauth/param"
)

// 增加资源操作
func (c *AuthenticationAdmin) CreateAction(resource_id string, params *param.CreateAction) (resp dto.ActionDto, err error) {
	if resource_id == "" {
		err = errors.New("resource_id is required")
		return
	}

	if err = c.validate.Struct(params); err != nil {
		return
	}

	uri := fmt.Sprintf(apiSpecialResourceAction, resource_id)
	body, err := c.SendHttpRequest(uri, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 删除资源操作
func (c *AuthenticationAdmin) DeleteAction(action_id string) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialAction, action_id)
	body, err := c.SendHttpRequest(uri, http.MethodDelete, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 更新资源操作
func (c *AuthenticationAdmin) UpdateAction(action_id string, params *param.UpdateAction) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialAction, action_id)
	body, err := c.SendHttpRequest(uri, http.MethodPut, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取指定资源操作
func (c *AuthenticationAdmin) GetAction(action_id string) (resp dto.ActionDto, err error) {
	if action_id == "" {
		err = errors.New("action_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialAction, action_id)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取指定资源下所有操作信息
func (c *AuthenticationAdmin) ListResourceActions(resource_id string, pagination param.QueryResource) (resp dto.ListResourceDto, err error) {
	if resource_id == "" {
		err = errors.New("resource_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialResourceAction, resource_id)
	body, err := c.SendHttpRequest(uri, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取所有资源下所有操作信息
func (c *AuthenticationAdmin) ListAllAction(pagination param.QueryResource) (resp dto.ListResourceDto, err error) {
	body, err := c.SendHttpRequest(apiAction, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
