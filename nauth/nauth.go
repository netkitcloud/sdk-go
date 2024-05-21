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

	apiResource            = "/resource"
	apiSpecialResource     = "/resource/%s"
	apiAuthorizeResource   = apiResource + "/authorize"
	apiUnauthorizeResource = apiResource + "/unauthorize"

	apiAction        = apiResource + "/action"
	apiSpecialAction = apiResource + "/action/%s"
	apiVerifyAction  = apiAction + "/verify"

	apiOrganization                    = "/organization"
	apiSpecialOrganization             = "/organization/%s"
	apiSpecialOrganizationMember       = apiSpecialOrganization + "/member"
	apiSpecialOrganizationMemberBind   = apiSpecialOrganizationMember + "/bind"
	apiSpecialOrganizationMemberUnbind = apiSpecialOrganizationMember + "/unbind"

	apiDepartment        = "/department"
	apiSpecialDepartment = "/department/%s"
)
