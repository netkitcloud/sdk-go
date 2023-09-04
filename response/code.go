package response

var (
	Success = NewError(0, "成功")

	// 通用错误码
	Base64EncodeError  = NewError(105, "base64编码错误")
	Base64DecodeError  = NewError(106, "base64解码错误")
	JsonMarshalError   = NewError(107, "json编码错误")
	JsonUnmarshalError = NewError(108, "json解码错误")
	CreateError        = NewError(109, "创建错误")
	UpdateError        = NewError(110, "更新错误")
	GetError           = NewError(111, "查询错误")
	DeleteError        = NewError(112, "删除错误")
	FromError          = NewError(113, "表单错误")
	TransformError     = NewError(114, "数据转换错误")
	ListError          = NewError(115, "获取列表错误")


	// 初始化客户端失败
	InitClientError = NewError(201, "init client error")
)
