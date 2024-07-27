package seanet

import (
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/seanet/dto"
	"github.com/netkitcloud/sdk-go/seanet/param"
)

// 获取设备列表
func (c *SeanetClient) GetProperty(params *param.GetProperty) (resp dto.GetPropertyDto, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	url := fmt.Sprintf(apiDeviceProperty, params.Devicekey, params.Property)

	body, err := c.SendHttpRequest(url, http.MethodGet, params.Query)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
