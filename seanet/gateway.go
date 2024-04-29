package seanet

import (
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/seanet/dto"
	"github.com/netkitcloud/sdk-go/seanet/param"
)

type Action string

const (
	Map        Action = "ZbMap"        // ZbMap 0x0000，查询该网关下所有设备
	Config     Action = "ZbConfig"     // ZbConfig ，查看网关当前信道
	Scan       Action = "ZbScan"       // ZbScan ，查看网关wifi下所有信道信号强度
	Info       Action = "ZbInfo"       // ZbInfo ， 查看网关知道的设备mac及ip，记录子设备的可达性和链接质量
	Forget     Action = "ZbForget"     // ZbForget <device> ， Remove a device from the Tasmota flash memory.
	Leave      Action = "ZbLeave"      // ZbLeave <device> ， request a device to leave the network. If the device is offline or sleeping, this will have no effect. It is not 100% guaranteed that the device will never connect again
	Restart    Action = "Restart"      // Restart 1 = restart device with configuration saved to flash
	PermitJoin Action = "ZbPermitJoin" // ZbPermitJoin 1 = enable pairing mode for 60 seconds

)

// 控制网关设备
func (c *SeanetClient) ControlGateway(devicekey string, cmd Action, content string) (resp common.BaseResponse, err error) {
	var param dto.CmdDeviceDto
	device, err := c.GetDevice(devicekey)
	if err != nil {
		return
	}

	param.Topic = fmt.Sprintf("/%s/%s/user/cmnd/%s", device.Data.ProductKey, device.Data.Sn, cmd)
	param.Content = content

	body, err := c.SendHttpRequest(apiCmdDevice, http.MethodPost, param)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 控制网关下子设备发送信息（从而实现控制设备的效果）
// 网关向子设备发送的主题为固定/:gatewayProductKey/:gatewayDeviceKey/user/cmnd/ZbSend，contents为控制设备的json内容
func (c *SeanetClient) SwitchDevice(params *param.SwitchDevice) (resp common.BaseResponse, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiStatusDevice, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 获取zigbee设备日志
func (c *SeanetClient) GetDeviceLog(sn string, pagination common.PaginationParams) (resp dto.DeviceLogDto, err error) {
	uri := fmt.Sprintf(apiLogDevice, sn)
	body, err := c.SendHttpRequest(uri, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 获取zigbee网关下的设备列表
func (c *SeanetClient) ListGatewaySubDevice(gatewaySn string, pagination common.PaginationParams) (resp dto.SubDevices, err error) {
	uri := fmt.Sprintf(apiSubDevices, gatewaySn)
	body, err := c.SendHttpRequest(uri, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}
