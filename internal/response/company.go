package response

import "altair-backend/internal/model"

type CompanyListResp struct {
	List []*model.Company
	Page
}

type MemberListResp struct {
	List []*model.CompanyMemberNoPassword
	Page
}
