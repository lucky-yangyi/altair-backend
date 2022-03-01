package model

// OrdinaryWall 元钱包表
type OrdinaryWall struct {
	Mixin
	Name       string `gorm:"column:name" json:"name"`               // 钱包别名
	Address    string `gorm:"column:address" json:"address"`         // 钱包公钥地址
	CurrencyId uint64 `gorm:"column:currency_id" json:"currency_id"` // 币种id
}

// TableName get sql table name.获取数据库表名
func (o *OrdinaryWall) TableName() string {
	return "ordinary_wallet"
}
