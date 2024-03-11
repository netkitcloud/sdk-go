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
		if err != nil {
			return
		}
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
