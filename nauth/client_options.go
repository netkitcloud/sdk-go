package nauth

type AuthenticationClientOptions struct {
	Host                    string
	Tenant                  string
	AccessKey               string
	Secret                  string
	AppId                   string
	AppSecret               string
	RedirectUri             string
	TokenEndPointAuthMethod string
	Issuer                  string
}

type AuthUrlResult struct {
	Url   string
	State string
	Nonce string
}
