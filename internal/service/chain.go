package service

import (
	"altair-backend/config"
	"altair-backend/internal/request"
	"altair-backend/internal/response"
	"altair-backend/log"
	"altair-backend/pkg/curl"
	"altair-backend/pkg/utils"
	"encoding/json"
	"errors"
)

func GetBalance(param request.Balance) (data response.Balance, err error) {
	// 发送请求
	returns, err := curl.ChainPost(config.ServerConfig.SeverHost+"/api/Addresses/Wallet/Balance", utils.StructToMap(param), "application/json")
	if err != nil {
		return data, errors.New(returns)
	}
	err = json.Unmarshal([]byte(returns), &data)
	if err != nil {
		return data, err
	}
	data.Data.Balance = data.Data.Balance / 1000000000000000000
	return data, nil
}

func ChainNewWallet() (data response.WalletNew, err error) {
	// 发送请求
	returns, err := curl.ChainPost(config.ServerConfig.SeverHost+"/api/Addresses/Wallet/New", utils.StructToMap(nil), "application/json")
	if err != nil {
		log.Fatal("节点请求失败：" + err.Error())
		return data, err
	}
	err = json.Unmarshal([]byte(returns), &data)
	if err != nil {
		log.Fatal("解析失败：" + err.Error())
		return data, err
	}
	return
}

func ChainNormalTransaction(param request.ChainCollectTransaction) (data response.CollectTransactionResponse, err error) {
	// 发送请求
	returns, err := curl.ChainPost(config.ServerConfig.SeverHost+"/api/Addresses/Wallet/Transaction/Collect", utils.StructToMap(param), "application/json")
	if err != nil {
		log.Fatal("节点请求失败：" + err.Error() + "||||" + returns)
		return data, errors.New(err.Error() + "||||" + returns)
	}
	err = json.Unmarshal([]byte(returns), &data)

	if err != nil {
		log.Fatal("解析失败：" + err.Error())
		data.Code = 400
		return data, errors.New("解析失败：" + err.Error())
	}
	return
}

func Receipt(param request.Receipt) (data response.CommonResponse, err error) {
	// 发送请求
	returns, err := curl.ChainPost(config.ServerConfig.SeverHost+"/api/Transactions/Receipt", utils.StructToMap(param), "application/json")
	if err != nil {
		log.Fatal("节点请求失败：" + err.Error() + "||||" + returns)
		return data, errors.New(err.Error() + "||||" + returns)
	}
	err = json.Unmarshal([]byte(returns), &data)

	if err != nil {
		log.Fatal("解析失败：" + err.Error())
		data.Code = 400
		return data, errors.New("解析失败：" + err.Error())
	}
	return
}

func SendRawTransaction(param request.SendRawTransaction) (data response.CollectTransactionResponse, err error) {
	// 发送请求
	returns, err := curl.ChainPost(config.ServerConfig.SeverHost+"/api/Transactions/SendRawTransaction", utils.StructToMap(param), "application/json")
	if err != nil {
		log.Fatal("节点请求失败：" + err.Error() + "||||" + returns)
		return data, errors.New(err.Error() + "||||" + returns)
	}
	err = json.Unmarshal([]byte(returns), &data)

	if err != nil {
		log.Fatal("解析失败：" + err.Error())
		data.Code = 400
		return data, errors.New("解析失败：" + err.Error())
	}
	return
}

func SingleTransaction(param request.SendFromOffline) (data response.CommonResponse, err error) {
	// 发送请求
	param.SignType = "createSingleTransaction"
	param.Symbol = "FIL"
	//var dataStruct requests.DataStruct
	returns, err := curl.ChainPost(config.ServerConfig.SeverHost+"/api/Transactions/SendFromOffline", utils.StructToMap(param), "application/json")
	if err != nil {
		log.Fatal("节点请求失败：" + err.Error() + "||||" + returns)
		return data, errors.New(err.Error() + "||||" + returns)
	}
	err = json.Unmarshal([]byte(returns), &data)

	if err != nil {
		log.Fatal("解析失败：" + err.Error())
		data.Code = 400
		return data, errors.New("解析失败：" + err.Error())
	}
	return
}

func CreateMulWallet(param request.SendFromOffline) (data response.CommonResponse, err error) {
	// 发送请求
	param.SignType = "createMultisigWalletTransaction"
	param.Symbol = "FIL"
	//	var dataStruct requests.CreateMsig
	returns, err := curl.ChainPost(config.ServerConfig.SeverHost+"/api/Transactions/SendFromOffline", utils.StructToMap(param), "application/json")
	if err != nil {
		log.Fatal("节点请求失败：" + err.Error() + "||||" + returns)
		return data, errors.New(err.Error() + "||||" + returns)
	}
	err = json.Unmarshal([]byte(returns), &data)

	if err != nil {
		log.Fatal("解析失败：" + err.Error())
		data.Code = 400
		return data, errors.New("解析失败：" + err.Error())
	}
	return
}

func ProposeMulCreate(param request.SendFromOffline) (data response.CommonResponse, err error) {
	// 发送请求
	param.SignType = "proposeMultisigTransaction"
	param.Symbol = "FIL"
	//var dataStruct requests.ProposeCreateMsig
	returns, err := curl.ChainPost(config.ServerConfig.SeverHost+"/api/Transactions/SendFromOffline", utils.StructToMap(param), "application/json")
	if err != nil {
		log.Fatal("节点请求失败：" + err.Error() + "||||" + returns)
		return data, errors.New(err.Error() + "||||" + returns)
	}
	err = json.Unmarshal([]byte(returns), &data)

	if err != nil {
		log.Fatal("解析失败：" + err.Error())
		data.Code = 400
		return data, errors.New("解析失败：" + err.Error())
	}
	return
}
