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

// 获取设备详情
func (c *SeanetClient) GetDevice(sn string) (*dto.Device, error) {
	uri := fmt.Sprintf(apiModifyDevice, sn)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return nil, err
	}

	return c.responseDevice(body)
}

// 更新设备
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

// 获取设备列表
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

// 控制设备
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

// 获取zigbee设备日志
func (c *SeanetClient) GetDeviceLog(sn string, pagination common.PaginationParams) (deviceLogsResp dto.DeviceLogDto, err error) {
	uri := fmt.Sprintf(apiLogDevice, sn)
	body, err := c.SendHttpRequest(uri, http.MethodGet, pagination)
	if err != nil {
		return deviceLogsResp, err
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

	err = json.Unmarshal(body, &deviceLogsResp)
	if err != nil {
		return
	}

	if !deviceLogsResp.Status {
		err = fmt.Errorf("code : %d", deviceLogsResp.Code)
	}

	return
}

// 获取zigbee网关下的设备列表
func (c *SeanetClient) ListGatewayDevice(gatewaySn string) (BelongDevicesResp dto.BelongDeviceDto, err error) {
	uri := fmt.Sprintf(apiBelongDevices, gatewaySn)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
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

	err = json.Unmarshal(body, &BelongDevicesResp)
	if err != nil {
		return
	}

	if !BelongDevicesResp.Status {
		err = fmt.Errorf("code : %d", BelongDevicesResp.Code)
	}

	return
}

const (
	Forget     = "ZbForget"
	Restart    = "Restart"
	PermitJoin = "ZbPermitJoin"
)

// 控制网关设备
func (c *SeanetClient) ControlGatewayDevice(devicekey, cmd, content string) (resp *common.BaseResponse, err error) {
	var param dto.CmdDeviceDto
	device, err := c.GetDevice(devicekey)
	if err != nil {
		return nil, err
	}

	param.Topic = fmt.Sprintf("/%s/%s/user/cmnd/%s", device.ProductKey, device.Sn, cmd)
	param.Content = content

	body, err := c.SendHttpRequest(apiCmdDevice, http.MethodPost, param)
	if err != nil {
		return nil, err
	}

	var result common.BaseResponse
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
