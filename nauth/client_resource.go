package nauth

import (
	"net/http"

	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/netkitcloud/sdk-go/nauth/param"
)

// 资源权限授权
func (c *AuthenticationClient) Authorize(params *param.AuthorizeResourcesDto) (resp common.BaseResponse, err error) {
	body, err := c.SendHttpRequest(apiAuthorizeResource, http.MethodPost, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 资源权限取消授权
func (c *AuthenticationClient) Unauthorize(params *param.AuthorizeResourcesDto) (resp common.BaseResponse, err error) {
	body, err := c.SendHttpRequest(apiUnauthorizeResource, http.MethodDelete, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}

	return
}

// 资源权限校验
func (c *AuthenticationClient) VerifyAction(params *param.QueryVerifyAction) (resp dto.VerifyActionDto, err error) {
	if err = c.validate.Struct(params); err != nil {
		return
	}

	if params.TenantID == "" {
		params.TenantID = c.options.Tenant
	}

	body, err := c.SendHttpRequest(apiVerifyAction, http.MethodGet, params)
	if err != nil {
		return
	}

	if err = common.ParserDto(body, &resp); err != nil {
		return
	}
	return
}
