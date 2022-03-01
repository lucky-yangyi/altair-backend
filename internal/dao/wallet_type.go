package dao

import (
	"altair-backend/internal/model"
)

const tableWalletType = "wallet_type"

//获取币种列表
func (d *Dao) GetWalletTypeList(name, code string, offset, size uint64) (list []*model.WalletType, err error) {
	db := d.dao
	db = db.Table(tableWalletType)
	if name != "" {
		db = db.Where(" name = ?", name)
	}
	if code != "" {
		db = db.Where("code = ?", code)
	}
	err = db.Where("is_del = ?", 0).Offset(int(offset)).Limit(int(size)).Find(&list).Error
	return

}

//获取获取币种列表总数
func (d *Dao) CountWalletTypeList(name, code string) (count int64, err error) {
	db := d.dao
	db = db.Table(tableWalletType)

	if name != "" {
		db = db.Where(" name = ?", name)
	}
	if name != "" {
		db = db.Where("code = ?", code)
	}
	err = db.Where("is_del = ?", 0).Count(&count).Error
	return
}
