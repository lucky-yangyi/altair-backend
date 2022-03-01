package dao

import (
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/log"

	"errors"
)

const (
	// TransactionStatusDoing 交易已发起
	TransactionStatusDoing uint8 = 1
	// TransactionStatusFailed 交易失败
	TransactionStatusFailed uint8 = 4
	// TransactionStatusSuccess 交易成功
	TransactionStatusSuccess uint8 = 127

	// TransactionTypeOut 交易类型：转出
	TransactionTypeOut uint8 = 1
	//TransactionTypeIn 交易类型：转入
	TransactionTypeIn uint8 = 2
	//TransactionTypeOther 交易类型：其它
	TransactionTypeOther uint8 = 3

	//TransactionSubTypeDefault 交易子类型：缺省
	TransactionSubTypeDefault uint8 = 1
	//TransactionSubTypeCreateWallet 交易子类型：创建多签钱包
	TransactionSubTypeCreateWallet uint8 = 2
	//TransactionSubTypeUpdateWalletSignNum 交易子类型：修改多签钱包签名数
	TransactionSubTypeUpdateWalletSignNum uint8 = 3

	NormalWallet uint8 = 2 //普通钱包
	MultiWallet  uint8 = 1 //多签
	FIL          uint8 = 1
)

//AddTransaction 添加交易
func (d *Dao) AddTransaction(tn model.Transaction, multiId uint64) (TransID uint64, err error) {
	//检查userId是否合法
	var user model.CompanyMember
	err = d.dao.Model(user).Where("id = ?", tn.CompanyMemberID).First(&user).Error
	if err != nil {
		log.Fatal("检查userId是否合法", err.Error())
		return 0, err
	} else if user.ID == 0 {
		return 0, errors.New("用户不存在")
	}

	//事务
	tx := d.dao.Begin()
	err = tx.Create(&tn).Error
	if err != nil {
		log.Fatal("添加交易事务", err.Error())
		tx.Rollback()
		return 0, err
	}
	//增加对应的detail
	//判断是否多签
	var wallet model.Wallet
	d.dao.Table(wallet.TableName()).Where("id = ?", tn.FromWalletID).First(&wallet)
	if wallet.TypeID == MultiWallet {
		//多签交易
		var members []model.MultiSignWalletDetail

		err = tx.Model(model.MultiSignWalletDetail{}).Where("wallet_id = ?", multiId).Find(&members).Error
		if err != nil {
			log.Fatal("添加交易 -> 拉取 MultiSignWalletMember ->", err.Error())
			tx.Rollback()
			return 0, err
		}
		for k, v := range members {
			var detail model.TransactionDetail
			detail.TransactionID = tn.ID
			detail.WalletID = v.OrdinaryWalletID
			detail.Status = TransactionStatusDoing
			//第一位指定为发起者
			if k == 0 {
				detail.IsSponsor = 1
			}
			err = tx.Table(detail.TableName()).Create(&detail).Error
			if err != nil {
				log.Fatal("添加交易 -> 添加详情 ->", err.Error())
				tx.Rollback()
				return 0, err // c.ResponseData(nil, http.BadRequest)
			}
		}
	} else {
		//普通钱包交易
		var detail model.TransactionDetail
		detail.TransactionID = tn.ID
		detail.WalletID = tn.FromWalletID
		detail.Status = TransactionStatusDoing
		detail.IsSponsor = 1
		err = tx.Table(detail.TableName()).Create(&detail).Error
	}

	tx.Commit()
	return tn.ID, nil
}

func (d *Dao) GetTransactionList(params request.TransactionList, offset, size uint64) (list []*model.MixTransaction, err error) {
	var par []interface{}
	sql := `SELECT * FROM transaction as t 
		    LEFT JOIN wallet w ON w.id = t.from_wallet_id
		 	LEFT JOIN coin c ON c.id = t.coin_id
			LEFT JOIN company_member cm ON cm.id = t.company_member_id WHERE t.is_del = 0`
	if params.WalletId != 0 {
		sql += " AND w.id = ?"
		par = append(par, params.WalletId)
	}
	if params.Type != 0 {
		sql += " AND t.type = ?"
		par = append(par, params.Type)
	}
	if params.StartTime != 0 {
		sql += " AND t.start_time > ?"
		par = append(par, params.StartTime)
	}
	if params.EndTime != 0 {
		sql += " AND t.end_time < ?"
		par = append(par, params.EndTime)
	}
	sql += " LIMIT ?,?"
	par = append(par, offset)
	par = append(par, size)
	err = d.dao.Raw(sql, par...).Find(&list).Error
	return
}

func (d *Dao) CountTransaction(params request.TransactionList) (count int64, err error) {
	var par []interface{}
	sql := `SELECT count(1) FROM transaction as t 
		    LEFT JOIN wallet w ON w.id = t.from_wallet_id
		 	LEFT JOIN coin c ON c.id = t.coin_id
			LEFT JOIN company_member cm ON cm.id = t.company_member_id  WHERE t.is_del = 0`
	if params.WalletId != 0 {
		sql += " AND w.id = ?"
		par = append(par, params.WalletId)
	}
	if params.Type != 0 {
		sql += " AND t.type = ?"
		par = append(par, params.Type)
	}
	if params.StartTime != 0 {
		sql += " AND t.start_time > ?"
		par = append(par, params.StartTime)
	}
	if params.EndTime != 0 {
		sql += " AND t.end_time < ?"
		par = append(par, params.EndTime)
	}
	err = d.dao.Raw(sql, par...).Count(&count).Error
	return
}

//获取交易详情列表
func (d *Dao) GetTransactionDetail(transactionId uint64) (list *model.TransactionDetail, err error) {
	err = d.dao.Model(model.TransactionDetail{}).Where("transaction_id = ? ", transactionId).Preload("Wallet").Preload("Transaction").Find(&list).Error
	return
}
