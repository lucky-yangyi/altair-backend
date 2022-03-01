package dao

import (
	"altair-backend/internal/model"
)

func (d *Dao) WalletCollect(wid uint64) (wc model.WalletCollect, err error) {
	err = d.dao.Table(wc.TableName()).Where("wallet_id = ?", wid).Find(&wc).Error
	return
}

//WalletCollectNew
func (d *Dao) WalletCollectUpdate(wid uint64) (AccessKey, SecretKey string, err error) {
	var wc model.WalletCollect
	var wcp model.WalletCollectPrivate
	query := d.dao.Table(wc.TableName())
	err = query.Where("wallet_id = ?", wid).Find(&wcp).Error
	if err != nil {
		return
	}
	err = query.Save(&wcp).Error
	return
}

//func (d *Dao) WalletCollectNew(wid uint64) (AccessKey, SecretKey string, err error) {
//	var wc model.WalletCollect
//	// 上链申请一个普通钱包
//	data := service.ChainNewWallet()
//
//	AccessKey = utils.GenerateMd5(fmt.Sprint(time.Now()) + strconv.Itoa(int(wid)))
//	SecretKey = utils.Hmac256AndBase256(AccessKey, fmt.Sprint(time.Now()))
//	wcp := model.WalletCollectPrivate{
//		WalletID: int64(wid),
//		Address: data.Data.Address,//todo
//		PrivateKey: string(data.Data.PrivateKey),//todo
//	}
//	err = d.dao.Table(wc.TableName()).Save(&wcp).Error
//	return
//}
