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
