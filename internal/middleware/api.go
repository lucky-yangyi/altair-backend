package middleware

import (
	"altair-backend/config"
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"github.com/gin-gonic/gin"
	"strings"
)

// ApiMiddleWare 中间件
func ApiMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		FilterApiSign(c)
	}
}

func FilterApiSign(ctx *gin.Context) bool {
	ok, message, data, code := NewValidateApiSign(ctx)
	if !ok {
		Unauthorized(ctx, message+data, code)
		ctx.Abort()
		return false
	}
	ctx.Next()

	return true
}

func ValidateApiSign(ctx *gin.Context) (bool, string, string, int) {
	// 查看是否授权
	AccessKey := ctx.GetHeader("AccessKey")
	Sign := ctx.GetHeader("Sign")
	var wallet model.Wallet

	//fmt.Println(AccessKey,"sssssssss",Sign)

	dao.DB.Table(wallet.TableName()).Where("access_key = ?", AccessKey).Where("sign = ?", Sign).First(&wallet)
	if wallet.ID != 0 {
		ctx.Set("wid", wallet.ID)
		return true, "", "", 200
	}
	return false, "AccessKey和Sign有误,请核实", "", 403
}

func NewValidateApiSign(ctx *gin.Context) (bool, string, string, int) {
	// 查看是否授权
	auth := ctx.GetHeader("Authorization")
	if auth == "" {
		return false, "请求头缺少 Authorization", "", 403
	}

	auths := strings.Split(auth, " ")
	if len(auths) != 2 {
		return false, "Authorization 格式不正确", "", 403
	}

	credential := strings.Split(auths[1], ":")
	if len(credential) != 2 {
		return false, "签名格式不正确", "", 403
	}

	AccessKey := credential[0]
	Sign := credential[1]
	var wallet model.WalletPrivate
	dao.DB.Table(wallet.TableName()).Where("access_key = ?", AccessKey).First(&wallet)
	if wallet.ID == 0 {
		return false, "AccessKey有误,请核实", "", 403
	}

	serverSign := GenerateSignature(
		ctx.Request.RequestURI,
		ctx.Request.Method,
		ctx.GetHeader("X-Authorization-Date"),
		wallet.SecretKey,
	)
	if config.ServerConfig.RunMode == "debug" {
		ctx.Set("wid", wallet.ID)
		ctx.Header("apiSign", serverSign)
		return true, "", "", 200
	}
	if serverSign != Sign {
		return false, "Sign有误,请核实", "", 403
	}
	ctx.Set("wid", wallet.ID)
	return true, "", "", 200
}
