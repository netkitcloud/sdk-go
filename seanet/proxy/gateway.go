package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/response"
	"github.com/netkitcloud/sdk-go/seanet"
	"github.com/netkitcloud/sdk-go/seanet/param"
)

// @description: 控制网关设备
// @Param body param.ControlGateway true "控制网关的指令"
// @Router /device/gateway [POST]
func ginControlGateway(c *gin.Context, cli *seanet.SeanetClient) {
	// 获取并绑定传入的参数
	params := new(param.ControlGateway)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	// 控制网关设备
	datas, err := cli.ControlGateway(params.Sn, seanet.Action(params.Cmd), params.Content)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 控制网关下子设备发送信息（从而实现控制设备的效果）
// @Param body param.SwitchDevice true "控制网关的指令"
// @Router /device/status [post]
func ginSwitchDevice(c *gin.Context, cli *seanet.SeanetClient) {
	// 获取并绑定请求参数
	params := new(param.SwitchDevice)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.SwitchDevice(params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 获取并返回所有设备信息
// @Param query queryKey string true "设备sn"
// @Param path current  true "当前分页"
// @Param path per_page  true "每个个数"
// @Router /device/:queryKey/log [GET]
func ginGetDeviceLog(c *gin.Context, cli *seanet.SeanetClient) {
	// 获取并绑定传入的参数
	params := new(common.PaginationParams)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	device_sn := c.Param("queryKey")
	datas, err := cli.GetDeviceLog(device_sn, common.PaginationParams{
		Current: params.Current,
		PerPage: params.PerPage,
	})
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// 获取所有的网关规则列表（List gateway rule）
func ginListGatewayRule(c *gin.Context, cli *seanet.SeanetClient) {
	// 获取并绑定传入的参数
	params := new(common.PaginationParams)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.ListGatewayRule(*params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// 获取单个网关规则信息下所有关联的设备信息
// @Router /gateway/rule/:queryKey/device [GET]
func ginListGatewayRuleDevices(c *gin.Context, cli *seanet.SeanetClient) {
	gateway_rule_id := c.Param("queryKey")
	// 获取并绑定传入的参数
	params := new(common.PaginationParams)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.ListGatewayRuleDevices(gateway_rule_id, *params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 获取zigbee网关下的设备列表
// @Param query queryKey string true "设备sn"
// @Param path current  true "当前分页"
// @Param path per_page  true "每个个数"
// @Router /device/:queryKey/gateway [GET]
func ginListGatewaySubDevice(c *gin.Context, cli *seanet.SeanetClient) {
	gateway_sn := c.Param("queryKey")
	// 获取并绑定传入的参数
	params := new(common.PaginationParams)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.ListGatewaySubDevice(gateway_sn, *params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 在指定的网关规则下添加关联设备（将网关与子设备绑定）
// @Param query queryKey string true "设备sn"
// @Router /gateway/rule/:queryKey/device [POST]
func ginCreateGatewayRuleDevice(c *gin.Context, cli *seanet.SeanetClient) {
	gateway_sn := c.Param("queryKey")
	params := new(param.CreateGatewayRuleDevices)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.CreateGatewayRuleDevice(gateway_sn, *params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 在指定的网关规则下删除关联设备（将网关与子设备解绑）
// @Param query queryKey string true "设备sn"
// @Param path current  true "当前分页"
// @Param path per_page  true "每个个数"
// @Router /gateway/rule/:queryKey/device [DELETE]
func ginDeleteGatewayRuleDevice(c *gin.Context, cli *seanet.SeanetClient) {
	gateway_sn := c.Param("queryKey")
	params := new(param.DeleteGatewayRuleDevices)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.DeleteGatewayRuleDevice(gateway_sn, *params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}
