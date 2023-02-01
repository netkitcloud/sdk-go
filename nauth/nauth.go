package nauth

const (
	SdkName         string = "nauth-golang-sdk"
	Version         string = "0.0.1"
	CoreAuthApiHost string = "https://nauth.netkit.cloud/api"
)

var commonHeaders = map[string]string{
	"x-nauth-request-from": SdkName,
	"x-nauth-sdk-version":  Version,
}

type User struct {
	UID       string `json:"uid,omitempty"`
	IsSuper   bool   `json:"issuper,omitempty"`
	UserName  string `json:"username,omitempty"`
	Nickname  string `json:"nickname,omitempty"`
	Email     string `json:"email,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Language  string `json:"language,omitempty"`
	UpdatedAt string `json:"updatedat,omitempty"`
	Token     string `json:"token,omitempty"`

	Birthdate string `json:"birthdate,omitempty"`
	Gender    string `json:"gender,omitempty"` //'M' | 'F' | 'U'
	Picture   string `json:"picture,omitempty"`
	Zoneinfo  string `json:"zoneinfo,omitempty"`
	Locale    string `json:"locale,omitempty"`
}
