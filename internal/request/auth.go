package request

type Login struct {
	Email    string `json:"email" binding:"required,email"` //邮箱
	Password string `json:"password" binding:"required"`    //密码
}

type AdminLogin struct {
	Email    string `json:"email" binding:"required"`    //邮箱
	Password string `json:"password" binding:"required"` //密码
}
