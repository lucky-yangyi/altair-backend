package request

type CreateCompanyReq struct {
	CompanyName string `json:"companyName" binding:"required,min=3,max=20"` //企业名称
	//Status      uint8 `json:"status" binding:"oneof=0 1"`
	//CompanyPwd  string `alias:"密码" valid:"Required" json:"company_pwd"`
	//Email       string `alias:"邮箱" json:"email" valid:"Required"`
	//Desc        string `alias:"备注" json:"desc"`
}

type CompanyListReq struct {
	Params string `json:"params"` //邮箱 企业名称
	PageReq
}

type UpdateCompanyReq struct {
	CompanyId   uint64 `json:"companyId" binding:"required"`                //企业ID
	CompanyName string `json:"companyName" binding:"required,min=3,max=10"` //企业名称
}

type EnableCompanyReq struct {
	CompanyId uint64 `json:"companyId"`                  //企业ID
	Status    uint64 `json:"status" binding:"oneof=0 1"` //禁用 停用
}

type AddCompanyMemberReq struct {
	CompanyID  uint64 `json:"companyId" binding:"required"`               //企业ID
	MemberName string `json:"memberName" binding:"required,min=3,max=10"` //姓名
	Email      string `json:"email" binding:"required,email"`             //邮箱
	MemberPwd  string `json:"memberPwd" binding:"required,gte=8"`         //密码
	Desc       string `json:"desc"`                                       //备注
}

type MemberListReq struct {
	CompanyId uint64 `json:"companyId" binding:"required"` //成员ID
	PageReq
}

type AdminMemberReq struct {
	MemberId uint64 `json:"memberId" binding:"required"` //成员ID
	Status   uint64 `json:"status" binding:"oneof=0 1"`  // 0 普通 1 管理员
}

type CreateMemberReq struct {
	CompanyID  uint64 `json:"companyId" binding:"required"`               // 企业ID
	MemberName string `json:"memberName" binding:"required,min=1,max=10"` //名称
	MemberPwd  string `json:"memberPwd" binding:"required,gte=8"`         //密码
	Email      string `json:"email" binding:"required,email"`             //邮箱
	Desc       string `json:"desc" binding:"max=100"`                     //备注
}
