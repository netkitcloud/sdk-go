package proxy

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/netkitcloud/sdk-go/common"
	"github.com/netkitcloud/sdk-go/response"
	"github.com/netkitcloud/sdk-go/seanet"
	"github.com/netkitcloud/sdk-go/seanet/param"
)

// @description: 根据设备sn获取设备信息
// @Param body sn string true "设备sn"
// @Router /device/:queryKey [get]
func ginGetDevice(c *gin.Context, cli *seanet.SeanetClient) {
	// 获取并绑定请求参数
	datas, err := cli.GetDevice(c.Param("queryKey"))
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 更新指定的设备信息
// @Param body param.UpdateDevice true "更新设备信息"
// @Router /device/:queryKey [PUT]
func ginUpdateDevice(c *gin.Context, cli *seanet.SeanetClient) {
	// 获取并绑定请求参数
	params := new(param.UpdateDevice)
	params.Sn = c.Param("queryKey")
	if err := c.ShouldBindJSON(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.UpdateDevice(params)
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}

// @description: 获取并返回所有设备信息
// @Param path current  true "当前分页"
// @Param path per_page  true "每个个数"
// @Router /device [GET]
func ginListDevice(c *gin.Context, cli *seanet.SeanetClient) {
	// 获取并绑定传入的参数
	params := new(common.PaginationParams)
	if err := c.ShouldBindQuery(params); err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.FromError, err.Error()))
		return
	}

	datas, err := cli.ListDevice(common.PaginationParams{
		Current: params.Current,
		PerPage: params.PerPage,
	})
	if err != nil {
		c.JSON(http.StatusOK, response.NewResponseMessage(response.GetError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, datas)
}
