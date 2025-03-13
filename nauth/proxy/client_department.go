package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth"
	"github.com/netkitcloud/sdk-go/nauth/param"
	"github.com/netkitcloud/sdk-go/response"
)

func ginCreateDepartment(c *gin.Context, cli *nauth.AuthenticationClient) {
	params := new(param.CreateDepartment)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.CreateDepartment(params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginDeleteDepartment(c *gin.Context, cli *nauth.AuthenticationClient) {
	department_id := c.Param("department_id")
	if department_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "department_id is required"))
		return
	}

	datas, err := cli.DeleteDepartment(department_id)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginUpdateDepartment(c *gin.Context, cli *nauth.AuthenticationClient) {
	department_id := c.Param("department_id")
	if department_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "department_id is required"))
		return
	}

	params := new(param.UpdateDepartment)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.UpdateDepartment(department_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginGetDepartment(c *gin.Context, cli *nauth.AuthenticationClient) {
	department_id := c.Param("department_id")
	if department_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "department_id is required"))
		return
	}

	datas, err := cli.GetDepartment(department_id)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginListDepartment(c *gin.Context, cli *nauth.AuthenticationClient) {
	pagination := new(param.QueryDepartment)
	if err := c.ShouldBindJSON(pagination); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.ListDepartment(*pagination)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginDepBindMember(c *gin.Context, cli *nauth.AuthenticationClient) {
	department_id := c.Param("department_id")
	if department_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "department_id is required"))
		return
	}

	params := new(param.DepartmentUser)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.DepBindMember(department_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginDepUnbindMember(c *gin.Context, cli *nauth.AuthenticationClient) {
	department_id := c.Param("department_id")
	if department_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "department_id is required"))
		return
	}

	params := new(param.DepartmentUser)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.DepUnbindMember(department_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginGetDepartmentUsers(c *gin.Context, cli *nauth.AuthenticationClient) {
	department_id := c.Param("department_id")
	if department_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "department_id is required"))
		return
	}

	pagination := new(common.PaginationParams)
	if err := c.ShouldBindJSON(pagination); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.GetDepartmentUsers(department_id, *pagination)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

func ginSetDepartmentUsersIsmanage(c *gin.Context, cli *nauth.AuthenticationClient) {
	department_id := c.Param("department_id")
	if department_id == "" {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, "department_id is required"))
		return
	}

	params := new(param.DepartmentUserManagerForm)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.SetDepartmentUsersIsmanage(department_id, params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}
