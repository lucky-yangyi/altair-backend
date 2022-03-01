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
	type AdminController struct {
		base
}

// @Summary 登陆
// @Description 用户登陆
// @Produce json
// @Param	body	body	request.AdminLogin	true	"email 邮箱; password 密码"
// @Success 200	{object}  cache.AdminLogin
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
// @router /api/v1/admin/login [post]
func (c *AdminController) Login(g *gin.Context) {
	var param request.AdminLogin
	err := g.BindJSON(&param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}
	var admin model.Admin

	admin, err = service.GetAdmin(param.Email)

	if err == nil {
		if param.Password == admin.Password {
			// 生成accesskey和secretkey
			adminInfo := utils.AdminLoginKey(admin, time.Now().Unix())
			err = cache.SetAdminInRedis(adminInfo)
			c.ResponseData(g, adminInfo, http.Ok)
			return
		} else {
			c.ResponseData(g, nil, http.Forbidden, "密码错误")
			return
		}
	} else {
		c.ResponseData(g, nil, http.BadRequest, "账号或密码错误")
		return
	}

}
