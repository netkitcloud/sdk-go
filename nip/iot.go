package nip

var commonHeaders = map[string]string{
	"x-nip-request-from": SdkName,
	"x-nip-sdk-version":  Version,
}

const (
	SdkName         string = "nip-golang-sdk"
	Version         string = "0.1.0"
	CoreAuthApiHost string = "https://iot.netkit.cloud/api"
	// CoreAuthApiHost string = "http://localhost:3000"
)

const (
	apiProduct       = "/products"
	apiModifyProduct = "/products/%s"

	apiDevice       = "/devices"
	apiModifyDevice = "/devices/%s"

	apiNatsUser       = "/nats"
	apiGetNatsStatus  = "/nats/status"
	apiNatsPermission = "/nats/permission"

	apiRuleForward       = "/rules/forwards"
	apiModifyRuleForward = "/rules/forwards/%s"
)
