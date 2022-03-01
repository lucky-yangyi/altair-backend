package dao

import (
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/pkg/utils"
	"fmt"
	"gorm.io/gorm/clause"
)

const (
	//MultiSignWalletStatusDoing 待签名
	MultiSignWalletStatusDoing uint8 = 1
	//MultiSignWalletStatusFailed 失败
	MultiSignWalletStatusFailed uint8 = 4
	//MultiSignWalletStatusSuccess 成功
	MultiSignWalletStatusSuccess uint8 = 127
	CollectYes                   uint8 = 1
	CollectNo                    uint8 = 2
)

func (d *Dao) WalletListNoPage(user model.CompanyMember) (data []model.WalletAuth, err error) {
	err = d.dao.Table(new(model.Wallet).TableName()).Where("company_id = ?", user.CompanyID).Order("id DESC").Find(&data).Error
	var wids []int
	DB.Table(new(model.WalletAuthorize).TableName()).Where("company_member_id = ?", user.ID).Pluck("wallet_id", &wids)
	for k, v := range data {
		data[k].IsHasAuth = utils.IsInArray(int(v.ID), wids)
	}
	return data, err
}

func (d *Dao) WalletList(ids []int, user model.CompanyMember, pageNo, pageSize uint64, status uint8) (data utils.Page, err error) {
	var wallets []model.WalletMix
	query := d.dao.Table(new(model.Wallet).TableName()).Where("id IN ?", ids)
	fmt.Println("sssssss:", user.IsAdmin)
	data, query, err = utils.Paginate(query, pageNo, pageSize, &wallets)
	if status != 0 {
		query.Where("status = ?", status)
	}
	if query != nil {
		query.Where("is_del = ?", 0).Preload(clause.Associations).Preload("MultiSignWalletDetail.Wallet").Order("id DESC").Find(&wallets)
	}

	return data, err
}

func (d *Dao) WalletUpdate(update request.WalletUpdate) (err error) {
	err = d.dao.Table(new(model.Wallet).TableName()).Where("id = ?", update.WalletID).Update("status", update.Status).Update("address", update.Address).Error
	return
}

func (d *Dao) WalletAdd(wallet model.Wallet, userId uint64) (walletId uint64, err error) {
	err = d.dao.Table(wallet.TableName()).Create(&wallet).Error
	// 添加对应的权限
	var wa model.WalletAuthorize
	wa.WalletID = wallet.ID
	wa.CompanyMemberID = userId
	d.dao.Table(wa.TableName()).Create(&wa)
	return wallet.ID, err
}

func (d *Dao) FilWalletMetaAdd(filWalletMeta model.FilWalletMeta) (walletId uint64, err error) {
	err = d.dao.Table(filWalletMeta.TableName()).Create(&filWalletMeta).Error
	return filWalletMeta.ID, err
}

func (d *Dao) MultiSignWalletDetailAdd(wd model.MultiSignWalletDetail) (err error) {
	err = d.dao.Table(wd.TableName()).Create(&wd).Error
	return err
}

func (d *Dao) WalletMemberCount(memberId uint64) (total int64, normalCount int64, collectCount int64, err error) {
	var ids []int //拥有权限的钱包id
	err = d.dao.Table(new(model.WalletAuthorize).TableName()).Where("company_member_id = ?", memberId).Pluck("wallet_id", &ids).Error
	if err != nil {
		return
	}
	var wallet model.Wallet
	err = d.dao.Table(wallet.TableName()).Where("id IN ?", ids).Count(&total).Error
	if err != nil {
		return
	}
	err = d.dao.Table(wallet.TableName()).Where("id IN ?", ids).Where("type_id = 2").Count(&normalCount).Error
	if err != nil {
		return
	}
	var wc model.WalletCollect
	err = d.dao.Table(wc.TableName()).Where("wallet_id IN ?", ids).Count(&collectCount).Error
	if err != nil {
		return
	}
	return
}

func (d *Dao) WalletAdminCount(companyId uint64) (total int64, normalCount int64, collectCount int64, err error) {
	var wallet model.Wallet
	err = d.dao.Table(wallet.TableName()).Where("company_id = ?", companyId).Count(&total).Error
	if err != nil {
		return
	}
	err = d.dao.Table(wallet.TableName()).Where("company_id = ?", companyId).Where("type_id = 2").Count(&normalCount).Error
	if err != nil {
		return
	}
	var wc model.WalletCollect
	var ids []int //拥有权限的钱包id
	err = d.dao.Table(wallet.TableName()).Where("company_id = ?", companyId).Pluck("id", &ids).Error
	err = d.dao.Table(wc.TableName()).Where("wallet_id IN ?", ids).Count(&collectCount).Error
	if err != nil {
		return
	}
	return
}

func (d *Dao) WalletIsExist(address string) (exist bool) {
	var num int64
	var wallet model.Wallet
	d.dao.Table(wallet.TableName()).Where("address = ?", address).Count(&num)
	if num > 0 {
		return true
	} else {
		return false
	}
}

func (d *Dao) CheckWalletExist(ids []int64) (exist bool) {
	var count int64
	count = 0
	err := d.dao.Table(new(model.Wallet).TableName()).Where("id IN (?)  ", ids).Count(&count).Error
	if err != nil {
		return false
	}
	return count == int64(len(ids))
}
