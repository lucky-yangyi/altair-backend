package api

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/internal/service"
	"altair-backend/log"
	"altair-backend/pkg/address"
	"altair-backend/pkg/constants/http"
	"altair-backend/pkg/utils"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"time"
)

type WalletController struct {
	base
}

//初始化
func GetWalletController() *WalletController {
	WalletController := &WalletController{}
	return WalletController
}

//@Summary 获取不分页钱包列表
//@Produce json
//@Success 200 {object} model.WalletAuth "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/list/no/page [get]
func (c *WalletController) WalletListNoPage(g *gin.Context) {
	userStruct := service.GetUserInfo(g)
	//fmt.Println(userStruct)
	if userStruct.ID == 0 {
		log.Fatal("获取用户信息失败", userStruct)
		c.ResponseData(g, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}

	data, errs := service.WalletListNoPage(userStruct)
	if errs != nil {
		log.Fatal("获取WalletList失败", errs.Error())
		c.ResponseData(g, nil, http.BadRequest, "获取WalletList失败")
		return
	}

	c.ResponseData(g, data, http.Ok)
	return
}

//@Summary 获取钱包列表
//@Produce json
// @Param	body	body	request.WalletList true "body参数"
//@Success 200 {object} response.Page "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/list [post]
func (c *WalletController) List(g *gin.Context) {
	var Param request.WalletList
	err := g.BindJSON(&Param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}

	c.getWalletList(Param, g)
	return
}

//@Summary app获取钱包列表
//@Produce json
// @Param	body	body	request.WalletList true "body参数"
//@Success 200 {object} model.WalletMix "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/list/app [post]
func (c *WalletController) AppList(g *gin.Context) {
	var Param request.WalletList
	err := g.BindJSON(&Param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}

	c.getWalletList(Param, g)
	return
}

func (c *WalletController) getWalletList(Param request.WalletList, g *gin.Context) {
	ids, isOk := c.getUserAuthIds(g)
	if !isOk {
		log.Fatal("获取ids失败", ids, isOk)
		c.ResponseData(g, nil, http.DataFieldInvalid, "获取钱包id失败")
		return
	}

	userStruct := service.GetUserInfo(g)
	//fmt.Println(userStruct)
	if userStruct.ID == 0 {
		log.Fatal("获取用户信息失败", userStruct)
		c.ResponseData(g, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}

	data, errs := service.WalletList(ids, userStruct, Param.PageNo, Param.PageSize, Param.Status)

	if errs != nil {
		log.Fatal("获取WalletList失败", errs.Error())
		c.ResponseData(g, nil, http.BadRequest, "获取WalletList失败")
		return
	}

	c.ResponseData(g, data, http.Ok)
	return
}

//@Summary 编辑钱包状态
//@Produce json
// @Param	body	body	request.WalletUpdate true "body参数"
//@Success 200   "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/update [put]
func (c *WalletController) Update(g *gin.Context) {
	var Param request.WalletUpdate
	err := g.BindJSON(&Param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}

	ids, isOk := c.getUserAuthIds(g)
	if !isOk {
		log.Fatal("获取ids失败", ids, isOk)
		c.ResponseData(g, nil, http.Forbidden, "获取钱包id失败")
		return
	}

	//判断是否有钱包权限
	isAuth := utils.IsInArray(Param.WalletID, ids)
	fmt.Println("wid  wallet:", Param.WalletID)
	fmt.Println("ids:", ids)
	if !isAuth {
		c.ResponseData(g, nil, http.Forbidden, "没有权限操作该钱包")
		return
	}
	err = service.WalletUpdate(Param)
	if err != nil {
		log.Fatal("wallet_id:", Param.WalletID, "status", Param.Status, err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, "数据解析错误")
		return
	}
	c.ResponseData(g, nil, http.Ok)
	return
}

//@Summary 上链新建普通钱包
//@Produce json
// @Param	body	body	request.WalletNormalNew true "body参数"
//@Success 200   "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/create [post]
func (c *WalletController) ChainCreate(g *gin.Context) {
	var Param request.WalletNormalNew
	err := g.BindJSON(&Param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}

	data, err := service.ChainNewWallet()
	if err != nil {
		log.Fatal(err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, "节点返回出错")
		return
	}
	//保存钱包
	userStruct := service.GetUserInfo(g)
	if userStruct.ID == 0 {
		log.Fatal("获取用户信息失败", userStruct)
		c.ResponseData(g, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}
	var wallet model.Wallet
	wallet.Address = data.Data.Address
	wallet.TypeID = 2
	wallet.CoinID = 1
	wallet.IsCollect = dao.CollectNo
	wallet.RequiredSigner = 1
	wallet.Name = Param.Name
	wallet.CompanyID = userStruct.CompanyID
	wallet.Status = dao.MultiSignWalletStatusSuccess
	//save
	_, errs := dao.GlobalDao.WalletAdd(wallet, userStruct.ID)
	if errs != nil {
		log.Fatal("添加钱包出错：" + err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, err.Error())
		return
	}
	var ki response.KeyInfo
	err = json.Unmarshal(data.Data.PrivateKey, &ki)
	if err != nil {
		log.Fatal("解析钱包私钥出错：" + err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, err.Error())
		return
	}
	data.Data.PrivateKey = ki.PrivateKey

	c.ResponseData(g, data.Data, http.Ok)
}

//@Summary 添加钱包
//@Produce json
// @Param	body	body	request.WalletAdd true "body参数"
//@Success 200   "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/add [post]
func (c *WalletController) Add(g *gin.Context) {
	var wallet model.Wallet
	var Param request.WalletAdd
	err := g.BindJSON(&Param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}

	userStruct := c.getUserInfo(g)
	if userStruct.ID == 0 {
		log.Fatal("获取用户信息失败", userStruct)
		c.ResponseData(g, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}

	//钱包地址检查是否已存在
	if service.WalletIsExist(Param.Address) {
		log.Fatal("钱包地址:" + Param.Address + "已存在")
		c.ResponseData(g, nil, http.BadRequest, "该钱包地址已存在")
		return
	}

	//检查钱包名是否有存在
	var num int64
	num = service.WalletCountByName(Param.Name)
	if num > 0 {
		log.Fatal("钱包名:" + Param.Name + "已存在")
		c.ResponseData(g, nil, http.DataExisted, "钱包名已存在")
		return
	}

	//币种有效性判断 && 设置coin_id
	var coin model.Coin
	coin, err = service.GetCoin(Param.Symbol)
	if err != nil {
		log.Fatal(err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, err.Error())
		return
	}
	if coin.ID == 0 {
		log.Fatal("币种类型有误,传入参数：", Param.Symbol)
		c.ResponseData(g, nil, http.DataFieldInvalid, "币种类型有误")
		return
	}
	wallet.CoinID = coin.ID

	if Param.Type == dao.MultiWallet {
		// 多签
		// 校验签名&关联钱包
		//ids := strings.Split(Param.OrdinaryWalletId, ",")
		if Param.RequiredSigner < 1 || Param.RequiredSigner > uint(len(Param.OrdinaryWalletId)) {
			c.ResponseData(g, nil, http.BadRequest, "签名数错误")
			return
		}

		if len(Param.OrdinaryWalletId) < 1 {
			c.ResponseData(g, nil, http.BadRequest, "最少关联一个钱包")
			return
		}
		err = service.MsWalletAdd(Param, wallet, userStruct)
	} else {
		//普通钱包：需检查地址有效性
		addr, err := address.NewFromString(Param.Address)
		if err != nil {
			c.ResponseData(g, nil, http.BadRequest, "钱包地址格式错误", err.Error())
			return
		}
		if &addr != nil {
			if &addr != nil && addr.Protocol() != address.SECP256K1 {
				c.ResponseData(g, nil, http.BadRequest, "钱包地址类型错误，只支持SECP256K1")
				return
			}
		} else {
			c.ResponseData(g, nil, http.BadRequest, "钱包地址校验异常")
			return
		}

		err = service.NormalWalletAdd(Param, wallet, userStruct)
	}

	if err != nil {
		log.Fatal(err.Error())
		c.ResponseData(g, nil, http.BadRequest, "添加钱包异常", err.Error())
		return
	}
	c.ResponseData(g, nil, http.Ok)
	return
}

//@Summary 开通归集/重置授权
//@Produce json
// @Param	body	body	request.OpenCollect true "body参数"
//@Success 200 {object} response.CollectKeyNew "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/collect/enable [post]
func (c *WalletController) EnableCollect(g *gin.Context) {
	var Param request.OpenCollect
	err := g.BindJSON(&Param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}

	var wallet model.WalletPrivate
	var w model.Wallet
	dao.DB.Table(w.TableName()).Where("id = ?", Param.ID).First(&wallet)

	tString := fmt.Sprint(time.Now())
	wallet.AccessKey = utils.GenerateMd5(tString + strconv.Itoa(int(wallet.ID)))
	wallet.SecretKey = utils.Hmac256AndBase256(wallet.AccessKey, fmt.Sprint(time.Now()))
	wallet.Sign = utils.Sha256(wallet.AccessKey + wallet.SecretKey)
	wallet.IsCollect = dao.CollectYes
	dao.DB.Table(w.TableName()).Where("id = ?", Param.ID).Save(&wallet)
	var req response.CollectKeyNew
	req.AccessKey = wallet.AccessKey
	req.SecretKey = wallet.SecretKey
	c.ResponseData(g, req, http.Ok)
	return
}

//@Summary 获取钱包个数
//@Produce json
//@Success 200 {object} response.WalletCount "成功"
//@Failure 400  "请求错误"
//@Failure 500  "内部错误"
//@Router /api/v1/wallet/count [post]
func (c *WalletController) Count(g *gin.Context) {
	userStruct := c.getUserInfo(g)
	if userStruct.ID == 0 {
		log.Fatal("获取用户信息失败", userStruct)
		c.ResponseData(g, nil, http.DataFieldInvalid, "获取用户信息失败")
		return
	}
	count, err := service.Count(userStruct)
	if err != nil {
		log.Fatal("获取钱包个数失败", err.Error())
		c.ResponseData(g, nil, http.DataFieldInvalid, "获取用户信息失败", err.Error())
		return
	}
	c.ResponseData(g, count, http.Ok)
	return
}

//@Summary 生成api公私钥
//@Produce json
//@Param id query string false "钱包id"
//@Success 200 {object} response.CollectKeyNew "成功"
//@Failure 400   "请求错误"
//@Failure 500   "内部错误"
//@Router /api/v1/wallet/collect/key [post]
func (c *WalletController) CollectKey(g *gin.Context) {
	var param request.GetNewCollectKey
	err := g.BindJSON(&param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}
	var resp response.CollectKeyNew
	resp.AccessKey, resp.SecretKey = service.CollectNewKey(param.ID)
	c.ResponseData(g, resp, http.Ok)
	return
}

//@Summary 用户钱包权限列表
//@Produce json
//@Failure 400   "请求错误"
//@Failure 500   "内部错误"
//@Router /api/v1/wallet/auth/ids [post]
func (c *WalletController) WalletAuthIds(g *gin.Context) {
	var param request.GetWalletAuthList
	err := g.BindJSON(&param)
	if err != nil {
		c.CommonValidateReturn(g, err)
		return
	}
	var wids []int
	dao.DB.Table(new(model.WalletAuthorize).TableName()).Where("company_member_id = ?", param.ID).Pluck("wallet_id", &wids)
	c.ResponseData(g, wids, http.Ok)
	return
}
