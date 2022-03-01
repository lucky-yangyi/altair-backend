package dao

import "altair-backend/internal/model"

const tableMonth = "month_bill"

//获取月付账单
func (d *Dao) GetMonthBillCount(walletId, payStatus int, month string) (count int64, err error) {
	var pars []interface{}
	sql := `select count(1) FROM month_bill mb LEFT JOIN wallet w ON mb.wallet_id = w.id WHERE 1 =1 `
	if walletId != 0 {
		sql += " AND w.id = ?"
		pars = append(pars, walletId)
	}
	if month != "" {
		sql += " AND mb.month = ?"
		pars = append(pars, month)
	}
	if payStatus != 0 {
		sql += " AND mb.pay_status = ?"
		pars = append(pars, payStatus)
	}
	if pars != nil {
		err = d.dao.Raw(sql, pars...).Count(&count).Error
	} else {
		err = d.dao.Raw(sql).Count(&count).Error
	}
	return
}

//获取月付账单
func (d *Dao) GetMonthBillList(walletId, payStatus int, month string, offset, size uint64) (list []*model.MixMonthBillWallet, err error) {
	var pars []interface{}
	sql := `select * FROM month_bill mb LEFT JOIN wallet w ON mb.wallet_id = w.id WHERE 1 =1 `
	if walletId != 0 {
		sql += " AND w.id = ?"
		pars = append(pars, walletId)
	}
	if month != "" {
		sql += " AND mb.month = ?"
		pars = append(pars, month)
	}
	if payStatus != 0 {
		sql += " AND mb.pay_status = ?"
		pars = append(pars, payStatus)
	}

	sql += " limit ?,? "
	pars = append(pars, offset, size)
	err = d.dao.Raw(sql, pars...).Scan(&list).Error
	return
}

//获取成员钱包的账单
func (d *Dao) GetWalletList(companyId uint64) (list []*model.Wallet, err error) {
	sql := `select name, address from wallet WHERE 1= 1`
	err = d.dao.Raw(sql, companyId).Scan(&list).Error
	return
}

//新增月账单
func (d *Dao) CreateMonthBill(bill *model.MonthBill) (err error) {
	return d.dao.Table(tableMonth).Create(&bill).Error

}
