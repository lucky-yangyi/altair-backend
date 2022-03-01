package middleware

import (
	"altair-backend/config"
	"altair-backend/internal/cache"
	"altair-backend/internal/model"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strings"
)

// AdminMiddleWare 中间件
func AdminMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		AdminFilterSignature(c)
	}
}

func AdminValidateSignature(ctx *gin.Context) (bool, string, string, int) {
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
	cacheData, err := cache.AdminInfoGet(accessKey)

	if err != nil {
		return false, "登陆已过期", "", 401
	}

	var adminInfo cache.AdminLogin
	err = json.Unmarshal(cacheData, &adminInfo)
	if err != nil {
		return false, "用户数据错误", "", 403
	}

	if config.ServerConfig.RunMode == "debug" {
		AdminSetCtxData(ctx, adminInfo.Admin)
		return true, "", "", 200
	}

	serverSign := GenerateSignature(
		ctx.Request.RequestURI,
		ctx.Request.Method,
		ctx.GetHeader("X-Authorization-Date"),
		adminInfo.Token.SecretKey,
	)

	fmt.Println("clientSign======>", clientSign)
	fmt.Println("serverSign======>", serverSign)

	if clientSign == serverSign {
		// 设置用户上下文数据
		AdminSetCtxData(ctx, adminInfo.Admin)
		return true, "", "", 200
	} else {
		return false, "签名不一致", serverSign, 403
	}
}

func AdminFilterSignature(ctx *gin.Context) bool {

	ok, message, data, code := AdminValidateSignature(ctx)
	if !ok {
		Unauthorized(ctx, message+data, code)
		ctx.Abort()
		return false
	}
	ctx.Next()
	return true
}

//
//// GenerateSignature token校验
//func GenerateSignature(Url string, Method string, Time string, AccessKey string, SecretKey string) (sign string) {
//	message := fmt.Sprintf("%s\n%s\n%s", Time, Method, Url)
//
//	key := utils.GenerateMd5(SecretKey)
//	sign = utils.Hmac256AndBase256(message, key)
//
//	fmt.Println("key ===>", key)
//	fmt.Println("message ===>", message)
//	fmt.Println("sign ===>", sign)
//
//	return
//}

//func Hmac256AndBase256(message string, secret string) string {
//	key := []byte(secret)
//	h := hmac.New(sha256.New, key)
//	h.Write([]byte(message))
//
//	sha := hex.EncodeToString(h.Sum(nil))
//
//	log.Println("sha256 ----->", sha)
//
//	return base64.StdEncoding.EncodeToString([]byte(sha))
//}
//
//func GenerateMd5(input string) (output string) {
//	hash := md5.New()
//	hash.Write([]byte(input))
//	output = fmt.Sprintf("%x", hash.Sum(nil))
//	return
//}

func AdminSetCtxData(c *gin.Context, adminInput model.AdminNoPassword) {
	c.Set("AdminInfo", adminInput)
	//a, b := c.Get("UserInfo")
	//fmt.Println("UserInfo:", a, b)
}
