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

func (cli *AuthenticationAdmin) ListUser(pagination common.PaginationParams) (userListResp dto.UserListDto, err error) {
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
		if err != nil {
			return
		}
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

// 发送短信
func (cli *AuthenticationAdmin) PostSMS(phone string, content string) error {
	body, err := cli.SendHttpRequest("/user/sms/content", http.MethodPost, map[string]interface{}{
		"phone":   phone,
		"content": content,
	})
	if err != nil {
		return err
	}
	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return err
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return err
		}
		return errors.New(string(msg))
	}

	return nil

}

func (c *AuthenticationAdmin) PostEmail(email string, content string, scene int) error {

	body, err := c.SendHttpRequest("/user/email", http.MethodPost, map[string]interface{}{
		"email":   email,
		"content": content,
		"scene":   scene,
	})
	if err != nil {
		return err
	}

	var p fastjson.Parser
	v, err := p.Parse(string(body))
	if err != nil {
		return err
	}

	if !v.GetBool("status") {
		msg := v.GetStringBytes("message")
		if err != nil {
			return err
		}
		return errors.New(string(msg))
	}

	return nil
}
