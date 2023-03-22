package nauth

import (
	"net/http"
)

type UpdateUserProfileDto struct {
	Nickname string `json:"nickname"`
}

// UpdateUserProfile
// 更新用户信息
func (c *AuthenticationClient) UpdateUserProfile(dto *UpdateUserProfileDto) error {
	body, err := c.SendHttpRequest("/user", http.MethodPut, dto)
	if err != nil {
		return err
	}

	return c.responseError(body)
}

// UpdateUserPassword
// 更新用户密码
func (c *AuthenticationClient) UpdateUserPassword(password string) error {
	body, err := c.SendHttpRequest("/user/password", http.MethodPut, map[string]string{
		"password": password,
	})
	if err != nil {
		return err
	}

	return c.responseError(body)
}

