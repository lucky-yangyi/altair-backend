package model

import (
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

// Wallet 钱包表
type Wallet struct {
	Mixin
	Name           string  `gorm:"column:name" json:"name"`                      // 钱包别名
	Address        string  `gorm:"column:address" json:"address"`                // 钱包地址
	CoinID         uint64  `gorm:"column:coin_id" json:"coinID"`                 // 币种id
	Balance        float64 `gorm:"column:balance" json:"balance"`                // 余额
	Status         uint8   `gorm:"column:status" json:"status"`                  // 签名状态：1-待签，4-失败，127-成功
	TypeID         uint8   `gorm:"column:type_id" json:"typeId"`                 // 钱包类型id:1-多签，2-普通
	RequiredSigner uint    `gorm:"column:required_signer" json:"requiredSigner"` // 签名个数
	CompanyID      uint64  `gorm:"column:company_id" json:"companyID"`           // 组织id
	IsCollect      uint8   `gorm:"column:is_collect" json:"isCollect"`           // 是否归集 1-是 2-否
	AccessKey      string  `gorm:"column:access_key" json:"accessKey"`           // 公钥
	Sign           string  `gorm:"column:sign" json:"-"`
	SecretKey      string  `gorm:"column:secret_key" json:"-"`
}

// TableName get sql table name.获取数据库表名
func (w *Wallet) TableName() string {
	return "wallet"
}

type WalletAuth struct {
	Wallet
	IsHasAuth bool `json:"IsHasAuth"` //是否有该钱包权限
}

type WalletPrivate struct {
	Mixin
	Name           string  `gorm:"column:name" json:"name"`                      // 钱包别名
	Address        string  `gorm:"column:address" json:"address"`                // 钱包地址
	CoinID         uint64  `gorm:"column:coin_id" json:"coinID"`                 // 币种id
	Balance        float64 `gorm:"column:balance" json:"balance"`                // 余额
	Status         uint8   `gorm:"column:status" json:"status"`                  // 签名状态：1-待签，4-失败，127-成功
	TypeID         uint8   `gorm:"column:type_id" json:"typeId"`                 // 钱包类型id:1-多签，2-普通
	RequiredSigner uint    `gorm:"column:required_signer" json:"requiredSigner"` // 签名个数
	CompanyID      uint64  `gorm:"column:company_id" json:"companyID"`           // 组织id
	IsCollect      uint8   `gorm:"column:is_collect" json:"isCollect"`           // 是否归集 1-是 2-否
	AccessKey      string  `gorm:"column:access_key" json:"accessKey"`           // 公钥
	Sign           string  `gorm:"column:sign" json:"sign"`
	SecretKey      string  `gorm:"column:secret_key" json:"secretKey"`
}

// TableName get sql table name.获取数据库表名
func (w *WalletPrivate) TableName() string {
	return "wallet"
}

type WalletMix struct {
	Wallet
	Meta
}

type Meta struct {
	WalletType            *WalletType              `gorm:"foreignKey:TypeID" json:"walletType"`                            //钱包类型
	WalletCollect         []*WalletCollect         `gorm:"foreignKey:WalletID;references:ID" json:"walletCollect"`         //归集钱包
	Coin                  *Coin                    `gorm:"foreignKey:CoinID" json:"coin"`                                  //币种类型
	MultiSignWalletDetail []*MultiSignWalletDetail `gorm:"foreignKey:WalletID;references:ID" json:"multiSignWalletDetail"` //多签详情
	FilWalletMeta         *FilWalletMeta           `gorm:"foreignKey:ID;references:WalletID" json:"filWalletMeta"`         //file详情
}

// FilWalletMeta FIL钱包详情表
type FilWalletMeta struct {
	Mixin
	WalletID       uint64 `gorm:"column:wallet_id" json:"walletID"`             // 币种id
	RequiredSigner uint   `gorm:"column:required_signer" json:"requiredSigner"` // 签名个数
	IsMultiWallet  uint8  `gorm:"column:is_multi_wallet" json:"isMultiWallet"`  // 是否多签：1-是，2-否
}

// TableName get sql table name.获取数据库表名
func (m *FilWalletMeta) TableName() string {
	return "fil_wallet_meta"
}

//获取币种列表
func (w *Wallet) List(db *gorm.DB, pageOffset, pageSize int) (wallets []*Wallet, err error) {
	if pageOffset >= 0 && pageSize > 0 {
		db = db.Preload(clause.Associations).Offset(pageOffset).Limit(pageSize)
	}
	if w.Name != "" {
		db = db.Where("name = ?", w.Name)
	}
	db = db.Where("is_del = ?", w.IsDel)
	err = db.Where("is_del = ?", 0).Find(&wallets).Error
	if err != nil {
		return nil, err
	}
	return

}
