package seanet

import (
	"fmt"
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/seanet/dto"
	"github.com/netkitcloud/sdk-go/seanet/param"
)

// 给指定网关创建网关规则信息
func (c *SeanetClient) CreateGatewayRule(params *param.CreateGatewayRule) (resp common.BaseResponse, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiGatewayRule, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 删除指网关规则
func (c *SeanetClient) DeleteGatewayRule(params *param.DeleteGatewayRule) (resp common.BaseResponse, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	body, err := c.SendHttpRequest(apiGatewayRule, http.MethodDelete, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 获取所有的网关规则列表（List gateway rule）
func (c *SeanetClient) ListGatewayRule(pagination common.PaginationParams) (resp dto.GatewayRulesDto, err error) {
	body, err := c.SendHttpRequest(apiGatewayRule, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 获取单个网关规则信息下所有关联的设备信息
func (c *SeanetClient) ListGatewayRuleDevices(gatewayRuleID string, pagination common.PaginationParams) (resp dto.GatewayRuleDevicesDto, err error) {
	uri := fmt.Sprintf(apiGatewayRuleDevices, gatewayRuleID)
	body, err := c.SendHttpRequest(uri, http.MethodGet, pagination)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 在指定的网关规则下添加关联设备（将网关与子设备绑定）
func (c *SeanetClient) CreateGatewayRuleDevice(gatewaySn string, params param.CreateGatewayRuleDevices) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiGatewayRuleDevices, gatewaySn)
	body, err := c.SendHttpRequest(uri, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 在指定的网关规则下删除关联设备（将网关与子设备解绑）
func (c *SeanetClient) DeleteGatewayRuleDevice(gatewaySn string, params param.DeleteGatewayRuleDevices) (resp common.BaseResponse, err error) {
	uri := fmt.Sprintf(apiGatewayRuleDevices, gatewaySn)
	body, err := c.SendHttpRequest(uri, http.MethodDelete, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}
