package nauth

import (
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/nauth/dto"
)

// UpdateUserProfile
// 更新用户信息
func (c *AuthenticationClient) UpdateUserProfile(profile *dto.UpdateUserProfileDto) error {
	body, err := c.SendHttpRequest("/user", http.MethodPut, profile)
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

func (c *AuthenticationClient) GetUserByToken(token string) (*dto.User, error) {
	c.AccessToken = token
	body, err := c.SendHttpRequest(fmt.Sprintf("/oauth/%s/userinfo", c.options.Tenant), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}
	return c.responseGetUser(body)
}
