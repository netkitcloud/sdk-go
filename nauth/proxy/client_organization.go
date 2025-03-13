package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/netkitcloud/sdk-go/nauth/param"
	"github.com/netkitcloud/sdk-go/response"
)

func ginCreateOrganization(c *gin.Context, cli *nauth.AuthenticationClient) {
	params := new(param.CreateOrganization)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.CreateOrganization(params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginDeleteOrganization(c *gin.Context, cli *nauth.AuthenticationClient) {
	organization_id := c.Param("organization_id")
	if organization_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "organization_id is required"))
		return
	}

	datas, err := cli.DeleteOrganization(organization_id)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginUpdateOrganization(c *gin.Context, cli *nauth.AuthenticationClient) {
	organization_id := c.Param("organization_id")
	if organization_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "organization_id is required"))
		return
	}

	params := new(param.UpdateOrganization)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.UpdateOrganization(organization_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginGetOrganization(c *gin.Context, cli *nauth.AuthenticationClient) {
	organization_id := c.Param("organization_id")
	if organization_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "organization_id is required"))
		return
	}

	datas, err := cli.GetOrganization(organization_id)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginListOrganization(c *gin.Context, cli *nauth.AuthenticationClient) {
	pagination := new(param.QueryOrganization)
	if err := c.ShouldBindJSON(pagination); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.ListOrganization(*pagination)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginOrgAddMember(c *gin.Context, cli *nauth.AuthenticationClient) {
	organization_id := c.Param("organization_id")
	if organization_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "organization_id is required"))
		return
	}

	params := new(dto.AddUserDto)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.OrgAddMember(organization_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginOrgBindMember(c *gin.Context, cli *nauth.AuthenticationClient) {
	organization_id := c.Param("organization_id")
	if organization_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "organization_id is required"))
		return
	}

	params := new(param.OrganizationUser)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.OrgBindMember(organization_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginOrgUnbindMember(c *gin.Context, cli *nauth.AuthenticationClient) {
	organization_id := c.Param("organization_id")
	if organization_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "organization_id is required"))
		return
	}

	params := new(param.OrganizationUser)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.OrgUnbindMember(organization_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginGetOrganizationUsers(c *gin.Context, cli *nauth.AuthenticationClient) {
	organization_id := c.Param("organization_id")
	if organization_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "organization_id is required"))
		return
	}

	pagination := new(common.PaginationParams)
	if err := c.ShouldBindJSON(pagination); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.GetOrganizationUsers(organization_id, *pagination)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginSetOrganizationUsersIsmanage(c *gin.Context, cli *nauth.AuthenticationClient) {
	organization_id := c.Param("organization_id")
	if organization_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "organization_id is required"))
		return
	}

	params := new(param.OrganizationUserManagerForm)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.SetOrganizationUsersIsmanage(organization_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}
