package param

import "github.com/netkitcloud/sdk-go/common"

type QueryUsers struct {
	common.PaginationParams
	CreatorUID string `json:"creator_uid" form:"creator_uid" binding:"omitempty"`
}
