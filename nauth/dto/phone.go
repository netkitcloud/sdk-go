package dto

type PhoneCodeRequestDtto struct {
	Phone string `json:"phone"`
}

// UpdatePhoneDto 用户修改手机表单
type UpdatePhoneDto struct {
	Phone    string `form:"phone" json:"phone"`
	Code     int    `form:"code" json:"code" binding:"min=1111,max=9999"`
	NewPhone string `form:"new_phone" json:"new_phone" binding:"required"`
	NewCode  int    `form:"new_code" json:"new_code" binding:"required,min=1111,max=9999"`
}