package request

type CreateCompanyMemberReq struct {
	MemberName string `json:"memberName" binding:"required,min=3,max=20"` //名称
	MemberPwd  string `json:"memberPwd" binding:"required,gte=8"`         //密码
	Email      string `json:"email" binding:"required,email"`             //邮箱
	Desc       string `json:"desc"`                                       //备注
}

type AddWalletPermissionReq struct {
	Ids      []int64 `json:"ids" binding:"required"`      //ids
	MemberId uint64  `json:"memberId" binding:"required"` //成员ID
}

type EnableMemberReq struct {
	MemberId uint64 `json:"memberId" binding:"required"` //成员ID
	Status   uint64 `json:"status" binding:"oneof=0 1"`  //状态
}

type CompanyMemberListReq struct {
	PageReq
}

type GetWalletPermissionReq struct {
	MemberId uint64 `json:"memberId" binding:"required"` //成员ID
}

type UpdateCompanyMemberReq struct {
	MemberId   uint64 `json:"memberId" binding:"required"`    //成员ID
	MemberName string `json:"memberName" binding:"required"`  //名称
	Email      string `json:"email" binding:"required,email"` //邮箱
	Desc       string `json:"desc"`                           //备注
}

type ChangePassword struct {
	MemberId  uint64 `json:"memberId" binding:"required"`        //成员ID
	MemberPwd string `json:"memberPwd" binding:"required,gte=8"` //密码
}
