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
}

type AuthUrlResult struct {
	Url   string
	State string
	Nonce string
}
