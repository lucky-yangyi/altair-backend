package service

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/log"
)

// UpdateTransactionDetail 修改交易详情
func UpdateTransactionDetail(param request.UpdateTransactionDetail) (ok bool, message string) {
	message = "更新失败"

	if param.GasAmount < 0 {
		return false, "Gas费用错误"
	}

	// 检查状态有效性
	statusList := [...]uint8{
		dao.TransactionStatusSuccess,
		dao.TransactionStatusDoing,
		dao.TransactionStatusFailed,
	}

	contains := false
	for k := range statusList {
		if param.Status == statusList[k] {
			contains = true
			break
		}
	}
	if !contains {
		return false, "状态不支持"
	}

	// 检查交易
	var trans model.Transaction
	if err := dao.DB.First(&trans, param.TransactionID).Error; err != nil {
		return false, "交易 ID 错误"
	}
	if trans.Status == dao.TransactionStatusSuccess || trans.Status == dao.TransactionStatusFailed {
		return false, "交易已结束，请勿重复提交"
	}

	// 检查钱包 ID
	var tmp model.Wallet
	if err := dao.DB.First(&tmp, param.WalletID).Error; err != nil {
		return false, "钱包 ID 错误"
	}

	// 多签钱包
	var wallet model.Wallet
	dao.DB.Table(wallet.TableName()).First(&wallet, trans.FromWalletID)

	var detail model.TransactionDetail
	var err error

	err = dao.DB.Model(detail).Where("transaction_id = ?", param.TransactionID).Where("wallet_id = ?", param.WalletID).First(&detail).Error
	if err != nil {
		return false, "交易详情不存在"
	}

	// 状态检查
	if detail.Status == param.Status {
		return false, "交易状态重复更新"
	}
	if detail.Status != dao.TransactionStatusDoing {
		return false, "交易非待签名状态，不允许签名"
	}

	detail.GasAmount = param.GasAmount
	detail.Status = param.Status
	detail.ChainTransactionID = param.ChainTransactionId

	// 开始事务
	tx := dao.DB.Begin()

	err = tx.Save(&detail).Error

	if err != nil {
		log.Fatal("更新交易详情失败", err.Error())
		tx.Rollback()
		return false, message
	}

	if wallet.TypeID == dao.MultiWallet {
		//多签
		// 需要更新钱包 Status 字段：交易子类型为（创建钱包 或 修改签名数）
		needChangeWalletStatus := false
		lastedStatus := wallet.Status

		//创建多签或者修改签名数时，可直接修改该交易状态
		if trans.SubType == dao.TransactionSubTypeCreateWallet || trans.SubType == dao.TransactionSubTypeUpdateWalletSignNum {
			needChangeWalletStatus = true
		}

		if param.Status == dao.TransactionStatusFailed {
			trans.Status = dao.TransactionStatusFailed
			// 修改钱包状态
			if needChangeWalletStatus {
				lastedStatus = dao.MultiSignWalletStatusFailed
			}
		} else {

			var num int64
			tx.Table(detail.TableName()).Where("transaction_id = ?", param.TransactionID).Where("status = ?", dao.TransactionStatusSuccess).Count(&num)

			if trans.SubType == dao.TransactionSubTypeCreateWallet && num == 1 {
				trans.Status = dao.TransactionStatusSuccess
				lastedStatus = dao.MultiSignWalletStatusSuccess
			} else {
				// 满足最小签名数
				if num >= int64(wallet.RequiredSigner) {
					// 更新交易状态
					trans.Status = dao.TransactionStatusSuccess

					// 修改钱包状态
					if needChangeWalletStatus {
						lastedStatus = dao.MultiSignWalletStatusSuccess
					}
				}
			}

		}

		// 更新钱包数据
		if needChangeWalletStatus && wallet.Status != lastedStatus {
			wallet.Status = lastedStatus
			err = tx.Save(&wallet).Error
			if err != nil {
				log.Fatal("更新钱包失败", err.Error())
				tx.Rollback()
				return false, message
			}
		}
	} else {
		if param.Status == dao.TransactionStatusFailed {
			trans.Status = dao.TransactionStatusFailed
		} else {
			trans.Status = dao.TransactionStatusSuccess
		}
	}

	trans.GasAmount += param.GasAmount

	err = tx.Save(&trans).Error
	if err != nil {
		log.Error("更新交易失败", err.Error())
		tx.Rollback()
		return false, message
	}

	tx.Commit()
	return true, message
}
