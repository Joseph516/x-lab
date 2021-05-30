package errcode

// 公共错误码

var (
	Success                   = NewError(200, "成功")
	ServerError               = NewError(1001, "服务内部错误")
	InvalidParams             = NewError(1002, "入参错误")
	NotFound                  = NewError(1003, "找不到")
	UnauthorizedAuthNotExist  = NewError(1004, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(1005, "鉴权失败，Token错误")
	UnauthorizedTokenTimeout  = NewError(1006, "鉴权失败，Token超时")
	UnauthorizedTokenGenerate = NewError(1007, "鉴权失败，Token生成失败")
	TooManyRequests           = NewError(1008, "请求过多")
)
