package dto

type PhoneCodeLoginDto struct {
	Phone string `json:"phone"`
	Code  int    `json:"code"`
}

type PhonePasswordLoginDto struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}

type UsernameLoginDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UsernameLoginResult struct {
	User
	AccessTokenClaims
}
