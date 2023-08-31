package nauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/valyala/fastjson"
)

func (cli *AuthenticationAdmin) GetAccessKey(accessKey string) (*dto.AccessKey, error) {
	if accessKey == "" {
		return nil, errors.New("accessKey is required")
	}

	uri := fmt.Sprintf("/admin/accesskey/%s", accessKey)
	body, err := cli.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return cli.responseAccessKey(body)
}

func (cli *AuthenticationAdmin) responseAccessKey(b []byte) (*dto.AccessKey, error) {
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

	byteAccessKey := v.GetObject("data").MarshalTo(nil)
	resultAccessKey := dto.AccessKey{}
	err = json.Unmarshal(byteAccessKey, &resultAccessKey)
	if err != nil {
		return nil, err
	}

	return &resultAccessKey, nil
}
