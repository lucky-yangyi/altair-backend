package request

type CoinListReq struct {
	Name string `json:"name"`
	Code string `json:"code"`
	PageReq
}

type CoinCreate struct {
	Name string `json:"name" binding:"required"`
	Code string `json:"code" binding:"required"`
}
