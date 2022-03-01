package dao

import "altair-backend/internal/model"

func (d *Dao) GetWalletAuthorizeList(member model.CompanyMember) (ids []int) {
	d.dao.Model(model.WalletAuthorize{}).Where("company_member_id = ?", member.ID).Pluck("wallet_id", &ids)
	return
}
