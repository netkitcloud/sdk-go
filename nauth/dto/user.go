package dto

import "github.com/netkitcloud/sdk-go/common"

type UpdateUserProfileDto struct {
	UserProfile
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
	UnionId   string `json:"unionId,omitempty"`
	OpenId    string `json:"openId,omitempty"`
}

type User struct {
	UID         string   `json:"uid,omitempty"`
	IsSuper     bool     `json:"issuper,omitempty"`
	UserName    string   `json:"username,omitempty"`
	Email       string   `json:"email,omitempty"`
	Phone       string   `json:"phone,omitempty"`
	AccessToken string   `json:"access_token,omitempty"`
	UpdatedAt   string   `json:"updatedAt,omitempty"`
	CrateddAt   string   `json:"createdAt,omitempty"`
	WxUsers     []WxUser `json:"wxUsers,omitempty"`
	UserProfile
}

type AddUserDto struct {
	UserName string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Phone    string `json:"phone,omitempty"`
	Password string `json:"password,omitempty"`
	UserProfile
}

type UpdateUserDto struct {
	UserName  string `json:"username,omitempty"`
	Email     string `json:"email,omitempty"`
	Phone     string `json:"phone,omitempty"`
	Password  string `json:"password,omitempty"`
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

type GetUserResponseDto struct {
	common.BaseResponse
	Data User
}

type UpdatePasswordDto struct {
	Password    string `json:"password,omitempty"`
	NewPassword string `json:"new_password,omitempty"`
	RePassword  string `json:"re_password,omitempty"`
}
