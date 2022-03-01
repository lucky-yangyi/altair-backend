package errcode

var (
	Success         = NewError(200, "成功")
	InvalidParams   = NewError(402, "入参错误")
	CheckEmailError = NewError(406, "邮箱不合法")
	//ServerError               = NewError(500, "服务器内部错误")

	NotFound        = NewError(10000002, "找不到")
	TooManyRequests = NewError(10000007, "请求过多")

	BadRequest      = NewError(400, "错误请求")
	PermissionError = NewError(404, "没有权限")
	// 业务错误
	WalletUnExistError   = NewError(1000, "钱包不存在")
	WalletAuthExistError = NewError(1001, "存在用户已经存在的钱包权限")
	CheckEmailExistError = NewError(1002, "邮箱已存在")

	AuthError     = NewError(401, "被封禁，拒绝访问")
	PasswordError = NewError(405, "密码错误")
	NoUser        = NewError(403, "用户不存在")
)
