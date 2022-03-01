package model

// MonthBill 月付账单
type MonthBill struct {
	Mixin
	Name      string  `gorm:"column:name;type:varchar(50);default:''" json:"name"`
	Address   string  `gorm:"column:address;type:varchar(50);default:''" json:"address"`
	WalletId  uint    `gorm:"column:wallet_id;type:int unsigned;default:0" json:"walletId"` // 钱包
	Month     string  `gorm:"column:month;type:varchar(32) default:'0'" json:"month"`       // 月份 yyyymm
	Amount    float64 `gorm:"column:amount;type:decimal(4,0) unsigned;not null;default:0" json:"amount"`
	PayStatus uint8   `gorm:"column:pay_status;type:tinyint unsigned;default:0" json:"PayStatus"` // 状态 0未付清、1已付清
	IsDel     uint8   `gorm:"column:is_del;type:tinyint unsigned;default:0" json:"isDel"`         // 0: 未删除 1: 删除

	//Wallet *Wallet `gorm:"foreignKey:WalletId" json:"wallet"`
}

// TableName get sql table name.获取数据库表名
func (m *MonthBill) TableName() string {
	return "month_bill"
}
