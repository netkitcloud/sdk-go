package nauth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
)

// 第三方授权登陆
func (c *AuthenticationClient) SocialLogin(tenant, identifier string, login *dto.SocialLoginDto) (sloginResp *dto.SocialLoginRespDto, err error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/login/%s", tenant, identifier)
	resp, err := c.SendHttpRequest(uri, http.MethodPost, login)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(resp, &sloginResp)
	return sloginResp, err
}

// 第三方授权登陆回调
func (c *AuthenticationClient) SocialLoginCallback(tenant, identifier string, login *dto.SocialLoginCallbackDto) (*dto.User, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/login/%s/callback", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodGet, login)
	if err != nil {
		return nil, err
	}

	return c.responseGetUser(body)
}

// 第三方授权绑定用户（需要token）
func (c *AuthenticationClient) SocialBind(tenant, identifier string, bind *dto.SocialBindDto) (bindDto *dto.SocialBindRespDto, err error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/bind/%s", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodPost, bind)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(body, &bindDto)
	return bindDto, err
}

// 第三方授权解绑用户（需要token）
func (c *AuthenticationClient) SocialUnBind(tenant, identifier string, bind *dto.SocialUnBindDto) (*common.BaseResponse, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/unbind/%s", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodPost, bind)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, err
}
