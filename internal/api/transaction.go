package api

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/log"
	"altair-backend/pkg/address"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm/clause"
	"strings"
)

var GloTransactionController *TransactionController

type TransactionController struct {
	base
	srv *service.TransactionService
}

func newTransactionController(srv *service.TransactionService) *TransactionController {
	return &TransactionController{
		srv: srv,
	}
}

//初始化
func GetTransactionController() *TransactionController {
	GloTransactionController = newTransactionController(service.GetTransactionService())
	return GloTransactionController
}

//@Summary 添加交易
//@Produce json
// @Param	body	body	request.AddTransaction true "body参数"
//@Success 200   "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/transaction/add [post]
func (controller *TransactionController) Add(g *gin.Context) {
	var param request.AddTransaction
	err := g.BindJSON(&param)
	if err != nil {
		controller.CommonValidateReturn(g, err)
		return
	}

	//验证 MultiSignWalletID 是否存在
	var wallet model.Wallet
	if err := dao.DB.First(&wallet, param.WalletID).Error; err != nil {
		log.Error(err.Error())
		controller.ResponseData(g, nil, http.DataNotFound, "多签 ID 不存在")
		return
	}

	// 判断钱包是否创建成功，不成功不允许产生交易
	if wallet.Status != dao.MultiSignWalletStatusSuccess {
		controller.ResponseData(g, nil, http.BadRequest, "钱包为不可用状态，不允许交易")
		return
	}

	//接收地址验证
	_, err = address.NewFromString(param.DestinationWallet)
	if err != nil {
		log.Error("地址：" + param.DestinationWallet + "，错误:" + err.Error())
		controller.ResponseData(g, nil, http.BadRequest, "钱包地址错误")
		return
	}
	// 检查金额
	if param.Amount <= 0 {
		controller.ResponseData(g, nil, http.BadRequest, "转账金额错误")
		return
	}
	if param.Amount > wallet.Balance {
		controller.ResponseData(g, nil, http.BadRequest, "转账金额超出余额")
		return
	}

	if amounts := strings.Split(fmt.Sprintf("%v", param.Amount), "."); len(amounts) == 2 && len(amounts[1]) > 5 {
		controller.ResponseData(g, nil, http.BadRequest, "转账金额最多支持 5 位小数")
		return //c.ResponseData(nil, http.BadRequest, "转账金额最多支持 5 位小数")
	}

	//判断是否有钱包权限
	ids, _ := controller.getUserAuthIds(g)
	isAuth := utils.IsInArray(int(param.WalletID), ids)
	if !isAuth {
		controller.ResponseData(g, nil, http.Forbidden, "没有权限操作该钱包")
		return
	}

	user := service.GetUserInfo(g)

	tn := model.Transaction{
		FromWalletID:      param.WalletID,
		Type:              dao.TransactionTypeOut,
		SubType:           dao.TransactionSubTypeDefault,
		Comment:           param.Comment,
		CompanyMemberID:   user.ID,
		Amount:            param.Amount,
		DestinationWallet: param.DestinationWallet,
		CoinID:            dao.FIL, //todo::
		Status:            dao.TransactionStatusDoing,
	}

	id, err := dao.GlobalDao.AddTransaction(tn, param.WalletID)

	if err == nil {
		controller.ResponseData(g, map[string]uint64{"ID": id}, http.Ok)
		return
	} else {
		log.Error("转出地址id：", param.WalletID, "", "执行添加交易失败："+err.Error())
		controller.ResponseData(g, nil, http.BadRequest)
		return //c.ResponseData(nil, http.BadRequest)
	}

}

//@Summary 交易回执，修改交易详情
//@Produce json
// @Param	body	body	request.UpdateTransactionDetail true "交易签名body参数"
//@Success 200   "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/transaction/signature [put]
func (controller *TransactionController) UpdateTransactionDetail(g *gin.Context) {
	var param request.UpdateTransactionDetail
	err := g.BindJSON(&param)
	if err != nil {
		controller.CommonValidateReturn(g, err)
		return
	}
	// 检查交易
	var trans model.Transaction
	if err := dao.DB.Preload(clause.Associations).First(&trans, param.TransactionID).Error; err != nil {
		//todo::打印日志
		log.Error("修改交易详情id:", param.TransactionID, "。error:"+err.Error())
		controller.ResponseData(g, nil, http.BadRequest, "交易 ID 错误")
		return
	}

	//判断是否有钱包权限
	ids, _ := controller.getUserAuthIds(g)
	isAuth := utils.IsInArray(int(param.WalletID), ids)
	if !isAuth {
		controller.ResponseData(g, nil, http.Forbidden, "没有权限操作该钱包")
		return
	}

	res, message := service.UpdateTransactionDetail(param)
	if res {
		controller.ResponseData(g, nil, http.Ok)
		return
	}
	controller.ResponseData(g, nil, http.BadRequest, message)
	return
}

//@Summary 交易列表
//@Produce json
// @Param	body	body	request.GetAllTransactionRequest true "交易签名body参数"
//@Success 200   {object} model.Tmix "成功"
//@Failure 400 "请求错误"
//@Failure 500 "内部错误"
//@Router /api/v1/transaction/all [post]
func (controller *TransactionController) GetAll(g *gin.Context) {
	var Param request.GetAllTransactionRequest

	err := g.BindJSON(&Param)
	if err != nil {
		controller.CommonValidateReturn(g, err)
		return
	}

	var list []model.Tmix
	var trans model.Transaction
	query := dao.DB.Table(trans.TableName()).Order("id DESC")

	// 筛选授权钱包的交易
	ids, isOk := controller.getUserAuthIds(g)
	if !isOk {
		log.Fatal("获取ids失败", ids, isOk)
		controller.ResponseData(g, nil, http.DataFieldInvalid, "获取钱包id失败")
		return
	}

	query.Where("from_wallet_id IN ?", ids)

	// 搜索钱包 ID
	if Param.FromWalletID != 0 {
		query = query.Where("from_wallet_id = ?", Param.FromWalletID)
	}

	if Param.ChainTransactionId != "" {
		query = query.Where("chain_transaction_id = ?", Param.ChainTransactionId)
	}

	if Param.DestinationWallet != "" {
		query = query.Where("destination_wallet = ?", Param.DestinationWallet)
	}
	if Param.SerialNumber != "" {
		query = query.Where("serial_number = ?", Param.SerialNumber)
	}

	if Param.Status == int8(dao.TransactionStatusDoing) {
		query = query.Where("status = 1")
	} else {
		query = query.Where("status <> 1")
	}

	if Param.Type != 0 && Param.Type != -1 {
		query = query.Where("type = ?", Param.Type)
	} else if Param.Type == -1 {
		query = query.Where("type <> ?", 2)
	}

	if Param.ID != 0 {
		query = query.Where("id = ?", Param.ID)
	}

	if len(Param.Daterange) == 2 {
		query = query.Where("created_at > ?", utils.StringToTimestamp(Param.Daterange[0])).Where("created_at < ?", utils.StringToTimestamp(Param.Daterange[1]))
	}

	page, query, err := utils.Paginate(query, Param.PageNo, Param.PageSize, &list)

	query.Preload(clause.Associations).Preload("TransactionDetail.Wallet").Find(&list)

	if err != nil {
		log.Fatal(err.Error())
		controller.ResponseData(g, nil, http.BadRequest)
		return
	}
	controller.ResponseData(g, page, http.Ok)

}
