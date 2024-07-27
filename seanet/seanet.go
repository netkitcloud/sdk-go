package seanet

const (
	SdkName         string = "seanet-golang-sdk"
	Version         string = "0.1.0"
	CoreAuthApiHost string = "https://cloud.sealan.tech/api"
	// CoreAuthApiHost string = "http://192.168.15.160:18000"
)

var commonHeaders = map[string]string{
	"x-seanet-request-from": SdkName,
	"x-seanet-sdk-version":  Version,
}

const (
	apiProduct       = "/product"
	apiModifyProduct = "/product/%s"

	apiDevice       = "/device"
	apiModifyDevice = "/device/%s"

	// gateway
	apiLogDevice    = "/device/%s/log"
	apiSubDevices   = "/device/%s/gateway"
	apiStatusDevice = "/device/status"
	apiCmdDevice    = "/device/cmd"

	// gateway rule
	apiGatewayRule        = "/gateway/rule"
	apiGatewayRuleDevices = "/gateway/rule/%s/device"

	// v2
	apiDeviceOperateProperty = "/device/operate/property"
	apiDeviceOperateCmd      = "/device/operate/cmd"

	// property
	apiDeviceProperty = "/device/%s/getproperty/%s"
)
