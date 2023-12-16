package nauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/valyala/fastjson"
)

func (cli *AuthenticationClient) AddAccessKey(dto dto.AddAccessKeyDto) (*dto.AccessKey, error) {
	body, err := cli.SendHttpRequest("/accesskey", http.MethodPost, dto)
	if err != nil {
		return nil, err
	}

	return cli.responseAccessKey(body)
}

func (cli *AuthenticationClient) DeleteAccessKey(accessKey string) (*common.BaseResponse, error) {
	if accessKey == "" {
		return nil, errors.New("accessKey is required")
	}

	uri := fmt.Sprintf("/accesskey/%s", accessKey)
	body, err := cli.SendHttpRequest(uri, http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationClient) RsetAccessSecret(accessKey string) (*dto.ResetAccessSecretDto, error) {
	if accessKey == "" {
		return nil, errors.New("accessKey is required")
	}

	uri := fmt.Sprintf("/accesskey/%s/reset", accessKey)
	body, err := cli.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var result dto.ResetAccessSecretDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationClient) UpdateAccessKey(dto dto.UpdateAccessKeyDto) (*common.BaseResponse, error) {
	if dto.AccessKey == "" {
		return nil, errors.New("accessKey is required")
	}

	uri := fmt.Sprintf("/accesskey/%s", dto.AccessKey)
	body, err := cli.SendHttpRequest(uri, http.MethodPut, dto)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationClient) GetAccessKey(accessKey string) (*dto.AccessKey, error) {
	if accessKey == "" {
		return nil, errors.New("accessKey is required")
	}

	uri := fmt.Sprintf("/accesskey/%s", accessKey)
	body, err := cli.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return cli.responseAccessKey(body)
}

func (cli *AuthenticationClient) ListAccessKey(pagination common.PaginationParams) (accesskeyListResp *dto.ListAccessKeyDto, err error) {
	body, err := cli.SendHttpRequest("/accesskey", http.MethodGet, pagination)
	if err != nil {
		return
	}

	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return
		}
		err = errors.New(string(msg))
		return
	}

	err = json.Unmarshal(body, &accesskeyListResp)
	if err != nil {
		return
	}

	if !accesskeyListResp.Status {
		err = fmt.Errorf("code : %d", accesskeyListResp.Code)
	}

	return
}

func (cli *AuthenticationClient) responseAccessKey(b []byte) (*dto.AccessKey, error) {
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
