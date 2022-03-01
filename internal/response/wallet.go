package response

type WalletCount struct {
	Total   int64 `json:"total"`   //总数
	Normal  int64 `json:"normal"`  //普通钱包数
	Collect int64 `json:"collect"` //归集钱包数
}

type CollectKeyNew struct {
	AccessKey string `json:"accessKey"` //公钥
	SecretKey string `json:"secretKey"` //私钥
}

type GetCollectKey struct {
	AccessKey string `json:"access_key"` //公钥
}

type Balance struct {
	Data    BalanceData `json:"data"`
	Code    int         `json:"code"`
	Success bool        `json:"success"`
}

type BalanceData struct {
	Balance float64 `json:"balance"`
}

type WalletNew struct {
	Data    New  `json:"data"`
	Code    int  `json:"code"`
	Success bool `json:"success"`
}

type New struct {
	Address    string `json:"address"`
	PrivateKey []byte `json:"privateKey"`
}
