package api

import (
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var GloBillController *BillController

type BillController struct {
	base
	srv *service.BillService
}

func newBillController(srv *service.BillService) *BillController {
	return &BillController{
		srv: srv,
	}
}

//初始化
func GetBillController() *BillController {
	GloBillController = newBillController(service.GetBillService())
	return GloBillController
}

// @Summary 运营后台月付账单
// @Produce json
// @Param	body	body	request.BillListReq	true	"body参数"
// @Success 200 {object} response.MonthListResp "成功"
// @Router /api/v1/bill/list [post]
func (controller *BillController) List(c *gin.Context) {
	var req request.BillListReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	page, size := utils.CheckIsPage(req.PageNo, req.PageSize)
	req.PageSize = size
	req.PageNo = page
	data, err := controller.srv.GetBillList(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, err.Error())
		return
	}
	controller.ResponseData(c, data, http.Ok)
	return

}

//@Summary 账单缴费
// @Produce json
// @Param	body	body	request.StatusBillReq	true	"body参数"
// @Success 200 "成功"
// @Router /api/v1/bill/status [put]
func (controller *BillController) StatusBill(c *gin.Context) {
	var req request.StatusBillReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	err = controller.srv.StatusBill(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest, err.Error())
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}
