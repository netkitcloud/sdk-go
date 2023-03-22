package nauth

import (
	"github.com/golang-jwt/jwt/v4"
)

type AuthenticationClientOptions struct {
	Host      string
	Tenant    string
	AccessKey string
	AppId     string
	AppSecret string
	RedirectUri string
	TokenEndPointAuthMethod string
	Issuer string
}

type Result struct {
	Code    int    `json:"code,omitempty"`
	Status  bool   `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

type AuthUrlResult struct {
	Url   string
	State string
	Nonce string
}

type AccessTokenClaims struct {
	jwt.StandardClaims
}
