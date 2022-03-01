package middleware

import (
	"altair-backend/config"
	"altair-backend/internal/cache"
	"altair-backend/internal/model"
	"altair-backend/internal/service"
	"altair-backend/log"
	"altair-backend/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// WebMiddleWare 中间件
func WebMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		FilterSignature(c)
	}
}

func ValidateSignature(ctx *gin.Context) (bool, string, string, int) {
	//date := ctx.GetHeader("X-Authorization-Date")
	//if date == "" {
	//	return false, "请求头缺少 X-Authorization-Date", ""
	//}
	//
	//if !ValidateTime(date, 300) {
	//	return false, "签名已过期"
	//}

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

	accessKey := credential[0]
	clientSign := credential[1]

	// 从 Redis 缓存中获取用户数据
	cacheData, err := cache.UserInfoGet(accessKey)

	if err != nil {
		log.Fatal("获取accessKey:" + accessKey + ",err:" + err.Error())
		return false, "登陆已过期", "", 401
	}

	var userInfo cache.Login
	err = json.Unmarshal(cacheData, &userInfo)
	if err != nil {
		return false, "用户数据错误", "", 403
	}

	status, err := cache.CompanyInfoGet(userInfo.User.CompanyID)
	if err != nil {
		companys := service.GetCompany(userInfo.User.CompanyID)
		status = companys.Enabled != 0
		cache.SetCompanyInRedis(userInfo.User.CompanyID, status)
		//return false, "公司数据错误", "", 403
	}
	if !status {
		return false, "公司被封禁", "", 403
	}
	serverSign := GenerateSignature(
		ctx.Request.RequestURI,
		ctx.Request.Method,
		ctx.GetHeader("X-Authorization-Date"),
		userInfo.Token.SecretKey,
	)
	if config.ServerConfig.RunMode == "debug" {
		SetCtxData(ctx, userInfo.User)
		ctx.Header("ServerSign", serverSign)
		return true, "", "", 200
	}

	fmt.Println("clientSign======>", clientSign)
	fmt.Println("serverSign======>", serverSign)

	if clientSign == serverSign {
		// 设置用户上下文数据
		SetCtxData(ctx, userInfo.User)
		return true, "", "", 200
	} else {
		return false, "签名不一致", serverSign, 403
	}
}

func Unauthorized(ctx *gin.Context, v interface{}, code int) {
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"code": code,
		"data": nil,
		"msg":  v,
	})
}

func FilterSignature(ctx *gin.Context) bool {
	if strings.TrimRight(ctx.Request.RequestURI, "/") != "/api/v1/login" {
		ok, message, data, code := ValidateSignature(ctx)
		if !ok {
			Unauthorized(ctx, message+data, code)
			ctx.Abort()
			return false
		}
		ctx.Next()
	}
	return true
}

// GenerateSignature token校验
func GenerateSignature(Url string, Method string, Time string, SecretKey string) (sign string) {
	message := fmt.Sprintf("%s\n%s\n%s", Time, Method, Url)
	fmt.Println("校验-入参：", Time, Method, Url)
	key := utils.GenerateMd5(SecretKey)
	sign = utils.Hmac256AndBase256(message, key)

	fmt.Println("key ===>", key)
	fmt.Println("message ===>", message)
	fmt.Println("sign ===>", sign)

	return
}

func SetCtxData(c *gin.Context, userInput model.CompanyMember) {
	c.Set("UserInfo", userInput)
}
