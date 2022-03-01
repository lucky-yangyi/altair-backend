package model

// MultiSignWalletDetail 多签钱包详情/成员表
type MultiSignWalletDetail struct {
	Mixin
	WalletID         uint64 `gorm:"column:wallet_id" json:"walletId"`                  // 多签钱包ID
	OrdinaryWalletID uint64 `gorm:"column:ordinary_wallet_id" json:"ordinaryWalletId"` // 普通钱包ID
	Status           uint8  `gorm:"column:status" json:"status"`                       // 签名状态：1-待签，4-驳回，127-完成

	Wallet *Wallet `gorm:"foreignKey:OrdinaryWalletID" json:"wallet"`
}

// TableName get sql table name.获取数据库表名
func (m *MultiSignWalletDetail) TableName() string {
	return "multi_sign_wallet_detail"
}
