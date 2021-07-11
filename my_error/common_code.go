package my_error

var (
	Success                   = NewErrorCode(0, "成功")
	ServerError               = NewErrorCode(10000000, "服务内部错误")
	InvalidParams             = NewErrorCode(10000001, "入参错误")
	NotFound                  = NewErrorCode(10000002, "找不到数据，请刷新后重试")
	UnauthorizedAuthNotExist  = NewErrorCode(10000003, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewErrorCode(10000004, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewErrorCode(10000005, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewErrorCode(10000006, "鉴权失败，Token生成失败")
	TooManyRequests           = NewErrorCode(10000007, "请求过多")
	UnauthorizedApi           = NewErrorCode(10000008, "鉴权失败，无当前接口访问权限")
	DBError                   = NewErrorCode(10000009, "服务内部错误") // 数据库操作失败
	ExtApiError               = NewErrorCode(10000010, "服务内部错误") // 调用外部api接口报错
	BusinessError             = NewErrorCode(10000011, "服务内部错误") // 业务异常
)
