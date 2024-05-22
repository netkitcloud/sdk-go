package nauth

const (
	SdkName         string = "nauth-golang-sdk"
	Version         string = "0.1.0"
	CoreAuthApiHost string = "https://nauth.netkit.cloud/api"
)

var commonHeaders = map[string]string{
	"x-nauth-request-from": SdkName,
	"x-nauth-sdk-version":  Version,
}

// api uri
const (
	apiRole        = "/role"
	apiSpecialRole = "/role/%s"
	apiUserRole    = "/role/user"

	apiResource              = "/resource"
	apiSpecialResource       = "/resource/%s"
	apiSpecialResourceAction = "/resource/%s/action"
	apiAuthorizeResource     = "/resource/authorize"
	apiUnauthorizeResource   = "/resource/unauthorize"

	apiAction        = "/resource/action"
	apiSpecialAction = "/resource/action/%s"
	apiVerifyAction  = "/resource/action/verify"

	apiOrganization                    = "/organization"
	apiSpecialOrganization             = "/organization/%s"
	apiSpecialOrganizationMember       = "/organization/%s/member"
	apiSpecialOrganizationMemberBind   = "/organization/%s/member/bind"
	apiSpecialOrganizationMemberUnbind = "/organization/%s/member/unbind"

	apiDepartment        = "/department"
	apiSpecialDepartment = "/department/%s"
)
