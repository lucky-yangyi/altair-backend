package response

import "altair-backend/internal/model"

type MemberWalletAuthResp struct {
	List []*model.Wallet
}

type CompanyMemberListResp struct {
	List []*model.CompanyMemberNoPassword
	Page
}
