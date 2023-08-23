package nip

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nip/dto"
	"github.com/valyala/fastjson"
)

// 获取该uid的所有NATS配置信息
func (c *NIPClient) GetNatsUser() (*dto.NatsUser, error) {
	body, err := c.SendHttpRequest(apiNatsUser, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return c.responseNatsUser(body)
}

// 查询该uid的NATS功能是否开通
func (c *NIPClient) GetNatsStatus() (*dto.NatsStatus, error) {
	body, err := c.SendHttpRequest(apiGetNatsStatus, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return c.responseNatsStatus(body)
}

// 若该uid未开通NATS功能，则开通该uid的NATS功能
func (c *NIPClient) CreateNatsUser() (*common.BaseResponse, error) {
	body, err := c.SendHttpRequest(apiNatsUser, http.MethodPost, nil)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// 获取该uid的NATS权限信息
func (c *NIPClient) GetNatsPermission() (*dto.NatsPermission, error) {
	body, err := c.SendHttpRequest(apiNatsPermission, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return c.responseNatsPermission(body)
}

// 更新该uid的NATS权限信息（添加监听的主题）
func (c *NIPClient) AddNatsPermission(dto *dto.AddNatsPermissionDto) (*common.BaseResponse, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	body, err := c.SendHttpRequest(apiNatsPermission, http.MethodPost, dto)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// 删除该uid的NATS权限信息（删除监听的主题）
func (c *NIPClient) DeleteNatsPermission(dto *dto.DeleteNatsPermissionDto) (*common.BaseResponse, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	body, err := c.SendHttpRequest(apiNatsPermission, http.MethodDelete, dto)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *NIPClient) responseNatsUser(b []byte) (*dto.NatsUser, error) {
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

	byteNatsUser := v.GetObject("data").MarshalTo(nil)
	resultNatsUser := dto.NatsUser{}
	err = json.Unmarshal(byteNatsUser, &resultNatsUser)
	if err != nil {
		return nil, err
	}

	return &resultNatsUser, nil
}

func (c *NIPClient) responseNatsStatus(b []byte) (*dto.NatsStatus, error) {
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

	byteNatsStatus := v.GetObject("data").MarshalTo(nil)
	resultNatsStatus := dto.NatsStatus{}
	err = json.Unmarshal(byteNatsStatus, &resultNatsStatus)
	if err != nil {
		return nil, err
	}

	return &resultNatsStatus, nil
}

func (c *NIPClient) responseNatsPermission(b []byte) (*dto.NatsPermission, error) {
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

	byteNatsPermission := v.GetObject("data").MarshalTo(nil)
	resultNatsPermission := dto.NatsPermission{}
	err = json.Unmarshal(byteNatsPermission, &resultNatsPermission)
	if err != nil {
		return nil, err
	}

	return &resultNatsPermission, nil
}
