package nauth

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
)

// 第三方微信公众号向用户发送模板消息
func (c *AuthenticationClient) SocialWxOfficeSendMsg(tenant, identifier string, msg *dto.SocialWxOfficeTemplateMsgDto) (*common.BaseResponse, error) {
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
func (c *AuthenticationClient) SocialWxUserList(tenant, identifier string, query *dto.WxUserQueryDto) (*dto.WxUser, error) {
	if tenant == "" || identifier == "" {
		return nil, fmt.Errorf("tenant and identifier are required")
	}

	uri := fmt.Sprintf("/%s/social/wxuser/%s", tenant, identifier)
	body, err := c.SendHttpRequest(uri, http.MethodGet, query)
	if err != nil {
		return nil, err
	}

	var result dto.WxUser
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}