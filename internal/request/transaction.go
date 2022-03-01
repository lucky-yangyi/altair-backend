package request

type TransactionList struct {
	WalletId  int   `json:"walletId"`
	Type      int   `json:"type"`
	StartTime int64 `json:"startTime"`
	EndTime   int64 `json:"endTime"`
	PageReq
}

type AddTransaction struct {
	WalletID          uint64  `json:"walletId" binding:"required"`          //钱包ID
	DestinationWallet string  `json:"destinationWallet" binding:"required"` //接收方钱包地址
	Amount            float64 `json:"amount" binding:"required"`            //金额
	Comment           string  `json:"comment"`                              //备注
}

type UpdateTransactionDetail struct {
	TransactionID      uint64  `json:"transactionId" binding:"required"` //交易id
	WalletID           uint64  `json:"walletId" binding:"required"`      //钱包id
	Status             uint8   `json:"status" binding:"required"`        //交易状态 127:已签 1:待签 4:驳回
	ChainTransactionId string  `json:"chainTransactionId"`               //链上id
	GasAmount          float64 `json:"gasAmount"`                        //gas费用
}

type TransactionDetail struct {
	TransactionId uint64 `json:"transactionId"` //交易id
}

type GetAllTransactionRequest struct {
	PageNo             uint64   `json:"pageNo" binding:"required"`   //`alias:"页码数"`
	PageSize           uint64   `json:"pageSize" binding:"required"` //`alias:"每页数量"`
	ID                 uint64   `json:"id"`                          //`alias:"交易ID"`
	SerialNumber       string   `json:"serialNumber"`                //`alias:"交易序列号"`
	FromWalletID       uint64   `json:"fromWalletId"`                //`alias:"发起方钱包ID"`
	DestinationWallet  string   `json:"destinationWallet"`           //`alias:"接收方钱包地址"`
	Status             int8     `json:"status"`                      //`alias:"交易状态:127-成功，1-已发起，4-失败
	Type               int8     `json:"type"`                        //`alias:"交易类型"`
	ChainTransactionId string   `json:"chainTransactionId"`          // `alias:"链上交易id"`
	Daterange          []string `json:"daterange"`                   //`alias:"时间区间"`
}
