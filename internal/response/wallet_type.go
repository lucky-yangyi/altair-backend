package response

import "altair-backend/internal/model"

type WalletTypeListResp struct {
	List []*model.WalletType
	Page
}
