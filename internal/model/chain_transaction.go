package model

// SysTransaction 平台交易数据
type ChainTransaction struct {
	Mixin
	Hash             string `gorm:"column:hash" json:"hash"`
	BlockNumber      string `gorm:"column:block_number" json:"blockNumber"`
	BlockHash        string `gorm:"column:block_hash" json:"blockHash"`
	FromAddress      string `gorm:"column:from_address" json:"fromAddress"`
	ToAddress        string `gorm:"column:to_address" json:"toAddress"`
	Gas              string `gorm:"column:gas" json:"gas"`
	Value            string `gorm:"column:value" json:"value"`
	TransactionIndex int64  `gorm:"column:transaction_index" json:"transactionIndex"`
	Timestamp        int64  `gorm:"column:timestamp" json:"timestamp"`
	Nonce            int64  `gorm:"column:nonce" json:"nonce"`
	BlockHeight      int64  `gorm:"column:block_height" json:"blockHeight"`
	Method           int    `gorm:"column:method" json:"method"` // 0-默认，1-to和from钱包都在系统中 2-单from在系统中（转出） 3-单to在系统中（转入）
	Status           int    `gorm:"column:status" json:"status"`
}

// TableName get sql table name.获取数据库表名
func (m *ChainTransaction) TableName() string {
	return "chain_transaction"
}
