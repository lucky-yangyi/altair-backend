package model

// Transaction [...]
type Transaction struct {
	Mixin
	ChainTransactionID string  `gorm:"column:chain_transaction_id" json:"chainTransactionId"` //链上交易id
	Type               uint8   `gorm:"column:type" json:"type"`                               // 交易类型 1 转出，2 转入，3 其它
	SubType            uint8   `gorm:"column:sub_type" json:"subType"`                        // 子类型:1-缺省/默认，2-创建多签，3-修改多签签名
	Comment            string  `gorm:"column:comment" json:"comment"`                         // 交易备注
	FromWalletID       uint64  `gorm:"column:from_wallet_id" json:"fromWalletId"`             // 发起方钱包ID
	DestinationWallet  string  `gorm:"column:destination_wallet" json:"destinationWallet"`    // 接收方钱包地址
	Amount             float64 `gorm:"column:amount" json:"amount"`                           // 金额
	GasAmount          float64 `gorm:"column:gas_amount" json:"gasAmount"`                    // gas金额
	Status             uint8   `gorm:"column:status" json:"status"`                           // 交易状态：127-成功，1-已发起，4-失败
	SerialNumber       Uuid    `gorm:"column:serial_number" json:"serialNumber"`              // 交易流水号
	CoinID             uint8   `gorm:"column:coin_id" json:"coinId"`                          // 币种id
	CompanyMemberID    uint64  `gorm:"column:company_member_id" json:"companyMemberId"`       // 发起人id
}

type Tmix struct {
	Transaction
	Wallet            *Wallet              `gorm:"foreignKey:FromWalletID" json:"wallet"`                //发起方钱包
	TransactionDetail []*TransactionDetail `gorm:"foreignKey:TransactionID;references:ID" json:"detail"` //交易详情
	CompanyMember     *CMMeta              `gorm:"foreignKey:CompanyMemberID" json:"companyMember"`      //用户（组织成员）详情
	Coin              *Coin                `gorm:"foreignKey:CoinID" json:"coin"`                        // 币种详情
}

// TableName get sql table name.获取数据库表名
func (m *Transaction) TableName() string {
	return "transaction"
}
