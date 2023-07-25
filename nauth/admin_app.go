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

func (cli *AuthenticationAdmin) AddApp(app dto.AddAppDto) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest("/apps", http.MethodPost, app)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if !result.Status {
		return &result, fmt.Errorf("code: %d", result.Code)
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) UpdateApp(appId string, app dto.UpdateAppDto) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest(fmt.Sprint("/apps/", appId), http.MethodPut, app)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) GetApp(appId string) (*dto.GetUserResponseDto, error) {
	body, err := cli.SendHttpRequest(fmt.Sprint("/apps/", appId), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var result dto.GetUserResponseDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) DeleteApp(appId string) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest(fmt.Sprint("/apps/", appId), http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) ListApp(pagination common.PaginationParams) (appListResp dto.ListAppDto, err error) {
	body, err := cli.SendHttpRequest("/apps", http.MethodGet, pagination)
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

	err = json.Unmarshal(body, &appListResp)
	if err != nil {
		return
	}

	if !appListResp.Status {
		err = fmt.Errorf("code : %d", appListResp.Code)
	}

	return
}
