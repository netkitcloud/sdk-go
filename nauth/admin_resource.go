package nauth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/netkitcloud/sdk-go/nauth/param"
)

// 增加资源
func (c *AuthenticationAdmin) CreateResource(params *param.CreateResource) (resp dto.ResourceDto, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiResource, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 删除资源
func (c *AuthenticationAdmin) DeleteResource(resource_id string) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialResource, resource_id)
	body, err := c.SendHttpRequest(uri, http.MethodDelete, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 更新资源
func (c *AuthenticationAdmin) UpdateResource(resource_id string, params *param.UpdateResource) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiSpecialResource, resource_id)
	body, err := c.SendHttpRequest(uri, http.MethodPut, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取指定资源
func (c *AuthenticationAdmin) GetResource(resource_id string) (resp dto.ResourceDto, err error) {
	if resource_id == "" {
		err = errors.New("resource_id is required")
		return
	}

	uri := fmt.Sprintf(apiSpecialResource, resource_id)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取资源列表
func (c *AuthenticationAdmin) ListResource(pagination param.QueryResource) (resp dto.ListResourceDto, err error) {
	if err = c.validate.Struct(pagination); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiResource, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
