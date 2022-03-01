package request

type MonthBillListReq struct {
	WalletId  int    `json:"wallet_id"`
	Month     string `json:"month"`
	PayStatus int    `json:"pay_status"`
	PageReq
}
type CreateBillReq struct {
	CompanyId uint64 `json:"company_id"`
}
