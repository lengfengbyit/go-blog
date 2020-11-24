package errcode

var (
	Success                   = NewError(0, "成功")
	ServerError               = NewError(500, "服务内部错误")
	InvalidParams             = NewError(400, "入参错误")
	NotFound                  = NewError(404, "找不到")
	UnauthorizedAuthNotExist  = NewError(411, "鉴权失败，找不到对应的AppKey和AppSecret")
	UnauthorizedTokenError    = NewError(412, "鉴权失败, Token 错误")
	UnauthorizedTokenTimeout  = NewError(413, "鉴权失败, Token超时")
	UnauthorizedTokenNotExist = NewError(414, "鉴权失败,Token不存在")
	UnauthorizedTokenGenerate = NewError(415, "鉴权失败, Token生成失败")
	TooManyRequests           = NewError(429, "请求过多")
)

