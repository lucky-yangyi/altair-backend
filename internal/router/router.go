package router

import (
	"altair-backend/internal/api"
	"altair-backend/internal/middleware"
	"github.com/gin-gonic/gin"

	_ "altair-backend/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	r.Use(middleware.Cors())
	//swagger 路由注册
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	v := r.Group("/api/v1")
	addAuthRouter(v)
	addAdminRouter(v)

	//运营后台
	vadmin := r.Group("/api/v1")
	vadmin.Use(middleware.AdminMiddleWare())
	addAdminBillRouter(vadmin)
	addCompanyRouter(vadmin)
	//数字资产
	v1 := r.Group("/api/v1")
	v1.Use(middleware.WebMiddleWare())
	addTransactionRouter(v1)
	addChainRouter(v1)
	addWalletRouter(v1)
	addCoinRouter(v1)
	addCompanyMemberRouter(v1)
	addMonthBillRouter(v1)
	AddWalletTypeRouter(v1)
	addHomeRouter(v1)

	//api鉴权
	vapi := r.Group("/api/v1")
	vapi.Use(middleware.ApiMiddleWare())
	addApiRouter(vapi)

	return r
}

//币种管理
func addCoinRouter(v1 *gin.RouterGroup) {
	coinController := api.GetCoinController()
	v1.POST("/coin/list", coinController.List)
}

//币类别管理
func AddWalletTypeRouter(v1 *gin.RouterGroup) {
	wallettypeController := api.GetWalletTypeController()
	v1.POST("/wallet/type/list", wallettypeController.List)
}

//成员管理
func addCompanyMemberRouter(v1 *gin.RouterGroup) {
	companymemberController := api.GetCompanyMemberController()
	v1.POST("/member/list", companymemberController.List)
	// 是否超级管理员
	//v1.POST("/member/create", companymemberController.IsAdmin, companymemberController.Create)
	v1.POST("/member/create", companymemberController.Create)
	v1.POST("/member/wallet/auth", companymemberController.AddWalletPermission)
	v1.POST("/member/wallet", companymemberController.MemberWallet)
	v1.PUT("/member/status", companymemberController.EnableMember)
	v1.PUT("/member/change/password", companymemberController.ChangePassword)
	v1.PUT("/member/update", companymemberController.Update)
}

//企业组织管理
func addCompanyRouter(v1 *gin.RouterGroup) {
	companyController := api.GetCompanyController()
	v1.POST("/admin/company/list", companyController.List)
	v1.POST("/admin/company/list/member", companyController.Member)
	v1.POST("/admin/company/create", companyController.Create)
	v1.POST("/admin/company/member", companyController.CreateMember)
	v1.POST("/admin/company/add", companyController.Add)
	v1.POST("/admin/company/update", companyController.Update)
	v1.PUT("/admin/company/status", companyController.EnableCompany)
	v1.PUT("/admin/company/state", companyController.AdminMember)
}

//资产月付账单
func addMonthBillRouter(v1 *gin.RouterGroup) {
	monthbillController := api.GetMonthBillController()
	v1.POST("/bill/month/list", monthbillController.List)
}

func addHomeRouter(v1 *gin.RouterGroup) {
	homeController := api.HomeController{}
	v1.POST("home/stat", homeController.Stat)
}

//运营月付账单
func addAdminBillRouter(v1 *gin.RouterGroup) {
	billController := api.GetBillController()
	v1.POST("/bill/list", billController.List)
	v1.PUT("/bill/status", billController.StatusBill)

}

//上链操作
func addChainRouter(v1 *gin.RouterGroup) {
	chainController := api.GetChainController()
	v1.GET("wallet/balance", chainController.Balance)
	v1.POST("chain/receipt", chainController.Receipt)
}

func addWalletRouter(v1 *gin.RouterGroup) {
	walletController := api.GetWalletController()
	v1.POST("wallet/create", walletController.ChainCreate)
	v1.POST("wallet/list", walletController.List)
	v1.POST("wallet/list/app", walletController.AppList)
	v1.GET("wallet/list/no/page", walletController.WalletListNoPage) //WalletListNoPage
	v1.PUT("wallet/update", walletController.Update)
	v1.POST("wallet/add", walletController.Add)
	v1.POST("wallet/count", walletController.Count)
	v1.POST("wallet/collect/key", walletController.CollectKey)
	v1.POST("wallet/collect/enable", walletController.EnableCollect)
	v1.POST("wallet/auth/ids", walletController.WalletAuthIds) //wallet/auth/ids
}

func addApiRouter(v1 *gin.RouterGroup) {
	coll := api.CollectController{}
	v1.POST("collect/new/wallet", coll.CollectNew)
	v1.POST("collect/transaction", coll.TransactionOut)
	v1.GET("collect/get/address", coll.GetSonAddresses) //collect/get/address  GetSonAddresses
}

//资产端登陆
func addAuthRouter(v1 *gin.RouterGroup) {
	auth := api.AuthController{}
	v1.POST("user/login", auth.Login)
}

//运营登陆
func addAdminRouter(v1 *gin.RouterGroup) {
	adminController := api.AdminController{}
	v1.POST("admin/login", adminController.Login)
}

//交易管理
func addTransactionRouter(v1 *gin.RouterGroup) {
	trans := api.GetTransactionController()
	v1.POST("transaction/add", trans.Add)
	v1.PUT("transaction/signature", trans.UpdateTransactionDetail)
	v1.POST("transaction/all", trans.GetAll) //
}
