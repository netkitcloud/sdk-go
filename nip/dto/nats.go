package dto

// 获取nats用户配置信息返回的结构体
type NatsUser struct {
	Jwt   string `json:"jwt"`
	Seed  string `json:"seed"`
	Creds string `json:"creds"`
}

// 获取nats状态返回的结构体
type NatsStatus struct {
	Activated bool `json:"activated"`
}

// 获取nats订阅主题权限返回的结构体
type NatsPermission struct {
	AllowSubs []string `json:"allowSubs"`
	RmSubs    []string `json:"rmSubs"`
}

// 增加nats订阅主题权限DTO
type AddNatsPermissionDto struct {
	AllowSubs []string `json:"allowSubs" validate:"required"`
}

// 删除nats订阅主题权限DTO
type DeleteNatsPermissionDto struct {
	RmSubs []string `json:"rmSubs" validate:"required"`
}
