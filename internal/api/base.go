package api

import (
	"altair-backend/internal/model"
	"altair-backend/internal/service"
	"altair-backend/log"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/validate"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"regexp"
	"strings"
)

type Status struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type base struct {
}

type JSONResponse struct {
	http.Status
	Data interface{} `json:"data"`
}

func (b *base) ResponseData(c *gin.Context, data interface{}, status http.Status, message ...string) {
	if len(message) > 0 {
		status.Msg = strings.Join(message[:], "; ")
	}

	jsonData := JSONResponse{
		status,
		data,
	}

	c.JSON(200, jsonData)
}

////统一错误处理方法
//func (b *base) render(c *gin.Context, data interface{}, err error) {
//	Code := getErrorCode(err)
//	response := model.Response{
//		Code: Code.Code,
//		Msg:  Code.Msg,
//		Data: getData(Code, data),
//	}
//	c.JSON(http2.StatusOK, response)
//}

////判断是否是组织管理员
//func (b *base) IsAdmin(c *gin.Context) {
//	var user model.CompanyMember
//	v, ok := c.Get("UserInfo")
//	if !ok {
//		b.render(c, nil, errcode.PermissionError)
//		return
//	}
//	user, ok = v.(model.CompanyMember)
//	if !ok {
//		b.render(c, nil, errcode.PermissionError)
//		return
//	}
//	if user.IsAdmin != true {
//		b.render(c, nil, errcode.PermissionError)
//		return
//	}
//	c.Next()
//	return
//}

func (b *base) getUserInfo(c *gin.Context) model.CompanyMember {
	var user model.CompanyMember
	v, ok := c.Get("UserInfo")
	if ok {
		user, ok = v.(model.CompanyMember)
		if ok {
			return user
		}
	}
	return user
}

//用户的钱包权限列表
func (b *base) getUserAuthIds(c *gin.Context) (wids []int, err bool) {
	userInput := b.getUserInfo(c)
	wids = service.GetWalletAuthorizeList(userInput)
	return wids, true
}

//func getErrorCode(err error) *errcode.Error {
//	if err == nil {
//		return errcode.Success
//	}
//	errCode, ok := err.(*errcode.Error)
//	if ok {
//		return errCode
//	}
//	// 包装一个普通的错误，默认500
//	return errcode.NewError(500, err.Error())
//}
//
//// 如果存在某些错误想返回数据，可以在这里控制
//func getData(err error, data interface{}) interface{} {
//	if err == errcode.Success {
//		return data
//	}
//	return nil
//}

// IsAdminS 是否是管理员
func (b *base) IsAdminS(c *gin.Context) bool {
	var user model.CompanyMember
	v, ok := c.Get("UserInfo")
	if !ok {
		return false
	}
	user, ok = v.(model.CompanyMember)
	if !ok {
		return false
	}
	return user.IsAdmin == true
}

// VerifyEmailFormat 检测邮箱
func (b *base) VerifyEmailFormat(Email string) bool {
	pattern := `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*` //匹配电子邮箱
	reg := regexp.MustCompile(pattern)
	return reg.MatchString(Email)
}

func (b *base) CommonValidateReturn(g *gin.Context, err error) {
	// 获取validator.ValidationErrors类型的errors
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		// 非validator.ValidationErrors类型错误直接返回
		log.Fatal("参数解析错误：" + err.Error())
		b.ResponseData(g, nil, http.DataFieldInvalid, err.Error())
		return
	}
	// validator.ValidationErrors类型错误则进行翻译
	// 并使用removeTopStruct函数去除字段名中的结构体名称标识
	res := validate.RemoveTopStruct(errs.Translate(validate.Trans))

	b.ResponseData(g, nil, http.DataFieldInvalid, res)
	return
}
