package api

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/internal/service"
	"altair-backend/pkg/constants/http"
	"github.com/gin-gonic/gin"
)

type HomeController struct {
	base
}

//@Summary 首页
//@Produce json
//@Success 200 {object} response.Stat "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/home/stat [post]
func (c *HomeController) Stat(g *gin.Context) {
	var resp response.Stat
	var msw []model.Wallet
	var ids []uint64 // 多签id集合

	query := dao.DB.Model(model.Wallet{})

	// 非管理员，则筛选授权钱包的交易
	wids, _ := c.getUserAuthIds(g)

	query.Where("id IN ?", wids)
	query.Find(&msw)

	// 1.TotalStat,ids
	ids, resp.TotalStat = service.TotalStat(msw)

	// 3.用户有权限的钱包列表
	resp.Msw, resp.Pipe = service.MultiSignWalletList(msw, resp.TotalStat.TotalAmount)

	// 2.6个月出入
	resp.SixMonthAmount = service.SixMonthAmountList(ids)

	//4 钱包近15天的收入支出金额 和次数
	resp.FifteenDayInOutNum = service.FifteenDayAmountInOutList(ids)

	//5 3个月每个钱包的3个月支出次数 ==> 钱包列表 中有
	resp.ThreeMonthNum = service.ThreeMonthNum(resp.Msw)

	c.ResponseData(g, resp, http.Ok)
}

//@Summary 获取钱包余额
//@Produce json
// @Param	body	body	request.Balance	true	"email 邮箱; password 密码"
//@Success 200 {object} response.Balance "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/balance [post]
func (c *HomeController) StatById(g *gin.Context) {
	var resp response.StatById
	var ids []uint64
	var param request.Ids

	err := g.BindJSON(&param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}

	//msw
	var msw []model.Wallet
	dao.DB.Model(model.Wallet{}).Where("id IN ?", param.Id).Find(&msw)

	// 1.TotalStat,ids
	ids, resp.TotalStat = service.TotalStat(msw)

	// 3.用户有权限的钱包列表
	resp.Msw, resp.Pipe = service.MultiSignWalletList(msw, resp.TotalStat.TotalAmount)

	// 4.钱包近15天的收入支出金额
	// 6.钱包近15天的收入支出次数
	resp.FifteenDayInOutNum = service.FifteenDayAmountInOutListById(ids)

	//5 3个月每个钱包的3个月支出次数
	//Msw[walletId].ThreeMonthOut/ThreeMonthIn
	resp.ThreeMonthNum = service.ThreeMonthNum(resp.Msw)

	c.ResponseData(g, resp, http.Ok)
}
