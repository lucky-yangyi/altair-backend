package model

// todo demo
type MixMemberCompany struct {
	Company
	member []CompanyMember
	auth   []WalletAuthorize
}

type MixMonthBillWallet struct {
	MonthBill
	Wallet
}

type MixTransaction struct {
	Transaction
	Wallet
	Coin
}

type MixUser struct {
	CompanyMember CompanyMember `json:"company_member"`
	Admin         Admin         `json:"admin"`
	AccessKey     string        `json:"accessKey"` // 用于登出
	System        int           `json:"system"`
}
