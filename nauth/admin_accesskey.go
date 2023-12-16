package nauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/valyala/fastjson"
)

func (cli *AuthenticationAdmin) GetAccessKey(accessKey string) (*dto.AccessKeyResponseDto, error) {
	if accessKey == "" {
		return nil, errors.New("accessKey is required")
	}

	uri := fmt.Sprintf("/admin/accesskey/%s", accessKey)
	body, err := cli.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return cli.responseAccessKeyResp(body)
}

func (cli *AuthenticationAdmin) responseAccessKeyResp(b []byte) (*dto.AccessKeyResponseDto, error) {
	var p fastjson.Parser
	v, err := p.Parse(string(b))
	if err != nil {
		return nil, err
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(msg))
	}

	byteAccessKey := v.MarshalTo(nil)
	resultAccessKey := dto.AccessKeyResponseDto{}
	err = json.Unmarshal(byteAccessKey, &resultAccessKey)
	if err != nil {
		return nil, err
	}

	return &resultAccessKey, nil
}
