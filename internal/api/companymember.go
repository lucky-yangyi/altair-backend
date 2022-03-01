package api

import (
	"altair-backend/internal/cache"
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/log"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var GloCompanyMemberController *CompanyMemberController

type CompanyMemberController struct {
	base
	srv *service.CompayMemberService
}

func newCompanyMemberController(srv *service.CompayMemberService) *CompanyMemberController {
	return &CompanyMemberController{
		srv: srv,
	}
}

//初始化
func GetCompanyMemberController() *CompanyMemberController {
	GloCompanyMemberController = newCompanyMemberController(service.GetCompanyMemberService())
	return GloCompanyMemberController
}

// @Summary 获取成员列表
// @Produce json
// @Param	body	body	request.CompanyMemberListReq	true	"body参数"
// @Success 200 {object} response.CompanyMemberListResp "成功"
//@Failure 500  "内部错误"
// @Router /api/v1/member/list [post]
func (controller *CompanyMemberController) List(c *gin.Context) {
	var req request.CompanyMemberListReq
	if !controller.IsAdminS(c) {
		controller.ResponseData(c, nil, http.Forbidden)
		return
	}
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	// 获取companyId
	user := controller.getUserInfo(c)
	if user.ID == 0 {
		log.Fatal("获取用户信息失败", user)
		controller.ResponseData(c, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}
	page, size := utils.CheckIsPage(req.PageNo, req.PageSize)
	req.PageSize = size
	req.PageNo = page
	data, err := controller.srv.GetCompanyMemberList(c, user.CompanyID, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, data, http.Ok)
	return
}

// @Summary 增加成员
// @Produce json
// @Param	body	body	request.CreateCompanyMemberReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
// @Failure 403 "没有权限访问"
//@Failure 500  "内部错误"
// @Router /api/v1/member/create [post]
func (controller *CompanyMemberController) Create(c *gin.Context) {
	var req request.CreateCompanyMemberReq
	if !controller.IsAdminS(c) {
		controller.ResponseData(c, nil, http.Forbidden)
		return
	}
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	// 获取companyId
	user := controller.getUserInfo(c)
	if user.ID == 0 {
		log.Fatal("获取用户信息失败", user)
		controller.ResponseData(c, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}
	err = controller.srv.CreateCompanyMember(user.CompanyID, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, err.Error())
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}

//@Summary 增加钱包权限
// @Produce json
// @Param	body	body	request.AddWalletPermissionReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
//@Failure 500  "内部错误"
// @Router /api/v1/member/wallet/auth [post]
func (controller *CompanyMemberController) AddWalletPermission(c *gin.Context) {
	var req request.AddWalletPermissionReq
	if !controller.IsAdminS(c) {
		controller.ResponseData(c, nil, http.Forbidden)
		return
	}
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	// 获取companyId
	user := controller.getUserInfo(c)
	if user.ID == 0 {
		log.Fatal("获取用户信息失败", user)
		controller.ResponseData(c, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}
	//判断是否在相同组织下
	member := controller.srv.GetCompanyMember(req.MemberId)
	if member.CompanyID != user.CompanyID {
		controller.ResponseData(c, nil, http.BadRequest, "该用户不在操作权限内")
		return
	}
	err = controller.srv.DelWalletPermission(c, req.MemberId)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, "参数错误")
		return
	}
	err = controller.srv.AddWalletPermission(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, "钱包不存在！")
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}

//@Summary 获取钱包权限
// @Produce json
// @Param	body	body	request.GetWalletPermissionReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
//@Failure 500  "内部错误"
// @Router /api/v1/member/wallet [post]
func (controller *CompanyMemberController) MemberWallet(c *gin.Context) {
	var req request.GetWalletPermissionReq
	if !controller.IsAdminS(c) {
		controller.ResponseData(c, nil, http.Forbidden)
		return
	}
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	// 获取companyId
	user := controller.getUserInfo(c)
	if user.ID == 0 {
		log.Fatal("获取用户信息失败", user)
		controller.ResponseData(c, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}
	//判断是否在相同组织下
	member := controller.srv.GetCompanyMember(req.MemberId)
	if member.CompanyID != user.CompanyID {
		controller.ResponseData(c, nil, http.BadRequest, "该用户不在操作权限内")
		return
	}
	data, err := controller.srv.GetMemberWalletPermission(c, req.MemberId)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, "获取钱包权限")
		return
	}
	controller.ResponseData(c, data, http.Ok)
	return
}

//@Summary 成员禁用启用
// @Produce json
// @Param	body	body	request.EnableMemberReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
//@Failure 500  "内部错误"
// @Router /api/v1/member/status [put]
func (controller *CompanyMemberController) EnableMember(c *gin.Context) {
	var req request.EnableMemberReq
	if !controller.IsAdminS(c) {
		controller.ResponseData(c, nil, http.Forbidden)
		return
	}
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	// 获取companyId
	user := controller.getUserInfo(c)
	if user.ID == 0 {
		log.Fatal("获取用户信息失败", user)
		controller.ResponseData(c, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}
	// 判断是否在相同组织下
	member := controller.srv.GetCompanyMember(req.MemberId)
	if member.CompanyID != user.CompanyID {
		controller.ResponseData(c, nil, http.BadRequest, "该用户不在操作权限内")
		return
	}
	// 删除用户redis
	cache.DelUserInIdRedis(req.MemberId)
	err = controller.srv.EnableMember(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}

//@Summary 成员修改
// @Produce json
// @Param	body	body	request.UpdateCompanyMemberReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
//@Failure 500  "内部错误"
// @Router /api/v1/member/update [put]
func (controller *CompanyMemberController) Update(c *gin.Context) {
	var req request.UpdateCompanyMemberReq
	if !controller.IsAdminS(c) {
		controller.ResponseData(c, nil, http.Forbidden)
		return
	}
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	// 获取companyId
	user := controller.getUserInfo(c)
	if user.ID == 0 {
		log.Fatal("获取用户信息失败", user)
		controller.ResponseData(c, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}
	//判断是否在相同组织下
	member := controller.srv.GetCompanyMember(req.MemberId)
	if member.CompanyID != user.CompanyID {
		controller.ResponseData(c, nil, http.BadRequest, "该用户不在操作权限内")
		return
	}
	err = controller.srv.UpdateConpanyMember(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, "邮箱已存在！")
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}

//@Summary 成员密码修改
// @Produce json
// @Param	body	   body	 request.ChangePassword	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
//@Failure  500  "内部错误"
// @Router /api/v1/member/change/password [put]
func (controller *CompanyMemberController) ChangePassword(c *gin.Context) {
	if !controller.IsAdminS(c) {
		controller.ResponseData(c, nil, http.Forbidden)
		return
	}
	var req request.ChangePassword
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	err = controller.srv.ChangePassword(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, "新密码与旧密码不能一样！")
		// 获取companyId
		user := controller.getUserInfo(c)
		if user.ID == 0 {
			log.Fatal("获取用户信息失败", user)
			controller.ResponseData(c, nil, http.DataFieldInvalid, "获取用户信息失败")
			return
		}
		//判断是否在相同组织下
		member := controller.srv.GetCompanyMember(req.MemberId)
		if member.CompanyID != user.CompanyID {
			controller.ResponseData(c, nil, http.BadRequest, "该用户不在操作权限内")
			return
		}
		err = controller.srv.ChangePassword(c, req)
		if err != nil {
			controller.ResponseData(c, nil, http.BadRequest)
			return
		}

	}
	controller.ResponseData(c, nil, http.Ok)
	return
}
