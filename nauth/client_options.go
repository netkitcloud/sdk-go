package nauth

import (
	"github.com/dgrijalva/jwt-go"
)

type AuthenticationClientOptions struct {
	Host      string
	Tenant    string
	Secret    string
	AccessKey string
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
