package dao

import (
	"altair-backend/internal/model"
	"gorm.io/gorm"
)

const tableCompanyMember = "company_member"
const tableMemberWallet = "wallet_authorize"

//添加成员
func (d *Dao) CreateCompanyMember(tx *gorm.DB, data *model.CompanyMember) error {
	if tx != nil {
		return tx.Table(tableCompanyMember).Create(&data).Error
	}
	return d.dao.Table(tableCompanyMember).Create(&data).Error
}

//判断钱包是否存在
func (d *Dao) CheckWalletAuthExist(ids []int64, memberId uint64) (exist bool, err error) {
	var count int64
	err = d.dao.Table(tableMemberWallet).Where("wallet_id IN (?) AND company_member_id = ? AND is_del = 0", ids, memberId).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

//判断邮箱是否存在
func (d *Dao) CheckEmailExist(Email string) (exist bool, err error) {
	var count int64
	err = d.dao.Table(tableCompanyMember).Where("email IN (?) AND is_del = 0", Email).Count(&count).Error
	if err != nil {
		return false, err
	}
	return count != 0, nil
}

//判断新密码和旧密码是否一致
func (d *Dao) CheckPassword(memberId uint64) (data *model.CompanyMember, err error) {
	err = d.dao.Table(tableCompanyMember).Where("id", memberId).Find(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil

}

// 增加钱包权限
func (d *Dao) CreateWalletPermission(tx *gorm.DB, data []*model.WalletAuthorize) error {
	for _, v := range data {
		if tx != nil {
			err := tx.Table(tableMemberWallet).Create(v).Error
			if err != nil {
				return err
			}
		} else {
			err := d.dao.Table(tableMemberWallet).Create(v).Error
			if err != nil {
				return err
			}
		}

	}
	return nil
}

//获取钱包权限
func (d *Dao) GetWalletByMemberId(memberId uint64) (list []*model.Wallet, err error) {
	sql := `SELECT * FROM wallet w LEFT JOIN wallet_authorize wa ON w.id = wa.wallet_id WHERE wa.company_member_id = ? AND wa.is_del = 0 AND w.is_del = 0 `
	err = d.dao.Raw(sql, memberId).Find(&list).Error
	return
}

//删除钱包权限
func (d *Dao) DelWalletPermission(memberId uint64) error {
	err := d.dao.Table(tableMemberWallet).Where("company_member_id", memberId).Delete(tableMemberWallet)
	if err != nil {
		return nil
	}
	return nil
}

//成员禁用/启用
func (d *Dao) UpdateMemberEnable(status uint64, memberId uint64) error {
	err := d.dao.Table(tableCompanyMember).Where("id = ?", memberId).UpdateColumn("enabled", status).Error
	return err
}

//获取成员列表
func (d *Dao) GetCompanyMemberList(companyId uint64, offset, size uint64) (list []*model.CompanyMemberNoPassword, err error) {
	db := d.dao.Table(tableCompanyMember)
	err = db.Offset(int(offset)).Limit(int(size)).Where("company_id", companyId).Order("id desc").Find(&list).Error
	return
}

//获取成员列表总数
func (d *Dao) GetCompanyMemberCount(companyId uint64) (count int64, err error) {
	db := d.dao.Table(tableCompanyMember)
	err = db.Where("company_id", companyId).Count(&count).Error
	return
}

//成员修改
func (d *Dao) UpdateCompanyMember(memberId uint64, member *model.CompanyMember) error {
	return d.dao.Table(tableCompanyMember).Where("id", memberId).Updates(&member).Error
}

//成员密码修改
func (d *Dao) ChangePassword(memberId uint64, member *model.CompanyMember) error {
	return d.dao.Table(tableCompanyMember).Where("id", memberId).Updates(&member).Error
}
