package nauth

type AuthenticationClientOptions struct {
	Host                    string
	Tenant                  string
	Secret                  string
	AppId                   string
	AppSecret               string
	RedirectUri             string
	TokenEndPointAuthMethod string
	Issuer                  string
	UserContext             map[string]interface{}
}

type AuthUrlResult struct {
	Url   string
	State string
	Nonce string
}
