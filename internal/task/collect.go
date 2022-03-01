package task

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/log"
	"encoding/base64"
	"fmt"
)

// CollectTransactionTask 定时扫描并执行归集
func CollectTransactionTask() {
	fmt.Println("开始执行归集")

	var collects []model.WalletCollectPrivate
	var collect model.WalletCollect
	dao.DB.Table(collect.TableName()).Find(&collects)
	for _, v := range collects {
		//1查询余额
		param := request.Balance{
			Symbol:  "FIL",
			Address: v.Address,
		}
		data, err := service.GetBalance(param)
		if err != nil {
			log.Fatal("子钱包地址：" + v.Address + "定时获取余额失败：" + err.Error())
			continue
		}
		if data.Data.Balance <= 0.001 {
			continue
		}
		fmt.Println("地址："+v.Address, "余额：", data.Data.Balance)

		privateByte, _ := base64.StdEncoding.DecodeString(v.PrivateKey)
		//// 执行交易
		var wallet model.Wallet
		dao.DB.Table(wallet.TableName()).Where("id = ?", v.WalletID).First(&wallet)
		params := request.ChainCollectTransaction{
			FromAddr:   v.Address,
			PrivateKey: privateByte,
			ToAddr:     wallet.Address,
		}
		resp, err := service.ChainNormalTransaction(params)
		if resp.Code != 200 {
			log.Fatal("v.Address:" + v.Address + "////错误信息：" + err.Error())
		}
		fmt.Println("子钱包地址："+v.Address, "交易回执：", resp)
		//保存数据
		var save = model.CollectTransaction{
			FromAddress: v.Address,
			ToAddress:   wallet.Address,
			Amount:      resp.Data.Amount,
			Cid:         resp.Data.Cid,
		}
		dao.DB.Table(save.TableName()).Create(&save)

	}

}
