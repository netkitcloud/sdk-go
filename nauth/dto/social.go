package dto

import (
	"encoding/json"

	"github.com/netkitcloud/sdk-go/common"
)

// WxUser 微信用户结构定义
type WxUser struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	OpenID    string `json:"openid"`
	UnionID   string `json:"unionid"`
	Phone     string `json:"phone"`
	UID       string `json:"uid"`
	SocialID  string `json:"socialid"`
	UpdatedAt string `json:"updatedAt,omitempty"`
	CreatedAt string `json:"createdAt,omitempty"`
}

type ListWxUserDto struct {
	common.BaseResponse
	Data []WxUser `json:"data,omitempty"`
}

type SocialLoginDto struct {
	// 微信小程序登录时需要传入的字段
	PhoneCode string `json:"phone_code,omitempty"`
	LoginCode string `json:"login_code,omitempty"`
	// 微信公众号授权登陆无需传入字段
	// 微信开放平台登录时无需传入字段
}

type SocialLoginCallbackDto struct {
	// 微信小程序登录无需传入字段
	// 微信公众号授权登陆后循环回调需要传入的字段
	Ticket string `json:"ticket,omitempty"` // 通过唯一的ticket确认是否是同一个用户，获取该用户是否授权成功，成功后返回该用户登陆信息

	// 微信开放平台登录时需要传入的字段
	Code  string `json:"code,omitempty"`  // 通过 code 换取网页授权 access_token.
	State string `json:"state,omitempty"` // 即identifier

}

type SocialLoginRespDto struct {
	common.BaseResponse
	Data struct {
		// 微信小程序授权登陆返回的字段
		User

		// 微信公众号授权登陆返回的字段
		Ticket       string `json:"ticket,omitempty"`
		RedirectURI  string `json:"redirect_uri,omitempty"`  // 前端通过该URI寻回回调地址
		QrcodepicURI string `json:"qrcodepic_uri,omitempty"` // 二维码图片地址

		// 微信开放平台返回的字段
		AuthCodeURL string `json:"auth_code_url,omitempty"` // 通过该URL获取用户授权code
	} `json:"data,omitempty"`
}

type SocialBindDto struct {
	// 微信小程序绑定时需要传入的字段
	PhoneCode string `json:"phone_code,omitempty"`
	LoginCode string `json:"login_code,omitempty"`
	// 微信公众号绑定时无需传入字段
	// TODO 微信开放平台
}

type SocialBindRespDto struct {
	common.BaseResponse
	Data struct {
		// 微信公众号绑定返回的字段
		AuthCodeURL string `json:"auth_code_url,omitempty"` // 通过该URL获取用户授权code
	} `json:"data,omitempty"`
}

type SocialUnBindDto struct {
	// 微信小程序绑定时无需传入字段
	// 微信公众号绑定时无需传入字段
	OpenID  string `json:"openid,omitempty"`
	UnionID string `json:"unionid,omitempty"`
	// TODO 微信开放平台
}

// 微信公众号模板消息
type SocialWxOfficeTemplateMsgDto struct {
	ToUser      string          `json:"touser"`                // 必须, 接受者OpenID
	TemplateId  string          `json:"template_id"`           // 必须, 模版ID
	URL         string          `json:"url,omitempty"`         // 可选, 用户点击后跳转的URL, 该URL必须处于开发者在公众平台网站中设置的域中
	MiniProgram *MiniProgram    `json:"miniprogram,omitempty"` // 可选, 跳小程序所需数据，不需跳小程序可不用传该数据
	Data        json.RawMessage `json:"data"`                  // 必须, 模板数据, JSON 格式的 []byte, 满足特定的模板需求
}

type MiniProgram struct {
	AppId    string `json:"appid"`    // 必选; 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
	PagePath string `json:"pagepath"` // 必选; 所需跳转到的小程序appid（该小程序appid必须与发模板消息的公众号是绑定关联关系）
}

// WxUserQueryDto 微信用户查询表单
type WxUserQueryDto struct {
	Phone   string `json:"phone" binding:"omitempty"`
	OpenID  string `json:"openid" binding:"omitempty"`
	UnionID string `json:"unionid" binding:"omitempty"`
}
