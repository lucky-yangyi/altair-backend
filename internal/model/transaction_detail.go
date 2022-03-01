package model

// TransactionDetail [...]
type TransactionDetail struct {
	Mixin
	TransactionID      uint64  `gorm:"column:transaction_id" json:"transactionId"`            // 交易ID
	WalletID           uint64  `gorm:"column:wallet_id" json:"walletId"`                      // 普通钱包ID
	GasAmount          float64 `gorm:"column:gas_amount" json:"gasAmount"`                    // gas金额
	Status             uint8   `gorm:"column:status" json:"status"`                           // 交易状态 127:已签 1:待签 4:驳回
	ChainTransactionID string  `gorm:"column:chain_transaction_id" json:"chainTransactionId"` // 链上cid
	SerialNumber       string  `gorm:"column:serial_number" json:"serialNumber"`              // 交易流水号
	IsSponsor          int     `gorm:"column:is_sponsor" json:"isSponsor"`                    // 是否是发起者

	Wallet      Wallet      `gorm:"foreignKey:WalletID" json:"wallet"`           //钱包
	Transaction Transaction `gorm:"foreignKey:TransactionID" json:"transaction"` //交易表
}

// TableName get sql table name.获取数据库表名
func (m *TransactionDetail) TableName() string {
	return "transaction_detail"
}
