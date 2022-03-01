package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
)

func GetWalletAuthorizeList(userInput model.CompanyMember) (ids []int) {
	if userInput.IsAdmin == true { // 管理员拥有组织下所有钱包权限
		dao.DB.Table(new(model.Wallet).TableName()).Where("company_id = ?", userInput.CompanyID).Pluck("id", &ids)
		return
	} else { // 非管理员只有在权限列表下的钱包
		return dao.GlobalDao.GetWalletAuthorizeList(userInput)
	}
}
