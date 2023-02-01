package nauth

import "net/http"

type PhoneCodeLoginDto struct {
	Phone string `json:"phone"`
	Code  string `json:"code"`
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
