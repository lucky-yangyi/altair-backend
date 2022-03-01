package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"context"
	"github.com/gin-gonic/gin"
)

type baseService struct {
	ctx context.Context
	dao *dao.Dao
}

func (base *baseService) IsAdmin(memberInfo model.CompanyMember) bool {
	return memberInfo.IsAdmin
}

func GetUserInfo(c *gin.Context) model.CompanyMember {
	var user model.CompanyMember
	v, ok := c.Get("UserInfo")
	if ok {
		user, ok = v.(model.CompanyMember)
		if ok {
			return user
		}
	}
	return user
}

//GetUserAuthIds 用户的钱包权限列表
func GetUserAuthIds(c *gin.Context) (wids []int, err bool) {
	userInput := GetUserInfo(c)
	wids = GetWalletAuthorizeList(userInput)
	return wids, true
}
