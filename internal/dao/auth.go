package dao

import (
	"altair-backend/internal/model"
)

func (d *Dao) GetMember(email string) (user model.CompanyMember, err error) {
	err = d.dao.Table(user.TableName()).Where("email = ?", email).Preload("Company").First(&user).Error
	return
}

func (d *Dao) GetAdmin(phone string) (user model.Admin, err error) {
	err = d.dao.Table("admin").Where("mobile = ?", phone).First(&user).Error
	return
}
