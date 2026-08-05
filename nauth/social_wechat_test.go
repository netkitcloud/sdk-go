package nauth

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/netkitcloud/sdk-go/nauth/dto"
)

func TestSocialWxOfficeMethodsMatchBackendContract(t *testing.T) {
	tests := []struct {
		name       string
		call       func(*AuthenticationClient) error
		wantMethod string
		wantPath   string
		wantQuery  map[string]string
	}{
		{
			name: "auth url",
			call: func(client *AuthenticationClient) error {
				_, err := client.SocialWxOfficeAuthURL("main", "office")
				return err
			},
			wantMethod: http.MethodPost,
			wantPath:   "/main/social/login/office",
		},
		{
			name: "login callback",
			call: func(client *AuthenticationClient) error {
				_, err := client.SocialWxOfficeLogin("main", "office", &dto.WxOfficeCallbackDto{Code: "code-1", State: "state-1"})
				return err
			},
			wantMethod: http.MethodGet,
			wantPath:   "/main/social/login/office/callback",
			wantQuery:  map[string]string{"code": "code-1", "state": "state-1"},
		},
		{
			name: "bind auth url",
			call: func(client *AuthenticationClient) error {
				_, err := client.SocialWxOfficeBindAuthURL("main", "office")
				return err
			},
			wantMethod: http.MethodPost,
			wantPath:   "/main/social/bind/office",
		},
		{
			name: "bind callback",
			call: func(client *AuthenticationClient) error {
				_, err := client.SocialWxOfficeBindCallback("main", "office", &dto.WxOfficeCallbackDto{Code: "code-2", State: "state-2"})
				return err
			},
			wantMethod: http.MethodGet,
			wantPath:   "/main/social/bind/office/callback",
			wantQuery:  map[string]string{"code": "code-2", "state": "state-2"},
		},
		{
			name: "rebind auth url",
			call: func(client *AuthenticationClient) error {
				_, err := client.SocialWxOfficeReBindAuthURL("main", "office")
				return err
			},
			wantMethod: http.MethodPost,
			wantPath:   "/main/social/rebind/office",
		},
		{
			name: "rebind callback",
			call: func(client *AuthenticationClient) error {
				_, err := client.SocialWxOfficeReBindCallback("main", "office", &dto.WxOfficeCallbackDto{Code: "code-3", State: "state-3"})
				return err
			},
			wantMethod: http.MethodGet,
			wantPath:   "/main/social/rebind/office/callback",
			wantQuery:  map[string]string{"code": "code-3", "state": "state-3"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != tt.wantMethod {
					t.Fatalf("expected method %s, got %s", tt.wantMethod, r.Method)
				}
				if r.URL.Path != tt.wantPath {
					t.Fatalf("expected path %s, got %s", tt.wantPath, r.URL.Path)
				}
				for key, want := range tt.wantQuery {
					if got := r.URL.Query().Get(key); got != want {
						t.Fatalf("expected query %s=%s, got %s", key, want, got)
					}
				}
				writeSocialWechatTestResponse(t, w, tt.wantPath)
			}))
			defer server.Close()

			client, err := NewClient(&AuthenticationClientOptions{
				Host:      server.URL,
				Tenant:    "main",
				AppId:     "app-1",
				AppSecret: "secret-1",
			})
			if err != nil {
				t.Fatalf("NewClient returned error: %v", err)
			}
			if err := tt.call(client); err != nil {
				t.Fatalf("method returned error: %v", err)
			}
		})
	}
}

func writeSocialWechatTestResponse(t *testing.T, w http.ResponseWriter, path string) {
	t.Helper()

	w.Header().Set("Content-Type", "application/json")
	if path == "/main/social/login/office" || path == "/main/social/bind/office" || path == "/main/social/rebind/office" {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": true,
			"code":   0,
			"data": map[string]interface{}{
				"auth_code_url": "https://wechat.example/auth",
				"state":         "state-1",
				"expires_in":    300,
			},
		})
		return
	}
	if path == "/main/social/login/office/callback" {
		_ = json.NewEncoder(w).Encode(map[string]interface{}{
			"status": true,
			"code":   0,
			"data": map[string]interface{}{
				"uid":          "uid-1",
				"access_token": "access-token",
			},
		})
		return
	}
	_ = json.NewEncoder(w).Encode(map[string]interface{}{"status": true, "code": 0})
}
