package proxy

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/netkitcloud/sdk-go/response"
	"github.com/netkitcloud/sdk-go/seanet"
)

type Action string

type NormalProxyFunc func(c *gin.Context, cli *seanet.SeanetClient)

type ActionFuncMap[KEY Action, VALUE NormalProxyFunc] map[KEY]VALUE

const (
	GetDevice         Action = "GetDevice"
	UpdateDevice      Action = "UpdateDevice"
	ListDevice        Action = "ListDevice"
	SwitchDevice      Action = "SwitchDevice"
	GetDeviceLog      Action = "GetDeviceLog"
	ListGatewayDevice Action = "ListGatewayDevice"
)

var ActionFunc ActionFuncMap[Action, NormalProxyFunc] = map[Action]NormalProxyFunc{
	GetDevice:         ginGetDevice,
	UpdateDevice:      ginUpdateDevice,
	ListDevice:        ginListDevice,
	SwitchDevice:      ginSwitchDevice,
	GetDeviceLog:      ginGetDeviceLog,
	ListGatewayDevice: ginListGatewayDevice,
}

func GinRouterClientProxy(action Action, options *seanet.SeanetClientOptions) func(*gin.Context) {
	return func(c *gin.Context) {
		cli, err := ginInitClinet(c, options)
		if err != nil {
			c.JSON(http.StatusOK, response.NewResponseMessage(response.InitClientError, err.Error()))
			return
		}

		ActionFunc[action](c, cli)
	}
}

func getBearToken(c *gin.Context) string {
	bearToken := c.Request.Header.Get("Authorization")
	//normally Authorization the_token_xxx
	strArr := strings.Split(bearToken, " ")
	if len(strArr) == 2 {
		return strArr[1]
	} else if len(strArr) == 1 && strArr[0] != "" {
		return strArr[0]
	}
	return ""
}

func ginInitClinet(c *gin.Context, options *seanet.SeanetClientOptions) (*seanet.SeanetClient, error) {
	// 获取token信息并保存到上下文中
	tokenString := getBearToken(c)
	c.Set(seanet.CurrentUserToken, tokenString)

	// Initial authentication client
	cli, err := seanet.NewClient(options)
	if err != nil {
		return nil, err
	}
	if err = cli.SetToken(c.GetString(seanet.CurrentUserToken)); err != nil {
		return nil, err
	}

	return cli, nil
}
