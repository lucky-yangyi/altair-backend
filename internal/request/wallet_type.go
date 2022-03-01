package request

type WalletTypeListReq struct {
	Name string `json:"name"`
	Code string `json:"code"`
	PageReq
}
