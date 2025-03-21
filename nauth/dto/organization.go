package dto

import "github.com/netkitcloud/sdk-go/common"

// 组织模型
type Organization struct {
	ID          uint   `json:"id"`
	CreatorUID  string `json:"creator_uid"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Country     string `json:"country"`
	City        string `json:"city"`
	Category    string `json:"category"`
	Description string `json:"description"`
	Createdat   string `json:"createdat"`
	Updatedat   string `json:"updatedat"`

	Departments []Department `json:"departments"`
	Users       []User       `json:"users"`
}

type OrganizationDto struct {
	Data Organization
	common.BaseResponse
}

type AddOrganizationMemberDto struct {
	Data User
	common.BaseResponse
}

// 获取指定组织下所有用户信息
type OrganizationUserDto struct {
	Data []User
	common.BaseResponse
}

type ListOrganizationDto struct {
	Data []Organization
	common.BaseListResponse
}

// 组织用户模型
type OrganizationUser struct {
	ID             uint   `json:"id"`
	OrganizationID uint   `json:"organization_id"`
	UserUID        string `json:"user_uid"`
	IsManager      bool   `json:"is_manager"`
	Createdat      string `json:"createdat"`
	Updatedat      string `json:"updatedat"`

	User         User         `json:"user"`
	Organization Organization `json:"organization"`
}

type ListOrganizationUserDto struct {
	Data []OrganizationUser
	common.BaseListResponse
}
