package task

import (
	"altair-backend/internal/dao"
	"altair-backend/internal/model"
	"altair-backend/internal/request"
	"altair-backend/internal/service"
	"altair-backend/log"
	"fmt"
)

type Response struct {
	Data    Data `json:"data"`
	Code    int  `json:"code"`
	Success bool `json:"success"`
}

type Data struct {
	Balance float64 `json:"balance"`
}

func getBalance() {
	var data request.Balance
	data.Symbol = "FIL"
	var wallet []model.Wallet

	dao.DB.Find(&wallet)
	for _, v := range wallet {
		data.Address = v.Address
		go GetAndSave(v, data)
	}
}

func GetAndSave(v model.Wallet, data request.Balance) {
	res, err := service.GetBalance(data)
	if err != nil {
		log.Error("获取余额失败：", err.Error())
		return
	}

	walletInfo := fmt.Sprintf("ID:%d, 地址:%s", v.ID, v.Address)

	if v.Balance != res.Data.Balance {
		v.Balance = res.Data.Balance
		dao.DB.Save(&v)
		fmt.Println("钱包余额更新成功", walletInfo)
	} else {
		fmt.Println("钱包余额未改变", walletInfo)
	}
}
