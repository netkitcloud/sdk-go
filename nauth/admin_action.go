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
func (c *AuthenticationAdmin) CreateAction(params *param.CreateAction) (resp dto.ActionDto, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiAction, http.MethodPost, params)
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
