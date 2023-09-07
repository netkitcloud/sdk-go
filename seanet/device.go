package seanet

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/seanet/dto"
	"github.com/valyala/fastjson"
)

func (c *SeanetClient) GetDevice(sn string) (*dto.Device, error) {
	uri := fmt.Sprintf(apiModifyDevice, sn)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return c.responseDevice(body)
}

func (c *SeanetClient) UpdateDevice(dto *dto.UpdateDeviceDto) (*common.BaseResponse, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(apiModifyDevice, dto.Sn)
	body, err := c.SendHttpRequest(uri, http.MethodPut, dto)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *SeanetClient) ListDevice(pagination common.PaginationParams) (deviceListResp dto.ListDeviceDto, err error) {
	body, err := c.SendHttpRequest(apiDevice, http.MethodGet, pagination)
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

	err = json.Unmarshal(body, &deviceListResp)
	if err != nil {
		return
	}

	if !deviceListResp.Status {
		err = fmt.Errorf("code : %d", deviceListResp.Code)
	}

	return
}

func (c *SeanetClient) SwitchDevice(param *dto.SwitchDeviceDto) (resp *common.BaseResponse, err error) {
	if err := c.validate.Struct(param); err != nil {
		return nil, err
	}

	body, err := c.SendHttpRequest(apiStatusDevice, http.MethodPost, param)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *SeanetClient) responseDevice(b []byte) (*dto.Device, error) {
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

	byteDevice := v.GetObject("data").MarshalTo(nil)
	resultDevice := dto.Device{}
	err = json.Unmarshal(byteDevice, &resultDevice)
	if err != nil {
		return nil, err
	}

	return &resultDevice, nil
}
