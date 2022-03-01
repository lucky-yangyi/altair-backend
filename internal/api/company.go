package api

import (
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var GloCompanyController *CompanyController

type CompanyController struct {
	base
	srv *service.CompanyService
}

func newCompanyController(srv *service.CompanyService) *CompanyController {
	return &CompanyController{
		srv: srv,
	}
}

//初始化
func GetCompanyController() *CompanyController {
	GloCompanyController = newCompanyController(service.GetCompanyService())
	return GloCompanyController
}

// @Summary 获取企业列表
// @Produce json
// @Param	body	body	request.CompanyListReq	true	"body参数"
// @Success 200 {object} response.CompanyListResp "成功"
// @Failed 400
// @Failure 500  "内部错误"
// @Router /api/v1/admin/company/list [post]
func (controller *CompanyController) List(c *gin.Context) {
	var req request.CompanyListReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	page, size := utils.CheckIsPage(req.PageNo, req.PageSize)
	req.PageSize = size
	req.PageNo = page
	data, err := controller.srv.GetCompanyList(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, data, http.Ok)
	return

}

// @Summary 新增企业
// @Produce json
// @Param	body	body	request.CreateCompanyReq	true	"body参数"
// @Success 200 "成功"
// @Failed 400
// @Router /api/v1/admin/company/create [post]
func (controller *CompanyController) Create(c *gin.Context) {
	var req request.CreateCompanyReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	company, err := controller.srv.CreateCompany(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, company, http.Ok)
	return
}

// @Summary 修改企业
// @Produce json
// @Param	body	body	request.UpdateCompanyReq	true	"body参数"
// @Success 200 "成功"
// @Failed 400
// @Failure 500  "内部错误"
// @Router /api/v1/admin/company/update [post]
func (controller *CompanyController) Update(c *gin.Context) {
	var req request.UpdateCompanyReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	err = controller.srv.UpdateConpany(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}

//@Summary 企业禁用
// @Produce json
// @Param	body	body	request.EnableCompanyReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
//@Failure  500  "服务器错误"
// @Router /api/v1/admin/company/status [put]
func (controller *CompanyController) EnableCompany(c *gin.Context) {
	var req request.EnableCompanyReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	err = controller.srv.EnableCompany(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}

// @Summary 增加成员
// @Produce json
// @Param	body	body	request.AddCompanyMemberReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
// @Failure 403 "没有权限访问"
//@Failure 500  "内部错误"
// @Router /api/v1/admin/company/add [post]
func (controller *CompanyController) Add(c *gin.Context) {
	var req request.AddCompanyMemberReq
	//if !controller.IsAdminS(c) {
	//	controller.ResponseData(c, nil, http.Forbidden)
	//	return
	//}
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	err = controller.srv.AddCompanyMember(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, "邮箱已存在！")
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}

// @Summary 获取企业成员列表
// @Produce json
// @Param	body	body	request.MemberListReq	true	"body参数"
// @Success 200 {object} response.MemberListResp "成功"
//@Failure 500  "内部错误"
// @Router /api/v1/admin/company/list/member [post]
func (controller *CompanyController) Member(c *gin.Context) {
	var req request.MemberListReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	page, size := utils.CheckIsPage(req.PageNo, req.PageSize)
	req.PageSize = size
	req.PageNo = page
	data, err := controller.srv.GetMemberList(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, data, http.Ok)
	return
}

//@Summary 设置管理员
// @Produce json
// @Param	body	   body	request.AdminMemberReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
//@Failure 500  "内部错误"
// @Router /api/v1/admin/company/state [put]
func (controller *CompanyController) AdminMember(c *gin.Context) {
	var req request.AdminMemberReq
	//if !controller.IsAdminS(c) {
	//	controller.ResponseData(c, nil, http.Forbidden)
	//	return
	//}
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	err = controller.srv.AdminMember(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}

// @Summary 增加企业成员
// @Produce json
// @Param	body	body	request.CreateMemberReq	true	"body参数"
// @Success 200 "成功"
// @Failure 400 "入参错误"
//@Failure 500  "内部错误"
// @Router /api/v1/admin/company/member [post]
func (controller *CompanyController) CreateMember(c *gin.Context) {
	var req request.CreateMemberReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	re := utils.VerifyPasswordRule(req.MemberPwd)
	if !re {
		controller.ResponseData(c, nil, http.BadRequest, "密码需要大小写字母+数字组合")
		return
	}
	err = controller.srv.CreateMember(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, "邮箱已存在！")
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}
