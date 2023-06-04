package dto

type OIDCTokenResponse struct {
	AccessToken      string `json:"access_token,omitempty"`
	IDToken          string `json:"id_token,omitempty"`
	RefreshToken     string `json:"refresh_token,omitempty"` //可选
	ExpiresIn        uint64 `json:"expires_in,omitempty"`
	TokenType        string `json:"token_type,omitempty"`
	Error            string `json:"error,omitempty"`
	ErrorDescription string `json:"error_description,omitempty"`
}
