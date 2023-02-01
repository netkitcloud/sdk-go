package nauth

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/valyala/fastjson"
)

type PhoneCodeRequestDtto struct {
	Phone string `json:"phone"`
}

type PhoneCodeRegisterDto struct {
	Phone string `json:"phone"`
	Code  int `json:"code"`
}

type UsernameRegisterDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// RegisterByUsername
// 使用用户名密码注册
func (c *AuthenticationClient) RegisterByUsername(dto *UsernameRegisterDto) error {
	body, err := c.SendHttpRequest("/register/username", http.MethodPost, dto)
	if err != nil {
		return err
	}

	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return err
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return err
		}
		return errors.New(string(msg))
	}

	return nil
}

// RegisterByPhoneCode
// 使用用户名密码注册
func (c *AuthenticationClient) RegisterByPhoneCode(dto *PhoneCodeRegisterDto) error {
	body, err := c.SendHttpRequest("/register/phone", http.MethodPost, dto)
	if err != nil {
		return err
	}

	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return err
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return err
		}
		return errors.New(string(msg))
	}

	return nil
}

// SendRegisterPhoneCode
// 发送手机注册验证码
func (c *AuthenticationClient) SendRegisterPhoneCode(phone string) (*Result, error) {
	body, err := c.SendHttpRequest("/register/getsms", http.MethodPost, map[string]interface{}{
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
