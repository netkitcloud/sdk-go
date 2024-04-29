package seanet

import (
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/seanet/dto"
	"github.com/netkitcloud/sdk-go/seanet/param"
)

// 获取设备详情
func (c *SeanetClient) GetDevice(sn string) (resp dto.DeviceDto, err error) {
	uri := fmt.Sprintf(apiModifyDevice, sn)
	body, err := c.SendHttpRequest(uri, http.MethodGet, nil)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 更新设备
func (c *SeanetClient) UpdateDevice(params *param.UpdateDevice) (resp common.BaseResponse, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	uri := fmt.Sprintf(apiModifyDevice, params.Sn)
	body, err := c.SendHttpRequest(uri, http.MethodPut, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}

// 获取设备列表
func (c *SeanetClient) ListDevice(pagination common.PaginationParams) (resp dto.ListDeviceDto, err error) {
	if err = c.validate.Struct(pagination); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiDevice, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
