package model

// WalletCollect 归集钱包表
type WalletCollect struct {
	Mixin
	WalletID   int64  `gorm:"column:wallet_id" json:"walletId"` // 普通钱包id
	Address    string `gorm:"column:address" json:"address"`    // 钱包公钥地址
	Symbol     string `gorm:"column:symbol" json:"symbol"`      //标识
	PrivateKey string `gorm:"column:private_key" json:"-"`      // 钱包私钥
}

// WalletCollectPrivate 私钥可见
type WalletCollectPrivate struct {
	Mixin
	WalletID   int64  `gorm:"column:wallet_id" json:"walletId"`     // 普通钱包id
	Address    string `gorm:"column:address" json:"address"`        // 钱包公钥地址
	Symbol     string `gorm:"column:symbol" json:"symbol"`          //标识
	PrivateKey string `gorm:"column:private_key" json:"privateKey"` // 私钥
}

// TableName get sql table name.获取数据库表名
func (m *WalletCollect) TableName() string {
	return "wallet_collect"
}
