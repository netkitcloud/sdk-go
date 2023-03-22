package nauth

import (
	"encoding/json"
	"net/http"
)

type PhoneCodeLoginDto struct {
	Phone string `json:"phone"`
	Code  int `json:"code"`
}

type PhonePasswordLoginDto struct {
	Phone string `json:"phone"`
	Password string `json:"password"`
}

type UsernameLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsernameLoginResult struct {
	User
	AccessTokenClaims
}

func (c *AuthenticationClient) UsernameLogin(dto *UsernameLoginDto) (*User, error) {
	body, err := c.SendHttpRequest("/login/username", http.MethodPost, dto)
	if err != nil {
		return nil, err
	}

	return c.responseGetUser(body)
}

func (c *AuthenticationClient) PhoneCodeLogin(dto *PhoneCodeLoginDto) (*User, error) {
	body, err := c.SendHttpRequest("/login/sms", http.MethodPost, dto)
	if err != nil {
		return nil, err
	}

	return c.responseGetUser(body)
}

func (c *AuthenticationClient) PhonePasswordLoginDto(dto *PhonePasswordLoginDto) (*User, error) {
	body, err := c.SendHttpRequest("/login/phone", http.MethodPost, dto)
	if err != nil {
		return nil, err
	}

	return c.responseGetUser(body)
}

// SendLoginPhoneCode
// 发送手机注册验证码
func (c *AuthenticationClient) SendLoginPhoneCode(phone string) (*Result, error) {
	body, err := c.SendHttpRequest("/login/getsms", http.MethodPost, map[string]interface{}{
		"phone": phone,
	})
	if err != nil {
		return nil, err
	}

	var result Result
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
