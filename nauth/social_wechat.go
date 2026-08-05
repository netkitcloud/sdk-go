package nauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/valyala/fastjson"
)

// SocialWxOfficeAuthURL 获取公众号网页授权登录 URL。
func (c *AuthenticationClient) SocialWxOfficeAuthURL(tenant, identifier string) (*dto.WxOfficeAuthURLRespDto, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/login/%s", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodPost, nil)
	if err != nil {
		return nil, err
	}

	var result dto.WxOfficeAuthURLRespDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if !result.Status {
		return &result, fmt.Errorf("code : %d", result.Code)
	}
	return &result, nil
}

// SocialWxOfficeLogin 用公众号网页授权回调的 code/state 换取 NAuth 登录结果。
func (c *AuthenticationClient) SocialWxOfficeLogin(tenant, identifier string, callback *dto.WxOfficeCallbackDto) (*dto.User, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/login/%s/callback", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodGet, callback)
	if err != nil {
		return nil, err
	}
	return c.responseGetUser(body)
}

// SocialWxOfficeBindAuthURL 获取当前登录用户绑定公众号账号的授权 URL。
func (c *AuthenticationClient) SocialWxOfficeBindAuthURL(tenant, identifier string) (*dto.WxOfficeAuthURLRespDto, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/bind/%s", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodPost, nil)
	if err != nil {
		return nil, err
	}

	var result dto.WxOfficeAuthURLRespDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if !result.Status {
		return &result, fmt.Errorf("code : %d", result.Code)
	}
	return &result, nil
}

// SocialWxOfficeBindCallback 提交公众号绑定回调 code/state。
func (c *AuthenticationClient) SocialWxOfficeBindCallback(tenant, identifier string, callback *dto.WxOfficeCallbackDto) (*common.BaseResponse, error) {
	return c.socialWxOfficeStateCallback(tenant, identifier, "bind", callback)
}

// SocialWxOfficeReBindAuthURL 获取当前登录用户换绑公众号账号的授权 URL。
func (c *AuthenticationClient) SocialWxOfficeReBindAuthURL(tenant, identifier string) (*dto.WxOfficeAuthURLRespDto, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/rebind/%s", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodPost, nil)
	if err != nil {
		return nil, err
	}

	var result dto.WxOfficeAuthURLRespDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if !result.Status {
		return &result, fmt.Errorf("code : %d", result.Code)
	}
	return &result, nil
}

// SocialWxOfficeReBindCallback 提交公众号换绑回调 code/state。
func (c *AuthenticationClient) SocialWxOfficeReBindCallback(tenant, identifier string, callback *dto.WxOfficeCallbackDto) (*common.BaseResponse, error) {
	return c.socialWxOfficeStateCallback(tenant, identifier, "rebind", callback)
}

func (c *AuthenticationClient) socialWxOfficeStateCallback(tenant, identifier, action string, callback *dto.WxOfficeCallbackDto) (*common.BaseResponse, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/%s/%s/callback", tenant, action, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodGet, callback)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}
	if !result.Status {
		return &result, fmt.Errorf("code : %d", result.Code)
	}
	return &result, nil
}

// 第三方微信公众号向用户发送模板消息
func (c *AuthenticationAdmin) SocialWxOfficeSendMsg(tenant, identifier string, msg *dto.SocialWxOfficeTemplateMsgDto) (*common.BaseResponse, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/weixin/%s/%s/sendmsg", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodPost, msg)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// 获取绑定的微信公众号用户信息
func (c *AuthenticationClient) SocialWxUserList(tenant, identifier string, query *dto.WxUserQueryDto) (wxUsersListResp *dto.ListWxUserDto, err error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/wxuser/%s", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		err = errors.New(string(msg))
		return
	}

	err = json.Unmarshal(body, &wxUsersListResp)
	if err != nil {
		return
	}

	if !wxUsersListResp.Status {
		err = fmt.Errorf("code : %d", wxUsersListResp.Code)
	}

	return
}
