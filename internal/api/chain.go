package api

import (
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/pkg/constants/http"
	"github.com/gin-gonic/gin"
)

type ChainController struct {
	base
}

func newChainController() *ChainController {
	return &ChainController{}
}

//初始化
func GetChainController() *ChainController {
	GloChainController := newChainController()
	return GloChainController
}

//@Summary 获取钱包余额
//@Produce json
// @Param	body	body	request.Balance	true	"email 邮箱; password 密码"
//@Success 200 {object} response.Balance "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/balance [post]
func (c *ChainController) Balance(g *gin.Context) {
	var param request.Balance
	err := g.BindJSON(&param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}
	data, err := service.GetBalance(param)
	if err != nil {
		c.ResponseData(g, nil, http.BadRequest, "获取失败，请核实钱包地址")
		return
	}
	c.ResponseData(g, data.Data, http.Ok)
}

//@Summary 交易回执
//@Produce json
//@Success 200 {object} response.Balance "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/chain/receipt [post]
func (c *ChainController) Receipt(g *gin.Context) {
	var param request.Receipt
	err := g.BindJSON(&param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}
	data, err := service.Receipt(param)
	if err != nil {
		c.ResponseData(g, nil, http.BadRequest, "获取失败，请核实txId参数")
		return
	}
	c.ResponseData(g, data.Data, http.Ok)
}
