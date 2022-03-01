package response

import "altair-backend/internal/model"

type MonthListResp struct {
	List []*model.MixMonthBillWallet
	Page
}

type CreateListResp struct {
	List []*model.Wallet
}
