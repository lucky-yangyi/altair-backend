package api

import (
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"github.com/gin-gonic/gin"
)

var GloCoinController *CoinController

type CoinController struct {
	base
	srv *service.CoinService
}

func newCoinController(srv *service.CoinService) *CoinController {
	return &CoinController{
		srv: srv,
	}
}

//初始化
func GetCoinController() *CoinController {
	GloCoinController = newCoinController(service.GetCoinService())
	return GloCoinController
}

// @Summary 获取币种列表
// @Produce json
// @Param	body	body	request.CoinListReq	true	"body参数"
// @Success 200 {object} response.CoinListResp "成功"
// @Failure 400 "入参错误"
//@Failure 500  "内部错误"
// @Router /api/v1/coin/list [post]
func (controller *CoinController) List(c *gin.Context) {
	var req request.CoinListReq
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	page, size := utils.CheckIsPage(req.PageNo, req.PageSize)
	req.PageSize = size
	req.PageNo = page
	data, err := controller.srv.GetCoinList(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, data, http.Ok)
	return
}

func (controller *CoinController) Create(c *gin.Context) {
	var req request.CoinCreate
	err := c.BindJSON(&req)
	if err != nil {
		controller.CommonValidateReturn(c, err)
		return
	}
	err = controller.srv.CreateCoin(c, req)
	if err != nil {
		controller.ResponseData(c, nil, http.BadRequest)
		return
	}
	controller.ResponseData(c, nil, http.Ok)
	return
}
