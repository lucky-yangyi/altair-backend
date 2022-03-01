package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/log"
	"altair-backend/pkg/utils"
	"encoding/base64"
	"fmt"
	uuid "github.com/satori/go.uuid"
	"strconv"
	"time"
)

func WalletIsExist(address string) bool {
	return dao.GlobalDao.WalletIsExist(address)
}

//多签添加
func MsWalletAdd(param request.WalletAdd, wallet model.Wallet, user model.CompanyMember) (err error) {

	//验证关联钱包ID
	ids := param.OrdinaryWalletId
	for _, v := range ids {
		//oId, _ := strconv.Atoi(v)
		// 验证关联钱包ID
		var tmp model.Wallet
		if err := dao.DB.First(&tmp, v).Error; err != nil {
			log.Fatal("验证关联钱包ID失败", err.Error())
			return err
		}
	}
	wallet.Address = uuid.NewV4().String()
	wallet.Status = dao.MultiSignWalletStatusDoing
	wallet.TypeID = param.Type
	wallet.Name = param.Name
	wallet.CompanyID = user.CompanyID
	wallet.RequiredSigner = param.RequiredSigner
	wallet.IsCollect = dao.CollectNo
	//save
	wallet.ID, err = dao.GlobalDao.WalletAdd(wallet, user.ID)
	if err != nil {
		return err
	}
	var filWalletMeta model.FilWalletMeta
	filWalletMeta.WalletID = wallet.ID
	filWalletMeta.IsMultiWallet = wallet.TypeID // 1：普通，2：多签
	filWalletMeta.RequiredSigner = param.RequiredSigner
	_, err = dao.GlobalDao.FilWalletMetaAdd(filWalletMeta)
	// 添加对应的detail
	// save multi_sign_wallet_detail
	for _, v := range ids {
		wd := model.MultiSignWalletDetail{
			WalletID:         wallet.ID,
			OrdinaryWalletID: uint64(v),
			Status:           dao.MultiSignWalletStatusDoing, //待签名
		}
		err = dao.GlobalDao.MultiSignWalletDetailAdd(wd)
		if err != nil {
			return err
		}
	}

	// 属性为多签钱包时，则需要添加交易记录
	trans := model.Transaction{
		FromWalletID:    wallet.ID,
		Type:            dao.TransactionTypeOther,
		SubType:         dao.TransactionSubTypeCreateWallet,
		Comment:         "创建多签钱包",
		CompanyMemberID: user.ID,
		Status:          dao.TransactionStatusDoing,
		CoinID:          dao.FIL, //todo::FIL
	}

	_, err = dao.GlobalDao.AddTransaction(trans, wallet.ID)
	// 生成交易记录
	if err != nil {
		return err
	}

	return nil
}

func NormalWalletAdd(param request.WalletAdd, wallet model.Wallet, user model.CompanyMember) (err error) {
	wallet.Address = param.Address
	wallet.TypeID = param.Type
	wallet.Name = param.Name
	wallet.CompanyID = user.CompanyID
	wallet.Status = dao.MultiSignWalletStatusSuccess
	wallet.IsCollect = dao.CollectNo
	//save
	_, errs := dao.GlobalDao.WalletAdd(wallet, user.ID)
	if errs != nil {
		return errs
	}

	return nil
}

func Count(member model.CompanyMember) (count response.WalletCount, err error) {
	if member.IsAdmin == true {
		count.Total, count.Normal, count.Collect, err = dao.GlobalDao.WalletAdminCount(member.CompanyID)
	} else {
		count.Total, count.Normal, count.Collect, err = dao.GlobalDao.WalletMemberCount(member.ID)
	}
	return
}

func WalletList(ids []int, user model.CompanyMember, pageNo, pageSize uint64, status uint8) (data utils.Page, err error) {
	return dao.GlobalDao.WalletList(ids, user, pageNo, pageSize, status)
}

func WalletListNoPage(user model.CompanyMember) (list []model.WalletAuth, err error) {
	return dao.GlobalDao.WalletListNoPage(user)
}

func WalletUpdate(update request.WalletUpdate) (err error) {
	return dao.GlobalDao.WalletUpdate(update)
}

func CollectNewKey(wid uint64) (AccessKey, SecretKey string) {
	AccessKey = utils.GenerateMd5(fmt.Sprint(time.Now()) + strconv.Itoa(int(wid)))
	SecretKey = utils.Hmac256AndBase256(AccessKey, fmt.Sprint(time.Now()))
	Sign := utils.Sha256(AccessKey + SecretKey) //http://tools.jb51.net/password/sha_encode
	//存表
	var wallet model.Wallet
	dao.DB.Model(wallet).Where("id = ?", wid).Update("access_key", AccessKey).Update("sign", Sign)
	return
}

func WalletCollectAdd(wid uint64, symbol string) (address string, err error) {
	var wc model.WalletCollect
	// 上链申请一个普通钱包
	data, err := ChainNewWallet()
	if err != nil {
		log.Fatal(err.Error(), "上链申请出错")
		return
	}
	PrivateKey := base64.StdEncoding.EncodeToString(data.Data.PrivateKey)
	wcp := model.WalletCollectPrivate{
		WalletID:   int64(wid),
		Address:    data.Data.Address, //todo
		PrivateKey: PrivateKey,        //todo
		Symbol:     symbol,
	}
	err = dao.DB.Table(wc.TableName()).Save(&wcp).Error
	return data.Data.Address, err
}

func WalletCountByName(name string) (num int64) {
	dao.DB.Table(new(model.Wallet).TableName()).Where("name = ?", name).Count(&num)
	return
}

func GetCoin(code string) (coin model.Coin, err error) {
	return dao.GlobalDao.GetCoin(code)
}

func WalletCollect(wid uint64) (wc model.WalletCollect, err error) {
	return dao.GlobalDao.WalletCollect(wid)
}

func WalletCollectUpdate(wid uint64) (AccessKey, SecretKey string, err error) {
	return dao.GlobalDao.WalletCollectUpdate(wid)
}
