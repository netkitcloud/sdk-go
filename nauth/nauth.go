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

type UserProfile struct {
	Nickname  string `json:"nickname,omitempty"`
	City      string `json:"city,omitempty"`
	Country   string `json:"country,omitempty"`
	Language  string `json:"language,omitempty"`
	Birthdate string `json:"birthdate,omitempty"`
	Gender    string `json:"gender,omitempty"` //'M' | 'F' | 'U'
	Picture   string `json:"picture,omitempty"`
	Timezone  string `json:"timezone,omitempty"`
	Locale    string `json:"locale,omitempty"`
}

type User struct {
	UID      string `json:"uid,omitempty"`
	IsSuper  bool   `json:"issuper,omitempty"`
	UserName string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	AccessToken    string `json:"access_token,omitempty"`
	UpdatedAt string `json:"updatedat,omitempty"`
	UserProfile
}
