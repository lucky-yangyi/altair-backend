package request

import (
	"altair-backend/pkg/address"
	"github.com/filecoin-project/go-state-types/abi"
)

type ChainCollectTransaction struct {
	FromAddr   string `json:"fromAddr"`
	PrivateKey []byte `json:"privateKey"`
	ToAddr     string `json:"toAddr"`
}

type SendFromOffline struct {
	Symbol   string `json:"symbol" binding:"required"`
	SignType string `json:"signType" binding:"required"`
	Address  string `json:"address" binding:"required"`
	Data     string `json:"data" binding:"required"`
	Post     bool   `json:"post" binding:"required"`
}

type SingleTransaction struct {
	From     string `json:"from" binding:"required"`
	To       string `json:"to" binding:"required"`
	Quantity string `json:"quantity" binding:"required"`
}

//创建多签钱包
type CreateMsig struct {
	Address  string `json:"addrstrs" binding:"required"`
	Require  uint64 `json:"required" binding:"required"`
	Quantity string `json:"quantity" binding:"required"`
	Duration uint64 `json:"duration" binding:"required"`
}

//提议多签
type ProposeCreateMsig struct {
	MultisigAddress    string `json:"multisigAddress" binding:"required"`
	DestinationAddress string `json:"destinationAddress" binding:"required"`
	Quantity           string `json:"quantity" binding:"required"`
}

//同意多签
type ApproveCreateMsig struct {
	MultisigAddress string `json:"multisigAddress" binding:"required"`
	MessageId       uint64 `json:"messageId" binding:"required"`
	TxId            string `json:"txId" binding:"required"`
}

// 上链交易
type SendRawTransaction struct {
	Signature  string `json:"signature" binding:"required"`
	EmptyTrans string `json:"emptyTrans" binding:"required,json"`
}

// 离线 EmptyTrans 结构体
type Message struct {
	Version    uint64          `json:"Version" binding:"required"`
	To         address.Address `json:"To" binding:"required"`
	From       address.Address `json:"From" binding:"required"`
	Nonce      uint64          `json:"Nonce" binding:"required"`
	Value      abi.TokenAmount `json:"Value" binding:"required"`
	GasLimit   int64           `json:"GasLimit" binding:"required"`
	GasFeeCap  abi.TokenAmount `json:"GasFeeCap" binding:"required"`
	GasPremium abi.TokenAmount `json:"GasPremium" binding:"required"`
	Method     uint64          `json:"Method" binding:"required"`
	Params     []byte          `json:"Params" binding:"required"`
}

type Receipt struct {
	Symbol string `json:"symbol" binding:"required"`
	TxtId  string `json:"txId" binding:"required"`
}

type DataStruct struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Quantity string `json:"quantity"`
}

type EmptyTransStruct struct {
	Version    int             `json:"version"`
	To         string          `json:"to"`
	From       string          `json:"from"`
	Nonce      uint64          `json:"nonce"`
	Value      abi.TokenAmount `json:"value"`
	GasLimit   int64           `json:"gasLimit"`
	GasFeeCap  string          `json:"gasFeeCap"`
	GasPremium string          `json:"gasPremium"`
	Method     uint64          `json:"method"`
	Params     []byte          `json:"params"`
}
