package nauth

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"sort"
	"strings"

	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/netkitcloud/sdk-go/nauth/dto"
)

func (c *AuthenticationClient) getSignature(dto map[string]string) string {
	var keys []string
	var keyValues []string
	for k := range dto {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	for _, k := range keys {
		keyValues = append(keyValues, fmt.Sprintf("%s=%s", k, dto[k]))
	}
	signatureString := strings.Join(keyValues, "&")
	h := hmac.New(sha256.New, []byte(c.options.AppSecret))
	h.Write([]byte(signatureString))
	return base64.StdEncoding.EncodeToString(h.Sum(nil))
}

func (c *AuthenticationClient) GetOAuthLoginUrl(state string) string {
	return fmt.Sprintf("%s/oauth/%s/authorize?appid=%s&state=%s&response_type=code",
		c.options.Host, c.options.Tenant, c.options.AppId, state)
}

func (c *AuthenticationClient) GetAccessTokenByCode(code string) (tokenResponse dto.OIDCTokenResponse, err error) {
	dto := map[string]string{
		"grant_type": "authorization_code",
		"code":       code,
		"appid":      c.options.AppId,
	}

	dto["signature"] = c.getSignature(dto)

	resp, err := c.SendHttpRequest("/login/username", http.MethodGet, dto)
	if err != nil {
		return tokenResponse, err
	}
	err = json.Unmarshal(resp, &tokenResponse)
	return tokenResponse, err
}

func (c *AuthenticationClient) Verify(token string) bool {
	// 创建一个context
	ctx := context.Background()
	// 创建一个oidc provider，传入issuer的URL
	provider, err := oidc.NewProvider(ctx, fmt.Sprintf(c.options.Host+"/oauth/%s", c.options.Tenant))
	if err != nil {
		// 处理错误
		fmt.Println(err)
		return false
	}

	verifier := provider.Verifier(&oidc.Config{
		ClientID:          c.options.AppId,
		SkipIssuerCheck:   false,
		SkipClientIDCheck: false,
		SkipExpiryCheck:   false,
	})
	_, err = verifier.Verify(ctx, token)
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
