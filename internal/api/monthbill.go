package api

import (
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var GloMonthBillController *MonthBillController

type MonthBillController struct {
	base
	srv *service.MonthBillService
}

func newMonthBillController(srv *service.MonthBillService) *MonthBillController {
	return &MonthBillController{
		srv: srv,
	}
}

//初始化
func GetMonthBillController() *MonthBillController {
	GloMonthBillController = newMonthBillController(service.GetMonthBillService())
	return GloMonthBillController
}

// @Summary 月付账单
// @Produce json
// @Param	body	body	request.MonthBillListReq	true	"body参数"
// @Success 200 {object} response.MonthListResp "成功"
// @Router /api/v1/bill/month/list [post]
func (controller *MonthBillController) List(c *gin.Context) {
	var req request.MonthBillListReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	//1.查企业 -钱包-账单
	page, size := utils.CheckIsPage(req.PageNo, req.PageSize)
	req.PageSize = size
	req.PageNo = page
	data, err := controller.srv.GetMonthBillList(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, err.Error())
		return
	}
	controller.ResponseData(c, data, http.Ok)
	return

}
