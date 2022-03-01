package dao

import (
	"altair-backend/internal/cache"
	"altair-backend/internal/model"
	"gorm.io/gorm"
)

const tableCompany = "company"

//新增企业列表
func (d *Dao) CreateCompany(company *model.Company) (err error) {
	return d.dao.Table(tableCompany).Create(&company).Error

}

//获取企业列表
func (d *Dao) GetCompanyList(params string, offset, size uint64) (list []*model.Company, err error) {
	db := d.dao.Table(tableCompany)
	if params != "" {
		db = db.Where("name = ?", params)
	}
	err = db.Offset(int(offset)).Limit(int(size)).Find(&list).Error
	return
}

//获取企业列表总数
func (d *Dao) GetCompanyCount(params string) (count int64, err error) {
	db := d.dao.Table(tableCompany)
	if params != "" {
		db = db.Where("name = ?", params)
	}
	err = db.Count(&count).Error
	return
}

//修改企业信息
func (d *Dao) UpdateCompany(CompanyId uint64, company *model.Company) error {
	return d.dao.Table(tableCompany).Where("id", CompanyId).Updates(&company).Error
}

//企业禁用/启用
func (d *Dao) UpdateCompanyEnable(status uint64, companyId uint64) error {
	err := cache.DelCompanyInRedis(companyId)
	err = d.dao.Table(tableCompany).Where("id = ?", companyId).UpdateColumn("enabled", status).Error
	return err
}

//添加企业成员
func (d *Dao) AddCompanyMember(data *model.CompanyMember) error {
	return d.dao.Table(tableCompanyMember).Create(&data).Error
}

//获取成员列表
func (d *Dao) GetMemberList(companyId uint64, offset, size uint64) (list []*model.CompanyMemberNoPassword, err error) {
	db := d.dao.Table(tableCompanyMember)
	err = db.Offset(int(offset)).Limit(int(size)).Where("company_id", companyId).Order("id desc").Find(&list).Error
	return
}

//获取成员列表总数
func (d *Dao) GetMemberCount(companyId uint64) (count int64, err error) {
	db := d.dao.Table(tableCompanyMember)
	err = db.Where("company_id", companyId).Count(&count).Error
	return
}

//设置管理员
func (d *Dao) UpdateMemberAdmin(status uint64, memberId uint64) error {
	err := d.dao.Table(tableCompanyMember).Where("id = ?", memberId).UpdateColumn("is_admin", status).Error
	return err
}

//添加成员
func (d *Dao) CreateMember(tx *gorm.DB, data *model.CompanyMember) error {
	if tx != nil {
		return tx.Table(tableCompanyMember).Create(&data).Error
	}
	return d.dao.Table(tableCompanyMember).Create(&data).Error
}
