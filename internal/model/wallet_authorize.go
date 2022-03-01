package model

// WalletAuthorize 钱包权限表
type WalletAuthorize struct {
	Mixin
	WalletID        uint64 `gorm:"column:wallet_id" json:"walletId"`                // 多签钱包ID
	CompanyMemberID uint64 `gorm:"column:company_member_id" json:"companyMemberId"` // 用户ID

	//Wallet        *Wallet        `gorm:"foreignKey:WalletID" json:"wallet"`
	//CompanyMember *CompanyMember `gorm:"foreignKey:CompanyMemberID" json:"companyMember"`
}

// TableName get sql table name.获取数据库表名
func (m *WalletAuthorize) TableName() string {
	return "wallet_authorize"
}
