package request

type WalletCount struct {
	Total   int `json:"total"`
	Normal  int `json:"normal"`
	Collect int `json:"collect"`
}

type CollectKeyNew struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type GetCollectKey struct {
	AccessKey string `json:"accessKey"`
	SecretKey string `json:"secretKey"`
}

type WalletList struct {
	PageReq
	Status uint8 `json:"status"`
}

type GetNewCollectKey struct {
	ID uint64 `json:"id"`
}

type GetWalletAuthList struct {
	ID uint64 `json:"id"` //用户id
}

// Balance 获得余额
type Balance struct {
	Symbol  string `json:"symbol"`  //币种标识
	Address string `json:"address"` //地址
}

type WalletUpdate struct {
	WalletID int    `json:"walletId" binding:"required"` //钱包id
	Status   int    `json:"status" binding:"required"`   //签名状态：1-待签，4-失败，127-成功
	Address  string `json:"address" binding:"required"`  //钱包地址
}

type WalletAdd struct {
	Symbol           string `json:"symbol" binding:"required"` //钱包类型,默认传"FIL"
	Name             string `json:"name" binding:"required"`   //钱包名称
	Address          string `json:"address"`                   //钱包地址
	Type             uint8  `json:"type" binding:"required"`   //钱包种类:1-多签，2-普通
	RequiredSigner   uint   `json:"requiredSigner" `           //签名人数（多签钱包时传）
	OrdinaryWalletId []int  `json:"ordinaryWalletId"`          //多个id时，逗号分开
}

type WalletNormalNew struct {
	Symbol string `json:"symbol" binding:"required"` //钱包类型,默认传"FIL"
	Name   string `json:"name" binding:"required"`   //钱包名称
}

type OpenCollect struct {
	ID uint64 `json:"id" binding:"required"` //钱包id
}
