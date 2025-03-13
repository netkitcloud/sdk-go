package proxy

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/netkitcloud/sdk-go/nauth"
	"github.com/netkitcloud/sdk-go/response"
)

type Action string

type NormalProxyFunc func(c *gin.Context, cli *nauth.AuthenticationClient)

type ActionFuncMap[KEY Action, VALUE NormalProxyFunc] map[KEY]VALUE

const (
	GetAccessKey           Action = "GetAccessKey"
	DeleteAccessKey        Action = "DeleteAccessKey"
	ListAccessKeys         Action = "ListAccessKeys"
	ResetAccessKey         Action = "ResetAccessKey"
	AddAccessKeyComment    Action = "AddAccessKeyComment"
	UpdateAccessKeyComment Action = "UpdateAccessKeyComment"

	CreateOrganization Action = "CreateOrganization"
	DeleteOrganization Action = "DeleteOrganization"
	UpdateOrganization Action = "UpdateOrganization"
	GetOrganization    Action = "GetOrganization"
	ListOrganization   Action = "ListOrganization"

	OrgAddMember                 Action = "OrgAddMember"
	OrgBindMember                Action = "OrgBindMember"
	OrgUnbindMember              Action = "OrgUnbindMember"
	GetOrganizationUsers         Action = "GetOrganizationUsers"
	SetOrganizationUsersIsmanage Action = "SetOrganizationUsersIsmanage"

	CreateDepartment Action = "CreateDepartment"
	DeleteDepartment Action = "DeleteDepartment"
	UpdateDepartment Action = "UpdateDepartment"
	GetDepartment    Action = "GetDepartment"
	ListDepartment   Action = "ListDepartment"

	DepBindMember              Action = "DepBindMember"
	DepUnbindMember            Action = "DepUnbindMember"
	GetDepartmentUsers         Action = "GetDepartmentUsers"
	SetDepartmentUsersIsmanage Action = "SetDepartmentUsersIsmanage"
)

var ActionFunc ActionFuncMap[Action, NormalProxyFunc] = map[Action]NormalProxyFunc{
	GetAccessKey:           ginGetAccessKey,
	DeleteAccessKey:        ginDeleteAccessKey,
	ListAccessKeys:         ginListAccessKey,
	ResetAccessKey:         ginResetAccessKey,
	AddAccessKeyComment:    ginAddAccessKeyComment,
	UpdateAccessKeyComment: ginUpdateAccessKeyComment,

	CreateOrganization: ginCreateOrganization,
	DeleteOrganization: ginDeleteOrganization,
	UpdateOrganization: ginUpdateOrganization,
	GetOrganization:    ginGetOrganization,
	ListOrganization:   ginListOrganization,

	OrgAddMember:                 ginOrgAddMember,
	OrgBindMember:                ginOrgBindMember,
	OrgUnbindMember:              ginOrgUnbindMember,
	GetOrganizationUsers:         ginGetOrganizationUsers,
	SetOrganizationUsersIsmanage: ginSetOrganizationUsersIsmanage,

	CreateDepartment: ginCreateDepartment,
	DeleteDepartment: ginDeleteDepartment,
	UpdateDepartment: ginUpdateDepartment,
	GetDepartment:    ginGetDepartment,
	ListDepartment:   ginListDepartment,

	DepBindMember:              ginDepBindMember,
	DepUnbindMember:            ginDepUnbindMember,
	GetDepartmentUsers:         ginGetDepartmentUsers,
	SetDepartmentUsersIsmanage: ginSetDepartmentUsersIsmanage,
}

func GinRouterClientProxy(action Action, options *nauth.AuthenticationClientOptions) func(*gin.Context) {
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

func ginInitClinet(c *gin.Context, options *nauth.AuthenticationClientOptions) (*nauth.AuthenticationClient, error) {
	// 获取token信息并保存到上下文中
	tokenString := getBearToken(c)
	c.Set(nauth.CurrentUserToken, tokenString)

	// Initial authentication client
	cli, err := nauth.NewClient(options)
	if err != nil {
		return nil, err
	}
	if err = cli.SetToken(c.GetString(nauth.CurrentUserToken)); err != nil {
		return nil, err
	}

	return cli, nil
}
