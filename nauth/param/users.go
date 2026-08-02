package param

import "github.com/netkitcloud/sdk-go/common"

type QueryUsers struct {
	common.PaginationParams
	CreatorUID string `json:"creator_uid" form:"creator_uid" binding:"omitempty"`
	// Username 用户名模糊搜索关键词。保留用于向后兼容。
	Username string `json:"username,omitempty" form:"username" binding:"omitempty"`
	// Search 用户名、邮箱、手机号通用模糊搜索关键词。
	Search string `json:"search,omitempty" form:"search" binding:"omitempty"`
}
