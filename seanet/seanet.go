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
	apiStatusDevice = "/device/status"
	apiModifyDevice = "/device/%s"
)
