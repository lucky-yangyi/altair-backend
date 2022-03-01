package model

import "time"

// CollectTransaction 归集交易
type CollectTransaction struct {
	ID          uint      `gorm:"primaryKey;column:id;type:int unsigned;not null" json:"-"`
	FromAddress string    `gorm:"column:from_address;type:varchar(128)" json:"fromAddress"`
	ToAddress   string    `gorm:"column:to_address;type:varchar(128)" json:"toAddress"`
	Amount      float64   `gorm:"column:amount;type:double" json:"amount"`
	Cid         string    `gorm:"column:cid;type:varchar(256)" json:"cid"`
	CreatedAt   time.Time `gorm:"column:created_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:timestamp;default:CURRENT_TIMESTAMP" json:"updatedAt"`
}

// TableName get sql table name.获取数据库表名
func (m *CollectTransaction) TableName() string {
	return "collect_transaction"
}
