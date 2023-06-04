package dto

type PhoneCodeRegisterDto struct {
	Phone string `json:"phone"`
	Code  int    `json:"code"`
}

type UsernameRegisterDto struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
