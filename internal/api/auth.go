package api

import (
	"altair-backend/internal/cache"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"github.com/gin-gonic/gin"
	"time"
)

// 登录验证
type AuthController struct {
	base
}

// @Summary 登陆
// @Description 用户登陆
// @Produce json
// @Param	body	body	request.Login	true	"email 邮箱; password 密码"
// @Success 200	{object}  cache.Login
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
// @router /api/v1/user/login [post]
func (c *AuthController) Login(g *gin.Context) {
	var param request.Login
	err := g.BindJSON(&param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}

	var user model.CompanyMember

	user, err = service.GetMember(param.Email)
	if err == nil {
		if user.Enabled == 0 {
			c.ResponseData(g, nil, http.Forbidden, "用户被封禁") //c.ResponseData(nil, http.Forbidden)
			return
		}

		if param.Password == user.Password {
			// 生成accesskey和secretkey

			userInfo := utils.LoginKey(user, time.Now().Unix())
			err = cache.SetUserInRedis(userInfo)
			c.ResponseData(g, userInfo, http.Ok)
			return
		} else {
			c.ResponseData(g, nil, http.BadRequest, "密码错误")
			return
		}
	} else {
		c.ResponseData(g, nil, http.BadRequest, "账号或密码错误")
		return
	}

}
