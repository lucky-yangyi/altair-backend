package request

type BillListReq struct {
	WalletId  int    `json:"walletId"`
	Month     string `json:"month"`
	PayStatus int    `json:"payStatus"`
	PageReq
}

type StatusBillReq struct {
	Id        int `json:"id"`
	PayStatus int `json:"payStatus"`
}
