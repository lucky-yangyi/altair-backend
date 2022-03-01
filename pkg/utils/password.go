package utils

import (
	"altair-backend/internal/cache"
	"altair-backend/internal/model"

	"fmt"

	"unicode"
)

// 密码强度校验
func VerifyPasswordRule(str string) bool {
	var (
		isUpper   = false
		isLower   = false
		isNumber  = false
		isSpecial = false
	)

	if len(str) < 10 {
		return false
	}

	for _, s := range str {
		switch {
		case unicode.IsUpper(s):
			isUpper = true
		case unicode.IsLower(s):
			isLower = true
		case unicode.IsNumber(s):
			isNumber = true
		case unicode.IsPunct(s) || unicode.IsSymbol(s):
			isSpecial = true
		default:
		}
	}

	if (isUpper && isLower) && (isNumber || isSpecial) {
		return true
	}
	return false
}

// LoginKey 登录返回数据
func LoginKey(user model.CompanyMember, time int64) cache.Login {

	return cache.Login{
		User: model.CompanyMember{
			Mixin:     user.Mixin,
			Name:      user.Name,
			Email:     user.Email,
			Company:   user.Company,
			IsAdmin:   user.IsAdmin,
			CompanyID: user.CompanyID,
		},
		Token: cache.Token{
			AccessKey: GenerateMd5(fmt.Sprint(time) + user.Password),
			SecretKey: GenerateMd5(user.Password + fmt.Sprint(time)),
		},
	}
}

// AdminLoginKey 登录返回数据
func AdminLoginKey(admin model.Admin, time int64) cache.AdminLogin {
	return cache.AdminLogin{
		Admin: model.AdminNoPassword{
			Mixin: admin.Mixin,
			Name:  admin.Name,
			Email: admin.Email,
		},
		Token: cache.Token{
			AccessKey: GenerateMd5(fmt.Sprint(time) + admin.Password),
			SecretKey: GenerateMd5(admin.Password + fmt.Sprint(time)),
		},
	}
}
