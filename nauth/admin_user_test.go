// @Title: 管理员用户 SDK 测试
// @Description: 验证管理员用户列表查询参数的请求契约
// @Author: Agent:claude
// @Date: 2026-08-02
package nauth

import (
	"encoding/json"
	"testing"

	"github.com/netkitcloud/sdk-go/nauth/param"
)

func TestQueryUsers_UsernameJSONContract(t *testing.T) {
	payload, err := json.Marshal(param.QueryUsers{Username: "alice"})
	if err != nil {
		t.Fatalf("failed to marshal query: %v", err)
	}
	if string(payload) != `{"current":0,"per_page":0,"creator_uid":"","username":"alice"}` {
		t.Fatalf("unexpected query payload: %s", payload)
	}
}

func TestQueryUsers_SearchJSONContract(t *testing.T) {
	payload, err := json.Marshal(param.QueryUsers{Search: "alice"})
	if err != nil {
		t.Fatalf("failed to marshal query: %v", err)
	}
	if string(payload) != `{"current":0,"per_page":0,"creator_uid":"","search":"alice"}` {
		t.Fatalf("unexpected query payload: %s", payload)
	}
}
