package nauth

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/netkitcloud/sdk-go/nauth/param"
	"github.com/valyala/fastjson"
)

func (cli *AuthenticationAdmin) AddUser(userInfo dto.AddUserDto) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest("/users", http.MethodPost, userInfo)
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

func (cli *AuthenticationAdmin) GetUserById(userId string) (*dto.GetUserResponseDto, error) {
	body, err := cli.SendHttpRequest(fmt.Sprint("/users/", userId), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var result dto.GetUserResponseDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) GetUserByUsername(username string) (*dto.GetUserResponseDto, error) {
	body, err := cli.SendHttpRequest(fmt.Sprint("/users/username/", username), http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	var result dto.GetUserResponseDto
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) UpdateUser(userId string, userInfo dto.UpdateUserDto) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest(fmt.Sprint("/users/", userId), http.MethodPut, userInfo)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) DeleteUser(userId string) (*common.BaseResponse, error) {
	body, err := cli.SendHttpRequest(fmt.Sprint("/users/", userId), http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (cli *AuthenticationAdmin) ListUser(pagination param.QueryUsers) (userListResp dto.UserListDto, err error) {
	body, err := cli.SendHttpRequest("/users", http.MethodGet, pagination)
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

	err = json.Unmarshal(body, &userListResp)
	if err != nil {
		return
	}

	if !userListResp.Status {
		err = fmt.Errorf("code : %d", userListResp.Code)
	}

	return
}
