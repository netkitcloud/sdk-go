package nauth

import (
	"encoding/json"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
)

func (c *AuthenticationClient) UsernameLogin(login *dto.UsernameLoginDto) (*dto.User, error) {
	body, err := c.SendHttpRequest("/login/username", http.MethodPost, login)
	if err != nil {
		return nil, err
	}

	return c.responseGetUser(body)
}

func (c *AuthenticationClient) PhoneCodeLogin(login *dto.PhoneCodeLoginDto) (*dto.User, error) {
	body, err := c.SendHttpRequest("/login/sms", http.MethodPost, login)
	if err != nil {
		return nil, err
	}

	return c.responseGetUser(body)
}

func (c *AuthenticationClient) PhonePasswordLoginDto(login *dto.PhonePasswordLoginDto) (*dto.User, error) {
	body, err := c.SendHttpRequest("/login/phone", http.MethodPost, login)
	if err != nil {
		return nil, err
	}

	return c.responseGetUser(body)
}

// SendLoginPhoneCode
// 发送手机注册验证码
func (c *AuthenticationClient) SendLoginPhoneCode(phone string) (*common.BaseResponse, error) {
	body, err := c.SendHttpRequest("/login/getsms", http.MethodPost, map[string]interface{}{
		"phone": phone,
	})
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
