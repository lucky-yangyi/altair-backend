package http

// Status 状态
type Status struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

var (
	Ok                   = Status{200, "成功"}
	BadRequest           = Status{400, "错误请求"}
	DataFieldInvalid     = Status{400, "数据字段不合法"}
	DataNotFound         = Status{400, "数据不存在"}
	DataExisted          = Status{400, "数据已存在"}
	Unauthorized         = Status{401, "认证失败"}
	Forbidden            = Status{403, "没有权限访问"}
	PageNotFound         = Status{404, "请求地址不存在"}
	MethodNotAllowed     = Status{405, "请求方法不允许"}
	TooManyRequests      = Status{426, "请求频率限制"}
	InternalServerError  = Status{500, "服务器内部错误"}
	CheckEmailExistError = Status{400, "服务器内部错误"}
)
