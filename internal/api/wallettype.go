package api

import (
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var GloWalletTypeController *WalletTypeController

type WalletTypeController struct {
	base
	srv *service.WalletTypeService
}

func newWalletTypeController(srv *service.WalletTypeService) *WalletTypeController {
	return &WalletTypeController{
		srv: srv,
	}
}

//GetWalletTypeController 初始化
func GetWalletTypeController() *WalletTypeController {
	GloWalletTypeController = newWalletTypeController(service.GetWalletTypeService())
	return GloWalletTypeController
}

// @Summary 获取币类别列表
// @Produce json
// @Success 200 {object} response.WalletTypeListResp "成功"
// @Failure 400 "入参错误"
// @Failure 500  "内部错误"
// @Router /api/v1/wallet/type/list [post]
func (controller *WalletTypeController) List(c *gin.Context) {
	var req request.WalletTypeListReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.ResponseData(c, nil, http.DataFieldInvalid, "数据解析错误")
		return
	}
	page, size := utils.CheckIsPage(req.PageNo, req.PageSize)

	req.PageNo = page
	req.PageSize = size
	data, err := controller.srv.GetWalletTypeList(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, data, http.Ok)
	return
}
