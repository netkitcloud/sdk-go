package nip

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nip/dto"
	"github.com/valyala/fastjson"
)

func (c *NIPClient) CreateDevice(dtop *dto.CreateDeviceDto) (*dto.Device, error) {
	// create and return that device
	if err := c.validate.Struct(dtop); err != nil {
		return nil, err
	}

	_, err := c.SendHttpRequest(apiDevice, http.MethodPost, dtop)
	if err != nil {
		return nil, err
	}

	return c.GetDevice(&dto.GetDeviceDto{
		Productkey: dtop.Productkey,
		Devicekey:  dtop.Devicekey,
	})
}

func (c *NIPClient) GetDevice(dto *dto.GetDeviceDto) (*dto.Device, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(apiModifyDevice, dto.Devicekey)
	body, err := c.SendHttpRequest(uri, http.MethodGet, dto)
	if err != nil {
		return nil, err
	}

	return c.responseDevice(body)
}

func (c *NIPClient) DeleteDevice(deviceKey string) (*common.BaseResponse, error) {
	if deviceKey == "" {
		return nil, errors.New("deviceKey is required")
	}

	uri := fmt.Sprintf(apiModifyDevice, deviceKey)
	body, err := c.SendHttpRequest(uri, http.MethodDelete, nil)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func (c *NIPClient) UpdateDevice(dto *dto.UpdateDeviceDto) (*dto.Device, error) {
	if err := c.validate.Struct(dto); err != nil {
		return nil, err
	}

	uri := fmt.Sprintf(apiModifyDevice, dto.Devicekey)
	body, err := c.SendHttpRequest(uri, http.MethodPut, dto)
	if err != nil {
		return nil, err
	}

	return c.responseDevice(body)
}

func (c *NIPClient) ListDevice(pagination common.PaginationParams) (deviceListResp dto.ListDeviceDto, err error) {
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

func (c *NIPClient) responseDevice(b []byte) (*dto.Device, error) {
	var p fastjson.Parser
	v, err := p.Parse(string(b))
	if err != nil {
		return nil, err
	}

	if !v.GetBool("status") || v.GetObject("data") == nil {
		msg := v.GetStringBytes("message")
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
