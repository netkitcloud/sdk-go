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

func (cli *AuthenticationAdmin) AddAppUser(appUser dto.AppUserCreateDto) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest("/app_user", http.MethodPost, appUser)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	if !result.Status {
		return &result, fmt.Errorf("code: %d, message: %s", result.Code, result.Message)
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) UpdateAppUser(appUser dto.AppUserUpdateDto) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest("/app_user", http.MethodPut, appUser)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) DeleteAppUser(appUser dto.AppUserDeleteDto) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest("/app_user", http.MethodDelete, appUser)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) ListAppUser(pagination common.PaginationParams) (resp dto.ListAppUserDto, err error) {
	body, err := cli.SendHttpRequest("/app_user", http.MethodGet, pagination)
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
		err = errors.New(string(msg))
		return
	}

	err = json.Unmarshal(body, &resp)
	if err != nil {
		return
	}

	if !resp.Status {
		err = fmt.Errorf("code : %d", resp.Code)
	}

	return
}
