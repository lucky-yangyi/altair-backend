package api

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/internal/service"
	"altair-backend/log"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
)

type CollectController struct {
	base
}

//初始化
func GetCollectController() *CollectController {
	Controller := &CollectController{}
	return Controller
}

//@Summary api：获得该钱包所有子地址
//@Produce json
//@Success 200 {object} response.CollectAddress "成功"
//@Failure 400   "请求错误"
//@Failure 500   "内部错误"
//@Router /api/v1/collect/get/address [post]
func (c *CollectController) GetSonAddresses(g *gin.Context) {
	var address []string
	id, exists := g.Get("wid")
	if exists {
		dao.DB.Table(new(model.WalletCollect).TableName()).Where("wallet_id = ?", id).Pluck("address", &address)
	}
	c.ResponseData(g, address, http.Ok)
}

//@Summary api：新建子（归集）钱包
//@Produce json
//@Param	body	body	request.AddWalletCollect true "body参数"
//@Success 200 {object} response.CollectNew "成功"
//@Failure 400   "请求错误"
//@Failure 500   "内部错误"
//@Router /api/v1/collect/new/wallet [post]
func (c *CollectController) CollectNew(g *gin.Context) {
	var param request.AddWalletCollect
	err := g.BindJSON(&param)
	if err != nil {
		log.Fatal(err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, "数据解析错误", err.Error())
		return
	}

	id, exists := g.Get("wid")
	if exists {
		var add response.CollectNew
		add.Address, err = service.WalletCollectAdd(id.(uint64), param.Symbol)
		if err != nil {
			log.Fatal(err.Error())
			c.ResponseData(g, nil, http.DataFieldInvalid, "数据解析错误", err.Error())
			return
		}
		c.ResponseData(g, add, http.Ok)
		return
	}
	c.ResponseData(g, nil, http.DataFieldInvalid, "AccessKey和Sign有误")
	return
}

//@Summary api：通过归集转账
//@Produce json
//@Param	body	body	request.CollectTransaction true "body参数"
//@Success 200  "成功"
//@Failure 400   "请求错误"
//@Failure 500   "内部错误"
//@Router /api/v1/collect/transaction [post]
func (c *CollectController) TransactionOut(g *gin.Context) {
	var param request.CollectTransaction
	err := g.BindJSON(&param)
	if err != nil {
		log.Fatal(err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, "数据解析错误", err.Error())
		return
	}
	id, exists := g.Get("wid")
	if !exists {
		c.ResponseData(g, nil, http.DataFieldInvalid, "获取钱包id有误，请重试")
		return
	}

	//判断余额是否足够
	var wallet model.Wallet
	dao.DB.Table(wallet.TableName()).Where("id = ?", id).First(&wallet)
	if wallet.Balance <= param.Amount {
		c.ResponseData(g, nil, http.DataFieldInvalid, "该钱包余额不足")
		return
	}

	// 创建普通交易
	tn := model.Transaction{
		FromWalletID:      id.(uint64),
		Type:              dao.TransactionTypeOut,
		SubType:           dao.TransactionSubTypeDefault,
		Comment:           "用户提币，需要从归集地址转账: " + param.Comment,
		CompanyMemberID:   9998,
		Amount:            param.Amount,
		DestinationWallet: param.Address,
	}

	ids, err := dao.GlobalDao.AddTransaction(tn, id.(uint64))

	if err == nil {
		c.ResponseData(g, map[string]uint64{"ID": ids}, http.Ok)
		return
	} else {
		log.Fatal("执行添加交易失败：" + err.Error())
		c.ResponseData(g, nil, http.BadRequest)
		return
	}

}

//@Summary api：归集签名测试
//@Produce json
//@Success 200  "成功"
//@Failure 400   "请求错误"
//@Failure 500   "内部错误"
//@Router /api/v1/collect/test [get]
func (c *CollectController) TestSign(g *gin.Context) {
	var param request.CollectTest
	err := g.BindJSON(&param)
	if err != nil {
		log.Fatal(err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, "数据解析错误", err.Error())
		return
	}
	var res response.CollectTest

	res.Message = fmt.Sprintf("%s\n%s\n%s", g.GetHeader("X-Authorization-Date"), g.Request.Method, g.Request.RequestURI)
	res.Key = utils.GenerateMd5(param.SecretKey)
	res.Sign = utils.Hmac256AndBase256(res.Message, res.Sign)
	c.ResponseData(g, res, http.Ok)
	return
}
