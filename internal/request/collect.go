package request

type AddWalletCollect struct {
	Symbol string `json:"symbol"` //标识
}

type CollectTransaction struct {
	Address string  `json:"address"` //转出地址
	Amount  float64 `json:"amount"`  //金额
	Comment string  `json:"comment"` //备注
}

type CollectTest struct {
	SecretKey string `json:"secretKey"` //转出地址
	AccessKey string `json:"accessKey"` //金额
}
