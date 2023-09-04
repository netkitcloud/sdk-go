package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/nauth"
	"github.com/netkitcloud/sdk-go/nauth/dto"
	"github.com/netkitcloud/sdk-go/response"
)

// @description: 创建用户的ak
// @Param body dto.AddAccessKeyDto true "创建用户ak的备注"
// @Router /accesskey/comment [post]
func ginAddAccessKey(c *gin.Context, cli *nauth.AuthenticationClient) {
	// 获取并绑定请求参数
	params := new(dto.AddAccessKeyDto)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.AddAccessKey(dto.AddAccessKeyDto{
		Comment: params.Comment,
	})
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 获取并返回用户指定的ak信息
// @Param path queryKey string true "用户accesskey"
// @Router /accesskey/:queryKey [GET]
func ginGetAccessKey(c *gin.Context, cli *nauth.AuthenticationClient) {
	datas, err := cli.GetAccessKey(c.Param("queryKey"))
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas))
}

// @description: 删除用户指定ak
// @Param path queryKey string true "用户accesskey"
// @Router /accesskey/:queryKey [DELETE]
func ginDeleteAccessKey(c *gin.Context, cli *nauth.AuthenticationClient) {
	datas, err := cli.DeleteAccessKey(c.Param("queryKey"))
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 更新用户指定ak的comment备注
// @Param body dto.AddAccessKeyDto true "更新用户ak的备注"
// @Router /accesskey/:queryKey/comment [PUT]
func ginUpdateAccessKey(c *gin.Context, cli *nauth.AuthenticationClient) {
	// 获取并绑定请求参数
	params := new(dto.AddAccessKeyDto)
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.UpdateAccessKey(dto.UpdateAccessKeyDto{
		AccessKey: c.Param("queryKey"),
		Comment:   params.Comment,
	})
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 获取并返回用户所有的ak信息
// @Param path current  true "当前分页"
// @Param path per_page  true "每个个数"
// @Router /accesskey?current=1&per_page=10 [GET]
func ginListAccessKey(c *gin.Context, cli *nauth.AuthenticationClient) {
	// 获取并绑定传入的参数
	params := new(common.PaginationParams)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.ListAccessKey(common.PaginationParams{
		Current: params.Current,
		PerPage: params.PerPage,
	})
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, response.NewResponseData(response.Success, datas.Data))
}

// @description: 重设用户指定ak的secret
// @Param path queryKey  string "用户accesskey"
// @Router /accesskey/:queryKey/reset [PUT]
func ginResetAccessKey(c *gin.Context, cli *nauth.AuthenticationClient) {
	datas, err := cli.RsetAccessSecret(c.Param("queryKey"))
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}
