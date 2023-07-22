package nauth

import (
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/nauth/dto"
)

// UpdateUserProfile
// 更新用户信息
func (c *AuthenticationClient) UpdateUserProfile(profile *dto.UpdateUserProfileDto) error {
	if len(c.AccessToken) <= 0 {
		return fmt.Errorf("TokenError")
	}
	body, err := c.SendHttpRequest("/user", http.MethodPut, profile)
	if err != nil {
		return err
	}

	return c.responseError(body)
}

// UpdateUserPassword
// 更新用户密码
func (c *AuthenticationClient) UpdateUserPassword(passwordDto *dto.UpdatePasswordDto) error {
	if len(c.AccessToken) <= 0 {
		return fmt.Errorf("TokenError")
	}
	body, err := c.SendHttpRequest("/user/password", http.MethodPut, passwordDto)
	if err != nil {
		return err
	}

	return c.responseError(body)
}

func (c *AuthenticationClient) GetUser() (*dto.User, error) {
	if len(c.AccessToken) <= 0 {
		return nil, fmt.Errorf("TokenError")
	}

	body, err := c.SendHttpRequest(fmt.Sprintf("/oauth/%s/userinfo", c.options.Tenant), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	return c.responseGetUser(body)
}

func (c *AuthenticationClient) GetUserByToken(token string) (*dto.User, error) {
	c.AccessToken = token
	body, err := c.SendHttpRequest(fmt.Sprintf("/oauth/%s/userinfo", c.options.Tenant), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	return c.responseGetUser(body)
}

// 修改手机号
func (c *AuthenticationClient) UpdatePhone(phoneDto dto.UpdatePhoneDto) error {
	if len(c.AccessToken) <= 0 {
		return fmt.Errorf("TokenError")
	}

	body, err := c.SendHttpRequest("/user/phone", http.MethodPut, phoneDto)
	if err != nil {
		return err
	}

	return c.responseError(body)
}

// 修改手机号
func (c *AuthenticationClient) GetPhoneSMS(phone string) error {
	if len(c.AccessToken) <= 0 {
		return fmt.Errorf("TokenError")
	}

	body, err := c.SendHttpRequest("/user/getsms", http.MethodPost, map[string]interface{}{
		"phone": phone,
	})
	if err != nil {
		return err
	}

	return c.responseError(body)
}

func (c *AuthenticationClient) GetSMS() error {
	if len(c.AccessToken) <= 0 {
		return fmt.Errorf("TokenError")
	}

	body, err := c.SendHttpRequest("/user/getsms", http.MethodGet, nil)
	if err != nil {
		return err
	}

	return c.responseError(body)
}