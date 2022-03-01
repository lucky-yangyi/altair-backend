package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
)

func GetAdmin(email string) (admin model.Admin, err error) {
	err = dao.DB.Model(admin).Where("email = ?", email).First(&admin).Error
	return admin, err
}
